package repositories

import (
	"context"
	"database/sql"
	"time"

	"github.com/aniket-skroman/skroman_sales_service.git/apis"
	db "github.com/aniket-skroman/skroman_sales_service.git/sqlc_lib"
	"github.com/google/uuid"
)

type LeadOrderRepository interface {
	Init() (context.Context, context.CancelFunc)
	CreateLeadOrder(args db.CreateLeadOrderParams) (db.LeadOrder, error)
	FetchOrdersByLeadId(lead_id uuid.NullUUID) ([]db.LeadOrder, error)
	DeleteLeadOrder(args db.DeleteLeadOrderParams) (sql.Result, error)
	UpdateLeadOrder(args db.UpdateLeadOrderParams) (db.LeadOrder, error)
	FetchOrdersByOrderId(order_id uuid.UUID) (db.LeadOrder, error)
	UploadOrderQuatation(args db.CreateNewOrderQuatationParams) error
	FetchOrderQutationsByLeadId(lead_id uuid.UUID) ([]db.OrderQuatation, error)
}

type lead_order_repo struct {
	db *apis.Store
}

func NewLeadOrderRepository(db *apis.Store) LeadOrderRepository {
	return &lead_order_repo{
		db: db,
	}
}

func (repo *lead_order_repo) Init() (context.Context, context.CancelFunc) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	return ctx, cancel
}
