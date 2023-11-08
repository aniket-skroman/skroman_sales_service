package repositories

import (
	"context"
	"time"

	"github.com/aniket-skroman/skroman_sales_service.git/apis"
	db "github.com/aniket-skroman/skroman_sales_service.git/sqlc_lib"
	"github.com/google/uuid"
)

type SalesRepository interface {
	Init() (context.Context, context.CancelFunc)
	CreateSalesLead(args db.CreateNewLeadParams) (db.SaleLeads, error)
	UpdateSalesLeadRef(args db.UpdateSaleLeadReferalParams) (db.SaleLeads, error)
	IncreaseQuatationCount(lead_id uuid.UUID) (int64, error)
	FetchAllLeads(args db.FetchAllLeadsParams) ([]db.SaleLeads, error)
	FetchLeadByLeadId(lead_id uuid.UUID) (db.SaleLeads, error)
}

type sale_repo struct {
	db *apis.Store
}

func NewSalesRepository(db *apis.Store) SalesRepository {
	return &sale_repo{
		db: db,
	}
}

func (repo *sale_repo) Init() (context.Context, context.CancelFunc) {
	ctx, cance := context.WithTimeout(context.Background(), 5*time.Second)
	return ctx, cance
}
