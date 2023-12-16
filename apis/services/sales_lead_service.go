package services

import (
	"database/sql"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"reflect"
	"strconv"
	"sync"

	"github.com/aniket-skroman/skroman_sales_service.git/apis/dto"
	"github.com/aniket-skroman/skroman_sales_service.git/apis/helper"
	proxycalls "github.com/aniket-skroman/skroman_sales_service.git/apis/proxy_calls"
	"github.com/aniket-skroman/skroman_sales_service.git/apis/repositories"
	db "github.com/aniket-skroman/skroman_sales_service.git/sqlc_lib"
	"github.com/aniket-skroman/skroman_sales_service.git/utils"
	"github.com/google/uuid"
)

type SalesLeadService interface {
	CreateNewLead(req dto.CreateNewLeadDTO) (dto.SaleLeadsDTO, error)
	FetchAllLeads(req dto.FetchAllLeadsRequestDTO) (interface{}, error)
	FetchLeadByLeadId(lead_id uuid.UUID) (interface{}, error)
	IncreaeQuatationCount(lead_id uuid.UUID) error
	FetchLeadCounts() (dto.FetchLeadCountsDTO, error)
	FetchLeadsByStatus(req dto.FetchLeadsByStatusRequestDTO) (interface{}, error)
	CancelLead(req dto.CancelLeadRequestDTO) error
	FetchCancelLeadByLeadId(lead_id string) (dto.CancelLeadsDTO, error)
	Fetch_users_data(user_id string) (interface{}, error)
}

type sale_service struct {
	sale_lead_repo  repositories.SalesRepository
	lead_order_serv LeadOrderService
	jwt_service     JWTService
}

func NewSalesLeadService(sale_lead_repo repositories.SalesRepository, lead_order_serv LeadOrderService, jwt_Service JWTService) SalesLeadService {
	return &sale_service{
		sale_lead_repo:  sale_lead_repo,
		lead_order_serv: lead_order_serv,
		jwt_service:     jwt_Service,
	}
}

func (ser *sale_service) CreateNewLead(req dto.CreateNewLeadDTO) (dto.SaleLeadsDTO, error) {
	lead_by, err := uuid.Parse(req.LeadBy)

	if err != nil {
		return dto.SaleLeadsDTO{}, err
	}

	_, err = strconv.Atoi(req.ReferalContact)

	if err != nil || len(req.ReferalContact) != 10 {
		return dto.SaleLeadsDTO{}, helper.ERR_REQUIRED_PARAMS
	}

	args := db.CreateNewLeadParams{
		LeadBy:         lead_by,
		ReferalName:    req.ReferalName,
		ReferalContact: req.ReferalContact,
		QuatationCount: sql.NullInt32{Int32: 1, Valid: true},
		Status:         "INIT",
	}

	new_lead, err := ser.sale_lead_repo.CreateSalesLead(args)
	if err != nil {
		return dto.SaleLeadsDTO{}, err
	}
	sale_lead := new(dto.SaleLeadsDTO).MakeSaleLeadsDTO(new_lead)
	return sale_lead.(dto.SaleLeadsDTO), nil
}
func (ser *sale_service) FetchAllLeads(req dto.FetchAllLeadsRequestDTO) (interface{}, error) {

	wg := sync.WaitGroup{}
	wg.Add(2)
	var err error
	var result []db.FetchAllLeadsRow
	sales_lead := []dto.SaleLeadsDTO{}
	go func() {
		defer wg.Done()
		args := db.FetchAllLeadsParams{
			Limit:  int32(req.PageSize),
			Offset: (int32(req.PageId) - 1) * int32(req.PageSize),
		}

		result, err = ser.sale_lead_repo.FetchAllLeads(args)
		if err != nil {
			return
		}

		sales_lead = make([]dto.SaleLeadsDTO, len(result))
		t_wg := sync.WaitGroup{}

		for i := range result {
			t_wg.Add(1)
			go ser.setLeadData(&t_wg, &sales_lead[i], &result[i])
		}
		t_wg.Wait()
	}()

	go func() {
		defer wg.Done()
		count, _ := ser.sale_lead_repo.CountSalesLead()
		helper.SetPaginationData(int32(req.PageId), int32(count))
	}()

	wg.Wait()

	if err != nil {
		return nil, err
	}
	if len(result) == 0 {
		return nil, sql.ErrNoRows
	}

	return sales_lead, err
}

