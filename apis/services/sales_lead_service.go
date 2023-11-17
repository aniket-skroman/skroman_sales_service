package services

import (
	"database/sql"
	"errors"
	"reflect"
	"strconv"
	"sync"

	"github.com/aniket-skroman/skroman_sales_service.git/apis/dto"
	"github.com/aniket-skroman/skroman_sales_service.git/apis/helper"
	"github.com/aniket-skroman/skroman_sales_service.git/apis/repositories"
	db "github.com/aniket-skroman/skroman_sales_service.git/sqlc_lib"
	"github.com/google/uuid"
)

type SalesLeadService interface {
	CreateNewLead(req dto.CreateNewLeadDTO) (db.SaleLeads, error)
	FetchAllLeads(req dto.FetchAllLeadsRequestDTO) (interface{}, error)
	FetchLeadByLeadId(lead_id uuid.UUID) (interface{}, error)
	IncreaeQuatationCount(lead_id uuid.UUID) error
}

type sale_service struct {
	sale_lead_repo  repositories.SalesRepository
	lead_order_serv LeadOrderService
}

func NewSalesLeadService(sale_lead_repo repositories.SalesRepository, lead_order_serv LeadOrderService) SalesLeadService {
	return &sale_service{
		sale_lead_repo:  sale_lead_repo,
		lead_order_serv: lead_order_serv,
	}
}

func (ser *sale_service) CreateNewLead(req dto.CreateNewLeadDTO) (db.SaleLeads, error) {
	lead_by, err := uuid.Parse(req.LeadBy)

	if err != nil {
		return db.SaleLeads{}, err
	}

	_, err = strconv.Atoi(req.ReferalContact)

	if err != nil || len(req.ReferalContact) != 10 {
		return db.SaleLeads{}, helper.ERR_REQUIRED_PARAMS
	}

	args := db.CreateNewLeadParams{
		LeadBy:         lead_by,
		ReferalName:    req.ReferalName,
		ReferalContact: req.ReferalContact,
		QuatationCount: sql.NullInt32{Int32: 1, Valid: true},
		Status:         "INIT",
	}

	new_lead, err := ser.sale_lead_repo.CreateSalesLead(args)

	return new_lead, err
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
		count, _ := ser.sale_lead_repo.CountSalesLead()
		helper.SetPaginationData(req.PageId, count)
	}()

	wg.Wait()

	if err != nil {
		return nil, err
	}
	if len(result) == 0 {
		return nil, sql.ErrNoRows
	}

	return sales_lead, err
	// data := new(dto.SaleLeadsDTO).MakeSaleLeadsDTO(result...)
	// if _, ok := data.(dto.SaleLeadsDTO); ok {
	// 	return []dto.SaleLeadsDTO{data.(dto.SaleLeadsDTO)}, nil
	// }

	//return data, nil
}

func (ser *sale_service) FetchLeadByLeadId(lead_id uuid.UUID) (interface{}, error) {

	wg := sync.WaitGroup{}
	wg.Add(2)
	var lead db.FetchLeadByLeadIdRow
	result := dto.SaleLeadsDTO{}
	var err error
	var orders []dto.LeadOrderDTO

	go func() {
		defer wg.Done()
		lead, err = ser.sale_lead_repo.FetchLeadByLeadId(lead_id)

		if err != nil {
			return
		}

		if (reflect.DeepEqual(lead, db.SaleLeads{})) {
			err = errors.New("lead not found")
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

		result.LeadInfo = &dto.GetLeadInfoDTO{
			ID:           lead.LeadInfoID,
			LeadID:       lead_id,
			Name:         lead.Name,
			Email:        lead.Email.String,
			Contact:      lead.Contact,
			AddressLine1: lead.AddressLine1.String,
			City:         lead.City.String,
			State:        lead.State.String,
			LeadType:     lead.LeadType.String,
			CreatedAt:    lead.LeadInfoCreatedAt,
			UpdatedAt:    lead.LeadInfoUpdatedAt,
		}

	}()

	go func() {
		defer wg.Done()
		orders, err = ser.lead_order_serv.FetchOrdersByLeadId(lead_id)
		if err != nil {
			result.LeadOrders = &[]dto.LeadOrderDTO{}
		} else {
			result.LeadOrders = &orders
		}
		result.LeadOrders = &orders
	}()

	wg.Wait()

	if err != nil {
		return dto.SaleLeadsDTO{}, err
	}

	return result, nil
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
