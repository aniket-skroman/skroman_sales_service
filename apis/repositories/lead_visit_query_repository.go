package repositories

import (
	"database/sql"

	db "github.com/aniket-skroman/skroman_sales_service.git/sqlc_lib"
	"github.com/google/uuid"
)

func (repo *lead_visit_repo) CreateLeadVisit(args db.CreateLeadVisitParams) (db.LeadVisit, error) {
	ctx, cancel := repo.Init()
	defer cancel()

	return repo.db.Queries.CreateLeadVisit(ctx, args)
}

func (repo *lead_visit_repo) FetchAllVisitByLead(lead_id uuid.UUID) ([]db.LeadVisit, error) {
	ctx, cancel := repo.Init()
	defer cancel()

	return repo.db.Queries.FetchAllVisitByLead(ctx, lead_id)
}

func (repo *lead_visit_repo) DeleteLeadVisit(args db.DeleteLeadVisitParams) (sql.Result, error) {
	ctx, cancel := repo.Init()
	defer cancel()

	return repo.db.Queries.DeleteLeadVisit(ctx, args)
}
