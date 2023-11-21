package repositories

import (
	"database/sql"
	"errors"

	db "github.com/aniket-skroman/skroman_sales_service.git/sqlc_lib"
	"github.com/google/uuid"
)

// create a db transaction for updating flags
func (repo *lead_order_repo) CreateLeadOrder(args db.CreateLeadOrderParams) (db.LeadOrder, error) {
	tx, err := repo.db.DBTransaction()

	if err != nil {
		return db.LeadOrder{}, err
	}

	txq := repo.db.WithTx(tx)

	ctx, cancel := repo.Init()
	defer cancel()

	// create order info first
	result, err := txq.CreateLeadOrder(ctx, args)

	if err != nil {
		return db.LeadOrder{}, err
	}

	// then update a is_order_info flag as true
	flag_args := db.UpdateIsLeadOrderParams{
		ID:          result.LeadID.UUID,
		IsOrderInfo: sql.NullBool{Bool: true, Valid: true},
	}

	_, err = txq.UpdateIsLeadOrder(ctx, flag_args)

	if err != nil {
		if r_err := tx.Rollback(); r_err != nil {
			return db.LeadOrder{}, r_err
		}

		return db.LeadOrder{}, err
	}

	err = tx.Commit()

	return result, err
}

func (repo *lead_order_repo) FetchOrdersByLeadId(lead_id uuid.NullUUID) ([]db.LeadOrder, error) {
	ctx, cancel := repo.Init()
	defer cancel()

	return repo.db.Queries.FetchOrdersByLeadId(ctx, lead_id)
}

// db trans
func (repo *lead_order_repo) DeleteLeadOrder(args db.DeleteLeadOrderParams) (sql.Result, error) {

	tx, err := repo.db.DBTransaction()

	if err != nil {
		return nil, err
	}

	txq := repo.db.WithTx(tx)

	ctx, cancel := repo.Init()
	defer cancel()

	// first delete the lead order
	result, err := txq.DeleteLeadOrder(ctx, args)
	if err != nil {
		return nil, err
	}

	// check all order has been deleted or not
	count, err := txq.CheckLeadHasOrder(ctx, uuid.NullUUID{UUID: args.LeadID.UUID, Valid: true})

	if err != nil {
		if err := tx.Rollback(); err != nil {
			return nil, err
		}

		return nil, err
	}

	// if there is no order against lead then change the flag
	if count == 0 {
		flag_args := db.UpdateIsLeadOrderParams{
			ID:          args.LeadID.UUID,
			IsOrderInfo: sql.NullBool{Bool: false, Valid: true},
		}

		_, err = txq.UpdateIsLeadOrder(ctx, flag_args)

		if err != nil {
			if err := tx.Rollback(); err != nil {
				return nil, err
			}

			return nil, err
		}
	}

	err = tx.Commit()

	return result, err
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

func (repo *lead_order_repo) UploadOrderQuatation(args db.CreateNewOrderQuatationParams) error {
	// make a db transaction
	tx, err := repo.db.DBTransaction()

	if err != nil {
		return err
	}

	txq := repo.db.WithTx(tx)

	ctx, cancel := repo.Init()
	defer cancel()

	// create a new order quatation
	_, err = txq.CreateNewOrderQuatation(ctx, args)

	if err != nil {
		return err
	}

	// then update the quatation count
	result, err := txq.IncreaeQuatationCount(ctx, args.LeadID)

	if err != nil {
		if err := tx.Rollback(); err != nil {
			return err
		}

		return err
	}

	if result == 0 {
		err := tx.Rollback()
		if err != nil {
			return err
		}

		return errors.New("failed to upload qutation")
	}

	err = tx.Commit()

	return err
}

func (repo *lead_order_repo) FetchOrderQutationsByLeadId(lead_id uuid.UUID) ([]db.OrderQuatation, error) {
	ctx, cancel := repo.Init()
	defer cancel()

	return repo.db.Queries.FetchQuatationByLeadId(ctx, lead_id)
}