func (ser *sale_service) setLeadData(wg *sync.WaitGroup, result *dto.SaleLeadsDTO, data *db.FetchAllLeadsRow) {
	defer wg.Done()
	temp := dto.SaleLeadsDTO{
		ID:             data.LeadID,
		LeadBy:         data.LeadBy,
		ReferalName:    data.ReferalName,
		ReferalContact: data.ReferalContact,
		Status:         data.Status,
		QuatationCount: data.QuatationCount.Int32,
		CreatedAt:      data.LeadCreatedAt,
		UpdatedAt:      data.LeadUpdatedAt,
		IsLeadInfo:     data.IsLeadInfo.Bool,
		IsOrderInfo:    data.IsOrderInfo.Bool,
		LeadInfo: &dto.GetLeadInfoDTO{
			ID:           data.LeadInfoID.UUID,
			Name:         data.Name.String,
			Email:        data.Email.String,
			Contact:      data.Contact.String,
			AddressLine1: data.AddressLine1.String,
			City:         data.City.String,
			State:        data.State.String,
			LeadType:     data.LeadType.String,
			CreatedAt:    data.LeadInfoCreatedAt.Time,
			UpdatedAt:    data.LeadInfoUpdatedAt.Time,
		},
	}

	if (temp.LeadInfo == &dto.GetLeadInfoDTO{}) {
		temp.LeadInfo = nil
	} else {
		temp.LeadInfo.LeadID = data.LeadID
	}

	*result = temp

}

func (ser *sale_service) FetchLeadByLeadId(lead_id uuid.UUID) (interface{}, error) {

	wg := sync.WaitGroup{}
	wg.Add(3)

	var result dto.SaleLeadsDTO
	err_chan := make(chan error)

	go func() {
		defer wg.Done()
		lead, err := ser.sale_lead_repo.FetchLeadByLeadId(lead_id)
		if err != nil {
			err_chan <- err
			return
		}

		result.ID = lead.LeadID
		result.LeadBy = lead.LeadBy
		result.ReferalName = lead.ReferalName
		result.ReferalContact = lead.ReferalContact
		result.Status = lead.Status
		result.QuatationCount = lead.QuatationCount.Int32
		result.IsLeadInfo = lead.IsLeadInfo.Bool
		result.IsOrderInfo = lead.IsOrderInfo.Bool
		result.CreatedAt = lead.LeadCreatedAt
		result.UpdatedAt = lead.LeadUpdatedAt

		if lead.LeadInfoID.UUID != uuid.Nil {
			result.LeadInfo = &dto.GetLeadInfoDTO{
				ID:           lead.LeadInfoID.UUID,
				LeadID:       lead_id,
				Name:         lead.Name.String,
				Email:        lead.Email.String,
				Contact:      lead.Contact.String,
				AddressLine1: lead.AddressLine1.String,
				City:         lead.City.String,
				State:        lead.State.String,
				LeadType:     lead.LeadType.String,
				CreatedAt:    lead.LeadInfoCreatedAt.Time,
				UpdatedAt:    lead.LeadInfoUpdatedAt.Time,
			}
		}
	}()

	go func() {
		defer wg.Done()
		if orders, err := ser.lead_order_serv.FetchOrdersByLeadId(lead_id); err != nil {
			err_chan <- err
			return
		} else {
			result.LeadOrders = &orders
		}
	}()

	go func() {
		defer wg.Done()
		order_result, err := ser.lead_order_serv.FetchQuatationByLeadId(lead_id)

		if err != nil {
			err_chan <- err
			return
		}

		quotations := make([]dto.OrderQuatation, len(order_result))
		t_w := sync.WaitGroup{}
		for i, quotation := range order_result {
			t_w.Add(1)
			go ser.setOrderQuotationData(&t_w, &quotations[i], &quotation)
		}
		t_w.Wait()

		if len(quotations) != 0 {
			result.OrderQuatations = &quotations
		}
	}()

	go func() {
		wg.Wait()
		close(err_chan)
	}()

	for data_err := range err_chan {
		if data_err != sql.ErrNoRows {
			return nil, data_err
		}
	}

	return result, nil
}

func (ser *sale_service) setOrderQuotationData(wg *sync.WaitGroup, result *dto.OrderQuatation, data *db.OrderQuatation) {
	defer wg.Done()

	*result = dto.OrderQuatation{
		ID:            data.ID,
		QuotationLink: utils.QUOTATION_PATH_URL + data.QuatationLink,
		CreatedAt:     data.CreatedAt,
		UpdatedAt:     data.UpdatedAt,
	}
}

func (ser *sale_service) IncreaeQuatationCount(lead_id uuid.UUID) error {

	result, err := ser.sale_lead_repo.IncreaeQuatationCount(lead_id)

	if err != nil {
		return err
	}

	if result == 0 {
		return errors.New("failed to increase a quatation count")
	}

	return nil
}

func (ser *sale_service) FetchLeadCounts() (dto.FetchLeadCountsDTO, error) {
	wg := sync.WaitGroup{}
	err_chan := make(chan error)
	result := dto.FetchLeadCountsDTO{}
	wg.Add(2)

	go func() {
		defer wg.Done()
		counts, err := ser.sale_lead_repo.FetchLeadCounts()

		if err != nil {
			err_chan <- err
			return
		}

		result.LeadCount = counts

	}()

	go func() {
		defer wg.Done()

		counts, err := ser.sale_lead_repo.FetchLeadCountMonthWise()
		if err != nil {
			err_chan <- err
			return
		}
		result.LeadMonthCounts = counts
	}()

	go func() {
		wg.Wait()
		close(err_chan)
	}()

	for err_data := range err_chan {
		return dto.FetchLeadCountsDTO{}, err_data
	}

	return result, nil
}

