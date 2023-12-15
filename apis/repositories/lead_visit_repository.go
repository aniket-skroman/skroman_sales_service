package repositories

import (
	"context"
	"database/sql"
	"time"

	"github.com/aniket-skroman/skroman_sales_service.git/apis"
	db "github.com/aniket-skroman/skroman_sales_service.git/sqlc_lib"
	"github.com/google/uuid"
)

type LeadVisitRepository interface {
	Init() (context.Context, context.CancelFunc)
	CreateLeadVisit(args db.CreateLeadVisitParams) (db.LeadVisit, error)
	FetchAllVisitByLead(lead_id uuid.UUID) ([]db.LeadVisit, error)
	DeleteLeadVisit(args db.DeleteLeadVisitParams) (sql.Result, error)
}

type lead_visit_repo struct {
	db *apis.Store
}

func NewLeadVisitRepository(db *apis.Store) LeadVisitRepository {
	return &lead_visit_repo{
		db: db,
	}
}

func (repo *lead_visit_repo) Init() (context.Context, context.CancelFunc) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	return ctx, cancel
}
