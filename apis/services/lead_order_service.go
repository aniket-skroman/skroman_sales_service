package services

import (
	"database/sql"
	"sync"

	"github.com/aniket-skroman/skroman_sales_service.git/apis/dto"
	"github.com/aniket-skroman/skroman_sales_service.git/apis/helper"
	"github.com/aniket-skroman/skroman_sales_service.git/apis/repositories"
	"github.com/aniket-skroman/skroman_sales_service.git/connections"
	db "github.com/aniket-skroman/skroman_sales_service.git/sqlc_lib"
	"github.com/google/uuid"
)

type LeadOrderService interface {
	CreateLeadOrder(req dto.CreateLeadOrderRequestDTO) (dto.LeadOrderDTO, error)
	FetchOrdersByLeadId(lead_id uuid.UUID) ([]dto.LeadOrderDTO, error)
	DeleteLeadOrder(req dto.DeleteLeadOrderRequestDTO) error
	UpdateLeadOrder(req dto.UpdateLeadOrderRequestDTO, order_id string) (dto.LeadOrderDTO, error)
	FetchOrdersByOrderId(order_id string) (dto.LeadOrderDTO, error)
	UploadOrderQuatation(req dto.UploadOrderQuatationRequestDTO) error
	FetchQuatationByLeadId(lead_id uuid.UUID) ([]db.OrderQuatation, error)
	DeleteQuotation(req dto.DeleteOrderQuotationRequestDTO) error
}

type lead_order_serv struct {
	order_repo repositories.LeadOrderRepository
}

func NewLeadOrderService(repo repositories.LeadOrderRepository) LeadOrderService {
	return &lead_order_serv{
		order_repo: repo,
	}
}

func (serv *lead_order_serv) CreateLeadOrder(req dto.CreateLeadOrderRequestDTO) (dto.LeadOrderDTO, error) {
	lead_obj_id, err := uuid.Parse(req.LeadID)

	if err != nil {
		return dto.LeadOrderDTO{}, helper.ERR_INVALID_ID
	}

	args := db.CreateLeadOrderParams{
		LeadID:      uuid.NullUUID{UUID: lead_obj_id, Valid: true},
		DeviceType:  sql.NullString{String: req.DeviceType, Valid: true},
		DeviceModel: sql.NullString{String: req.DeviceModel, Valid: true},
		DevicePrice: sql.NullInt32{Int32: req.DevicePrice, Valid: true},
		DeviceName:  sql.NullString{String: req.DeviceName, Valid: true},
		Quantity:    sql.NullInt32{Int32: req.Quantity, Valid: true},
	}

	order, err := serv.order_repo.CreateLeadOrder(args)

	err = helper.Handle_db_err(err)

	if err != nil {
		return dto.LeadOrderDTO{}, err
	}

	return new(dto.LeadOrderDTO).MakeLeadOrderDTO(order).(dto.LeadOrderDTO), nil
}

func (serv *lead_order_serv) FetchOrdersByLeadId(lead_id uuid.UUID) ([]dto.LeadOrderDTO, error) {

	args := uuid.NullUUID{UUID: lead_id, Valid: true}

	orders, err := serv.order_repo.FetchOrdersByLeadId(args)

	if err != nil {
		return nil, err
	}

	if len(orders) == 0 {
		return nil, sql.ErrNoRows
	}

	result := make([]dto.LeadOrderDTO, len(orders))
	wg := sync.WaitGroup{}

	for i, data := range orders {
		wg.Add(1)
		go serv.setOrderData(&wg, &result[i], &data)
	}
	wg.Wait()

	return result, nil

}

func (ser *lead_order_serv) setOrderData(wg *sync.WaitGroup, result *dto.LeadOrderDTO, data *db.LeadOrder) {
	defer wg.Done()

	*result = dto.LeadOrderDTO{
		ID:          data.ID,
		LeadID:      data.LeadID,
		DeviceType:  data.DeviceType.String,
		DeviceModel: data.DeviceModel.String,
		DeviceName:  data.DeviceName.String,
		DevicePrice: data.DevicePrice.Int32,
		Quantity:    data.Quantity.Int32,
		CreatedAt:   data.CreatedAt,
		UpdatedAt:   data.UpdatedAt,
	}
}

