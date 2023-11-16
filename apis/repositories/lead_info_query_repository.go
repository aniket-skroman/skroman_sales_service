package repositories

import (
	"database/sql"

	db "github.com/aniket-skroman/skroman_sales_service.git/sqlc_lib"
	"github.com/google/uuid"
)

// implement a DB transaction
func (repo *lead_info_repo) CreateLeadInfo(args db.CreateLeadInfoParams) (db.LeadInfo, error) {
	tx, err := repo.db.DBTransaction()

	if err != nil {
		return db.LeadInfo{}, err
	}

	ctx, cancel := repo.Init()
	defer cancel()

	txq := repo.db.WithTx(tx)

	// first create lead info
	result, err := txq.CreateLeadInfo(ctx, args)
	if err != nil {
		return db.LeadInfo{}, err
	}

	// then update a is_lead_info flag
	flag_args := db.UpdateIsLeadInfoParams{
		ID:         result.LeadID.UUID,
		IsLeadInfo: sql.NullBool{Bool: true, Valid: true},
	}

	_, err = txq.UpdateIsLeadInfo(ctx, flag_args)

	if err != nil {
		if err := tx.Rollback(); err != nil {
			return db.LeadInfo{}, err
		}

		return db.LeadInfo{}, err
	}

	err = tx.Commit()

	return result, err
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

// db transaction if lead info get deleted make flag flase
func (repo *lead_info_repo) DeleteLeadInfo(lead_id uuid.NullUUID) (sql.Result, error) {
	tx, err := repo.db.DBTransaction()

	if err != nil {
		return nil, err
	}

	txq := repo.db.WithTx(tx)

	ctx, cancel := repo.Init()
	defer cancel()

	// delete first lead info
	result, err := txq.DeleteLeadInfoByLeadId(ctx, lead_id)
	if err != nil {
		return nil, err
	}

	// then update flag in sales_lead
	flag_args := db.UpdateIsLeadInfoParams{
		ID:         lead_id.UUID,
		IsLeadInfo: sql.NullBool{Bool: false, Valid: true},
	}

	_, err = txq.UpdateIsLeadInfo(ctx, flag_args)

	if err != nil {
		if err := tx.Rollback(); err != nil {
			return nil, err
		}
		return nil, err
	}

	err = tx.Commit()

	return result, err
}
