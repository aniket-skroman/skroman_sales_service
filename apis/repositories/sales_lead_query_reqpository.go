package repositories

import (
	db "github.com/aniket-skroman/skroman_sales_service.git/sqlc_lib"
	"github.com/google/uuid"
)

func (repo *sale_repo) CreateSalesLead(args db.CreateNewLeadParams) (db.SaleLeads, error) {
	ctx, cancel := repo.Init()
	defer cancel()

	return repo.db.Queries.CreateNewLead(ctx, args)
}

func (repo *sale_repo) UpdateSalesLeadRef(args db.UpdateSaleLeadReferalParams) (db.SaleLeads, error) {
	ctx, cancel := repo.Init()
	defer cancel()

	return repo.db.Queries.UpdateSaleLeadReferal(ctx, args)
}

func (repo *sale_repo) IncreaseQuatationCount(lead_id uuid.UUID) (int64, error) {
	ctx, cancel := repo.Init()
	defer cancel()

	return repo.db.Queries.IncreaeQuatationCount(ctx, lead_id)
}

func (repo *sale_repo) FetchAllLeads(args db.FetchAllLeadsParams) ([]db.FetchAllLeadsRow, error) {
	ctx, cancel := repo.Init()
	defer cancel()

	return repo.db.Queries.FetchAllLeads(ctx, args)
}

func (repo *sale_repo) FetchLeadByLeadId(lead_id uuid.UUID) (db.FetchLeadByLeadIdRow, error) {
	ctx, cancel := repo.Init()
	defer cancel()

	return repo.db.Queries.FetchLeadByLeadId(ctx, lead_id)
}

func (repo *sale_repo) IncreaeQuatationCount(lead_id uuid.UUID) (int64, error) {
	ctx, cancel := repo.Init()
	defer cancel()

	return repo.db.Queries.IncreaeQuatationCount(ctx, lead_id)
}

func (repo *sale_repo) CountSalesLead() (int64, error) {
	ctx, cancel := repo.Init()
	defer cancel()

	return repo.db.Queries.CountOfLeads(ctx)
}
