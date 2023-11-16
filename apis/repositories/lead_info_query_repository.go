package repositories

import (
	"database/sql"

	db "github.com/aniket-skroman/skroman_sales_service.git/sqlc_lib"
	"github.com/google/uuid"
)

func (repo *lead_info_repo) CreateLeadInfo(args db.CreateLeadInfoParams) (db.LeadInfo, error) {
	ctx, cancel := repo.Init()
	defer cancel()

	return repo.db.Queries.CreateLeadInfo(ctx, args)
}

func (repo *lead_info_repo) FetchLeadInfoByLead(lead_id uuid.NullUUID) (db.LeadInfo, error) {
	ctx, cancel := repo.Init()
	defer cancel()

	return repo.db.Queries.FetchLeadInfoByLeadID(ctx, lead_id)
}

func (repo *lead_info_repo) UpdateLeadInfo(args db.UpdateLeadInfoParams) (db.LeadInfo, error) {
	ctx, cancel := repo.Init()
	defer cancel()

	return repo.db.Queries.UpdateLeadInfo(ctx, args)
}

func (repo *lead_info_repo) DeleteLeadInfo(lead_id uuid.NullUUID) (sql.Result, error) {
	ctx, cancel := repo.Init()
	defer cancel()

	return repo.db.Queries.DeleteLeadInfoByLeadId(ctx, lead_id)
}
