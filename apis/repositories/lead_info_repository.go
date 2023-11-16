package repositories

import (
	"context"
	"database/sql"
	"time"

	"github.com/aniket-skroman/skroman_sales_service.git/apis"
	db "github.com/aniket-skroman/skroman_sales_service.git/sqlc_lib"
	"github.com/google/uuid"
)

type LeadInfoRepository interface {
	Init() (context.Context, context.CancelFunc)
	CreateLeadInfo(args db.CreateLeadInfoParams) (db.LeadInfo, error)
	FetchLeadInfoByLead(lead_id uuid.NullUUID) (db.LeadInfo, error)
	UpdateLeadInfo(args db.UpdateLeadInfoParams) (db.LeadInfo, error)
	DeleteLeadInfo(lead_id uuid.NullUUID) (sql.Result, error)
}

type lead_info_repo struct {
	db *apis.Store
}

func NewLeadInfoRepository(db *apis.Store) LeadInfoRepository {
	return &lead_info_repo{
		db: db,
	}
}

func (repo *lead_info_repo) Init() (context.Context, context.CancelFunc) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	return ctx, cancel
}
