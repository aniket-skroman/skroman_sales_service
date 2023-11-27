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
	FetchAllLeads(args db.FetchAllLeadsParams) ([]db.FetchAllLeadsRow, error)
	FetchLeadByLeadId(lead_id uuid.UUID) (db.FetchLeadByLeadIdRow, error)
	IncreaeQuatationCount(lead_id uuid.UUID) (int64, error)
	CountSalesLead() (int64, error)
	FetchLeadCounts() (db.FetchLeadCountsRow, error)
	FetchLeadCountMonthWise() ([]db.FetchLeadCountByMonthRow, error)
	FetchLeadsByStatus(args db.FetchLeadsByStatusParams) ([]db.FetchLeadsByStatusRow, error)
	FetchPGCountLeadsByStatus(status string) (int64, error)
	CancelLead(args db.CreateCancelLeadParams) error
	FetchCancelLead(lead_id uuid.UUID) (db.CancelLeads, error)
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
