package services

import (
	"database/sql"

	"github.com/aniket-skroman/skroman_sales_service.git/apis/dto"
	"github.com/aniket-skroman/skroman_sales_service.git/apis/helper"
	"github.com/aniket-skroman/skroman_sales_service.git/apis/repositories"
	db "github.com/aniket-skroman/skroman_sales_service.git/sqlc_lib"
	"github.com/google/uuid"
)

type LeadOrderService interface {
	CreateLeadOrder(req dto.CreateLeadOrderRequestDTO) (dto.LeadOrderDTO, error)
	FetchOrdersByLeadId(lead_id uuid.UUID) ([]dto.LeadOrderDTO, error)
	DeleteLeadOrder(req dto.DeleteLeadOrderRequestDTO) error
	UpdateLeadOrder(req dto.UpdateLeadOrderRequestDTO, order_id string) (dto.LeadOrderDTO, error)
	FetchOrdersByOrderId(order_id string) (dto.LeadOrderDTO, error)
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
	// lead_obj_id, err := uuid.Parse(lead_id)

	// if err != nil {
	// 	return nil, helper.ERR_INVALID_ID
	// }

	args := uuid.NullUUID{UUID: lead_id, Valid: true}

	orders, err := serv.order_repo.FetchOrdersByLeadId(args)

	if err != nil {
		return nil, err
	}

	if len(orders) == 0 {
		return nil, sql.ErrNoRows
	}

	result := new(dto.LeadOrderDTO).MakeLeadOrderDTO(orders...)
	if data, ok := result.(dto.LeadOrderDTO); ok {
		return []dto.LeadOrderDTO{data}, nil
	}

	return result.([]dto.LeadOrderDTO), nil

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