func (ser *sale_service) FetchLeadsByStatus(req dto.FetchLeadsByStatusRequestDTO) (interface{}, error) {
	wg := sync.WaitGroup{}
	wg.Add(2)
	var err error
	var result []db.FetchLeadsByStatusRow
	sales_lead := []dto.SaleLeadsDTO{}
	go func() {
		defer wg.Done()
		args := db.FetchLeadsByStatusParams{
			Limit:  int32(req.PageSize),
			Offset: (int32(req.PageId) - 1) * int32(req.PageSize),
			Status: req.Status,
		}

		result, err = ser.sale_lead_repo.FetchLeadsByStatus(args)

		if err != nil {
			return
		}

		for i := range result {
			temp := dto.SaleLeadsDTO{
				ID:             result[i].LeadID,
				LeadBy:         result[i].LeadBy,
				ReferalName:    result[i].ReferalName,
				ReferalContact: result[i].ReferalContact,
				Status:         result[i].Status,
				QuatationCount: result[i].QuatationCount.Int32,
				CreatedAt:      result[i].LeadCreatedAt,
				UpdatedAt:      result[i].LeadUpdatedAt,
				IsLeadInfo:     result[i].IsLeadInfo.Bool,
				IsOrderInfo:    result[i].IsOrderInfo.Bool,
				LeadInfo: &dto.GetLeadInfoDTO{
					ID:           result[i].LeadInfoID.UUID,
					Name:         result[i].Name.String,
					Email:        result[i].Email.String,
					Contact:      result[i].Contact.String,
					AddressLine1: result[i].AddressLine1.String,
					City:         result[i].City.String,
					State:        result[i].State.String,
					LeadType:     result[i].LeadType.String,
					CreatedAt:    result[i].LeadInfoCreatedAt.Time,
					UpdatedAt:    result[i].LeadInfoUpdatedAt.Time,
				},
			}
			if (reflect.DeepEqual(temp.LeadInfo, &dto.GetLeadInfoDTO{})) {
				temp.LeadInfo = nil
			} else {
				temp.LeadInfo.LeadID = result[i].LeadID
			}
			sales_lead = append(sales_lead, temp)
		}
	}()

	go func() {
		defer wg.Done()
		count, _ := ser.sale_lead_repo.FetchPGCountLeadsByStatus(req.Status)
		helper.SetPaginationData(int32(req.PageId), int32(count))
	}()

	wg.Wait()

	if err != nil {
		return nil, err
	}
	if len(result) == 0 {
		return nil, sql.ErrNoRows
	}

	return sales_lead, err
}

func (ser *sale_service) CancelLead(req dto.CancelLeadRequestDTO) error {
	lead_obj_id, err := uuid.Parse(req.LeadId)

	if err != nil {
		return helper.ERR_INVALID_ID
	}

	cancel_by, err := uuid.Parse(utils.TOKEN_ID)
	if err != nil {
		return helper.Err_Something_Wents_Wrong
	}

	args := db.CreateCancelLeadParams{
		LeadID:   lead_obj_id,
		Reason:   req.Reason,
		CancelBy: cancel_by,
	}

	err = ser.sale_lead_repo.CancelLead(args)

	return helper.Handle_db_err(err)
}

func (ser *sale_service) FetchCancelLeadByLeadId(lead_id string) (dto.CancelLeadsDTO, error) {
	lead_obj_id, err := uuid.Parse(lead_id)

	if err != nil {
		return dto.CancelLeadsDTO{}, helper.ERR_INVALID_ID
	}

	result, err := ser.sale_lead_repo.FetchCancelLead(lead_obj_id)

	if err != nil {
		return dto.CancelLeadsDTO{}, err
	}

	user, err := ser.Fetch_users_data(result.CancelBy.String())
	cancel_lead := dto.CancelLeadsDTO{}

	if err != nil {
		cancel_lead.CancelBy = nil
	} else {
		cancel_lead.CancelBy = user
	}

	cancel_lead.ID = result.ID
	cancel_lead.LeadID = result.LeadID
	cancel_lead.Reason = result.Reason
	cancel_lead.CreatedAt = result.CreatedAt
	cancel_lead.UpdatedAt = result.UpdatedAt

	return cancel_lead, nil
}

func (ser *sale_service) recover_ser() {
	if err := recover(); err != nil {
		log.Println("RECOVERD : ", err)
	}
}

func (ser *sale_service) Fetch_users_data(user_id string) (interface{}, error) {
	defer ser.recover_ser()
	token := ser.jwt_service.GenerateTempToken(user_id, "EMP", "SALES")

	api_resquest := proxycalls.NewAPIRequest("fetch-user", "GET", false, nil, nil, map[string]string{
		"Authorization": token,
	})

	response, err := api_resquest.MakeRequest()

	if err != nil {
		return nil, err
	}

	defer response.Body.Close()
	if response.StatusCode == http.StatusOK {
		var response_data map[string]interface{}
		err = json.NewDecoder(response.Body).Decode(&response_data)

		return response_data["user_data"], err
	}

	return nil, nil
}
