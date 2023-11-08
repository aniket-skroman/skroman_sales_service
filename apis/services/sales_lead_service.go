package services

import (
	"database/sql"
	"errors"
	"fmt"
	"reflect"
	"strconv"

	"github.com/aniket-skroman/skroman_sales_service.git/apis/dto"
	"github.com/aniket-skroman/skroman_sales_service.git/apis/repositories"
	db "github.com/aniket-skroman/skroman_sales_service.git/sqlc_lib"
	"github.com/google/uuid"
)

type SalesLeadService interface {
	CreateNewLead(req dto.CreateNewLeadDTO) (db.SaleLeads, error)
	FetchAllLeads(req dto.FetchAllLeadsRequestDTO) ([]db.SaleLeads, error)
	FetchLeadByLeadId(lead_id string) (db.SaleLeads, error)
}

type sale_service struct {
	sale_lead_repo repositories.SalesRepository
}

func NewSalesLeadService(sale_lead_repo repositories.SalesRepository) SalesLeadService {
	return &sale_service{
		sale_lead_repo: sale_lead_repo,
	}
}

func (ser *sale_service) CreateNewLead(req dto.CreateNewLeadDTO) (db.SaleLeads, error) {
	lead_by, err := uuid.Parse(req.LeadBy)

	if err != nil {
		return db.SaleLeads{}, err
	}

	_, err = strconv.Atoi(req.ReferalContact)

	if err != nil || len(req.ReferalContact) != 10 {
		return db.SaleLeads{}, errors.New("invalid contact please check params")
	}

	args := db.CreateNewLeadParams{
		LeadBy:         lead_by,
		ReferalName:    req.ReferalName,
		ReferalContact: req.ReferalContact,
		QuatationCount: sql.NullInt32{Int32: 1, Valid: true},
		Status:         "INIT",
	}

	new_lead, err := ser.sale_lead_repo.CreateSalesLead(args)
	fmt.Println("Error from serv : ", err)
	return new_lead, err
}
func (ser *sale_service) FetchAllLeads(req dto.FetchAllLeadsRequestDTO) ([]db.SaleLeads, error) {

	args := db.FetchAllLeadsParams{
		Limit:  int32(req.PageSize),
		Offset: (int32(req.PageId) - 1) * int32(req.PageSize),
	}

	result, err := ser.sale_lead_repo.FetchAllLeads(args)

	if err != nil {
		return nil, err
	}
	if len(result) == 0 {
		return nil, sql.ErrNoRows
	}

	return result, nil
}

func (ser *sale_service) FetchLeadByLeadId(lead_id string) (db.SaleLeads, error) {
	lead_obj_id, err := uuid.Parse(lead_id)

	if err != nil {
		return db.SaleLeads{}, err
	}

	lead, err := ser.sale_lead_repo.FetchLeadByLeadId(lead_obj_id)

	if err != nil {
		return db.SaleLeads{}, err
	}

	if (reflect.DeepEqual(lead, db.SaleLeads{})) {
		return db.SaleLeads{}, errors.New("lead not found")
	}

	return lead, nil
}
