package repositories

import db "github.com/aniket-skroman/skroman_sales_service.git/sqlc_lib"

func (repo *lead_info_repo) CreateLeadInfo(args db.CreateLeadInfoParams) (db.LeadInfo, error) {
	ctx, cancel := repo.Init()
	defer cancel()

	return repo.db.Queries.CreateLeadInfo(ctx, args)
}
