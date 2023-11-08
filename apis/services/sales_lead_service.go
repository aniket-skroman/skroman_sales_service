package services

import (
	"database/sql"
	"errors"
	"fmt"
	"strconv"

	"github.com/aniket-skroman/skroman_sales_service.git/apis/dto"
	"github.com/aniket-skroman/skroman_sales_service.git/apis/repositories"
	db "github.com/aniket-skroman/skroman_sales_service.git/sqlc_lib"
	"github.com/google/uuid"
)

type SalesLeadService interface {
	CreateNewLead(req dto.CreateNewLeadDTO) (db.SaleLeads, error)
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

	if err != nil {
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
