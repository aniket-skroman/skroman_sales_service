// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.23.0
// source: lead_info.sql

package db

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
)

const createLeadInfo = `-- name: CreateLeadInfo :one
insert into lead_info (
    lead_id,
    name,
    email,
    contact,
    address_line_1,
    city,
    state,
    lead_type
) values (
    $1,$2,$3,$4,$5,$6,$7,$8
) returning id, lead_id, name, email, contact, address_line_1, city, state, lead_type, created_at, updated_at
`

type CreateLeadInfoParams struct {
	LeadID       uuid.NullUUID  `json:"lead_id"`
	Name         string         `json:"name"`
	Email        sql.NullString `json:"email"`
	Contact      string         `json:"contact"`
	AddressLine1 sql.NullString `json:"address_line_1"`
	City         sql.NullString `json:"city"`
	State        sql.NullString `json:"state"`
	LeadType     sql.NullString `json:"lead_type"`
}

// create lead info
func (q *Queries) CreateLeadInfo(ctx context.Context, arg CreateLeadInfoParams) (LeadInfo, error) {
	row := q.db.QueryRowContext(ctx, createLeadInfo,
		arg.LeadID,
		arg.Name,
		arg.Email,
		arg.Contact,
		arg.AddressLine1,
		arg.City,
		arg.State,
		arg.LeadType,
	)
	var i LeadInfo
	err := row.Scan(
		&i.ID,
		&i.LeadID,
		&i.Name,
		&i.Email,
		&i.Contact,
		&i.AddressLine1,
		&i.City,
		&i.State,
		&i.LeadType,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deleteLeadInfoByLeadId = `-- name: DeleteLeadInfoByLeadId :execresult
delete from lead_info
where lead_id = $1
`

func (q *Queries) DeleteLeadInfoByLeadId(ctx context.Context, leadID uuid.NullUUID) (sql.Result, error) {
	return q.db.ExecContext(ctx, deleteLeadInfoByLeadId, leadID)
}

const fetchLeadInfoByLeadID = `-- name: FetchLeadInfoByLeadID :one
select id, lead_id, name, email, contact, address_line_1, city, state, lead_type, created_at, updated_at from lead_info
where lead_id = $1
`

func (q *Queries) FetchLeadInfoByLeadID(ctx context.Context, leadID uuid.NullUUID) (LeadInfo, error) {
	row := q.db.QueryRowContext(ctx, fetchLeadInfoByLeadID, leadID)
	var i LeadInfo
	err := row.Scan(
		&i.ID,
		&i.LeadID,
		&i.Name,
		&i.Email,
		&i.Contact,
		&i.AddressLine1,
		&i.City,
		&i.State,
		&i.LeadType,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const updateLeadInfo = `-- name: UpdateLeadInfo :one
update lead_info
set name = $2,
email = $3, contact=$4,
address_line_1=$5, city=$6,
state=$7, lead_type=$8,
updated_at = CURRENT_TIMESTAMP
where id = $1
returning id, lead_id, name, email, contact, address_line_1, city, state, lead_type, created_at, updated_at
`

type UpdateLeadInfoParams struct {
	ID           uuid.UUID      `json:"id"`
	Name         string         `json:"name"`
	Email        sql.NullString `json:"email"`
	Contact      string         `json:"contact"`
	AddressLine1 sql.NullString `json:"address_line_1"`
	City         sql.NullString `json:"city"`
	State        sql.NullString `json:"state"`
	LeadType     sql.NullString `json:"lead_type"`
}

func (q *Queries) UpdateLeadInfo(ctx context.Context, arg UpdateLeadInfoParams) (LeadInfo, error) {
	row := q.db.QueryRowContext(ctx, updateLeadInfo,
		arg.ID,
		arg.Name,
		arg.Email,
		arg.Contact,
		arg.AddressLine1,
		arg.City,
		arg.State,
		arg.LeadType,
	)
	var i LeadInfo
	err := row.Scan(
		&i.ID,
		&i.LeadID,
		&i.Name,
		&i.Email,
		&i.Contact,
		&i.AddressLine1,
		&i.City,
		&i.State,
		&i.LeadType,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
