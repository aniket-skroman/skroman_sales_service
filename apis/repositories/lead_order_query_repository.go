package repositories

import (
	"database/sql"

	db "github.com/aniket-skroman/skroman_sales_service.git/sqlc_lib"
	"github.com/google/uuid"
)

func (repo *lead_order_repo) CreateLeadOrder(args db.CreateLeadOrderParams) (db.LeadOrder, error) {
	ctx, cancel := repo.Init()
	defer cancel()

	return repo.db.Queries.CreateLeadOrder(ctx, args)
}

func (repo *lead_order_repo) FetchOrdersByLeadId(lead_id uuid.NullUUID) ([]db.LeadOrder, error) {
	ctx, cancel := repo.Init()
	defer cancel()

	return repo.db.Queries.FetchOrdersByLeadId(ctx, lead_id)
}

func (repo *lead_order_repo) DeleteLeadOrder(args db.DeleteLeadOrderParams) (sql.Result, error) {
	ctx, cancel := repo.Init()
	defer cancel()

	return repo.db.Queries.DeleteLeadOrder(ctx, args)
}

func (repo *lead_order_repo) UpdateLeadOrder(args db.UpdateLeadOrderParams) (db.LeadOrder, error) {
	ctx, cancel := repo.Init()
	defer cancel()

	return repo.db.Queries.UpdateLeadOrder(ctx, args)
}

func (repo *lead_order_repo) FetchOrdersByOrderId(order_id uuid.UUID) (db.LeadOrder, error) {
	ctx, cancel := repo.Init()
	defer cancel()

	return repo.db.Queries.FetchLeadOrderByOrderId(ctx, order_id)
}