func (serv *lead_order_serv) DeleteLeadOrder(req dto.DeleteLeadOrderRequestDTO) error {
	order_obj_id, err := uuid.Parse(req.OrderId)

	if err != nil {
		return helper.ERR_INVALID_ID
	}

	lead_obj_id, err := uuid.Parse(req.LeadId)

	if err != nil {
		return helper.ERR_INVALID_ID
	}

	args := db.DeleteLeadOrderParams{
		ID:     order_obj_id,
		LeadID: uuid.NullUUID{UUID: lead_obj_id, Valid: true},
	}

	result, err := serv.order_repo.DeleteLeadOrder(args)

	if err != nil {
		return err
	}

	if rows, _ := result.RowsAffected(); rows == 0 {
		return helper.Err_Delete_Failed
	}

	return nil
}

func (serv *lead_order_serv) UpdateLeadOrder(req dto.UpdateLeadOrderRequestDTO, order_id string) (dto.LeadOrderDTO, error) {
	order_obj_id, err := uuid.Parse(order_id)

	if err != nil {
		return dto.LeadOrderDTO{}, helper.ERR_INVALID_ID
	}

	lead_obj_id, err := uuid.Parse(req.LeadID)

	if err != nil {
		return dto.LeadOrderDTO{}, helper.ERR_INVALID_ID
	}

	args := db.UpdateLeadOrderParams{
		ID:          order_obj_id,
		LeadID:      uuid.NullUUID{UUID: lead_obj_id, Valid: true},
		DeviceType:  sql.NullString{String: req.DeviceType, Valid: true},
		DeviceModel: sql.NullString{String: req.DeviceModel, Valid: true},
		DevicePrice: sql.NullInt32{Int32: req.DevicePrice, Valid: true},
		Quantity:    sql.NullInt32{Int32: req.Quantity, Valid: true},
		DeviceName:  sql.NullString{String: req.DeviceName, Valid: true},
	}

	result, err := serv.order_repo.UpdateLeadOrder(args)

	if err != nil {
		return dto.LeadOrderDTO{}, err
	}

	return new(dto.LeadOrderDTO).MakeLeadOrderDTO(result).(dto.LeadOrderDTO), nil
}

func (serv *lead_order_serv) FetchOrdersByOrderId(order_id string) (dto.LeadOrderDTO, error) {
	order_obj_id, err := uuid.Parse(order_id)

	if err != nil {
		return dto.LeadOrderDTO{}, helper.ERR_INVALID_ID
	}

	result, err := serv.order_repo.FetchOrdersByOrderId(order_obj_id)

	return new(dto.LeadOrderDTO).MakeLeadOrderDTO(result).(dto.LeadOrderDTO), err
}

func (serv *lead_order_serv) UploadOrderQuatation(req dto.UploadOrderQuatationRequestDTO) error {
	// upload a file first
	s3_connection := connections.NewS3Connection()
	path, err := s3_connection.UploadOrderQuatation(req.QuatationFile, &req.FileHandler)

	if err != nil {
		return err
	}

	// setting a file path to process with DB
	req.QuatationLink = path

	// prepare for args
	args := db.CreateNewOrderQuatationParams{
		LeadID:        req.LeadId,
		GeneratedBy:   req.GeneratedBy,
		QuatationLink: req.QuatationLink,
	}

	err = serv.order_repo.UploadOrderQuatation(args)
	return err
}

func (serv *lead_order_serv) FetchQuatationByLeadId(lead_id uuid.UUID) ([]db.OrderQuatation, error) {

	return serv.order_repo.FetchOrderQutationsByLeadId(lead_id)
}

// delete a quotation
func (ser *lead_order_serv) DeleteQuotation(req dto.DeleteOrderQuotationRequestDTO) error {
	// validate id's
	lead_obj_id, err := helper.ValidateUUID(req.LeadId)

	if err != nil {
		return err
	}

	quotation_obj_id, err := helper.ValidateUUID(req.QuotationId)

	if err != nil {
		return err
	}

	// fetch file path
	quotation, err := ser.order_repo.FetchQuotationById(quotation_obj_id)

	if err != nil {
		return err
	}

	if quotation.QuatationLink == "" {
		return sql.ErrNoRows
	}

	// delete first file from remote server
	s3_connection := connections.NewS3Connection()
	err = s3_connection.DeleteFiles(quotation.QuatationLink)

	if err != nil {
		return err
	}

	//  then remove file ref
	args := db.DeleteOrderQuotationParams{
		LeadID: lead_obj_id,
		ID:     quotation_obj_id,
	}

	_, err = ser.order_repo.DeleteQuotation(args)

	return err
}
