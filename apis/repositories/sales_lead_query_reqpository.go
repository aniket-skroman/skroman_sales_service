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

func (repo *sale_repo) FetchLeadCounts() (db.FetchLeadCountsRow, error) {
	ctx, cancel := repo.Init()
	defer cancel()

	return repo.db.Queries.FetchLeadCounts(ctx)
}

func (repo *sale_repo) FetchLeadCountMonthWise() ([]db.FetchLeadCountByMonthRow, error) {
	ctx, cancel := repo.Init()
	defer cancel()

	return repo.db.Queries.FetchLeadCountByMonth(ctx)
}

func (repo *sale_repo) FetchLeadsByStatus(args db.FetchLeadsByStatusParams) ([]db.FetchLeadsByStatusRow, error) {
	ctx, cancel := repo.Init()
	defer cancel()

	return repo.db.Queries.FetchLeadsByStatus(ctx, args)
}

func (repo *sale_repo) FetchPGCountLeadsByStatus(status string) (int64, error) {
	ctx, cancel := repo.Init()
	defer cancel()

	return repo.db.Queries.PGCountByLeadStatus(ctx, status)
}

func (repo *sale_repo) CancelLead(args db.CreateCancelLeadParams) error {

	// db tx
	tx, err := repo.db.DBTransaction()

	if err != nil {
		return err
	}

	txq := repo.db.WithTx(tx)

	ctx, cancel := repo.Init()
	defer cancel()

	// first create cancel lead obj
	_, err = txq.CreateCancelLead(ctx, args)

	if err != nil {
		return err
	}

	// then update a lead status
	up_args := db.UpdateLeadStatusParams{
		ID:     args.LeadID,
		Status: "CANCELD",
	}

	_, err = txq.UpdateLeadStatus(ctx, up_args)

	if err != nil {
		if err := tx.Rollback(); err != nil {
			return err
		}
		return err
	}

	return tx.Commit()
}

func (repo *sale_repo) FetchCancelLead(lead_id uuid.UUID) (db.CancelLeads, error) {
	ctx, cancel := repo.Init()
	defer cancel()

	return repo.db.Queries.FetchCancelLeadByLeadId(ctx, lead_id)
}
