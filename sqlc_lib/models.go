// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.23.0

package db

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type LeadInfo struct {
	ID           uuid.UUID      `json:"id"`
	LeadID       uuid.NullUUID  `json:"lead_id"`
	Name         string         `json:"name"`
	Email        sql.NullString `json:"email"`
	Contact      string         `json:"contact"`
	AddressLine1 sql.NullString `json:"address_line_1"`
	City         sql.NullString `json:"city"`
	State        sql.NullString `json:"state"`
	LeadType     sql.NullString `json:"lead_type"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
}

type LeadOrder struct {
	ID          uuid.UUID      `json:"id"`
	LeadID      uuid.NullUUID  `json:"lead_id"`
	DeviceType  sql.NullString `json:"device_type"`
	DeviceModel sql.NullString `json:"device_model"`
	DeviceName  sql.NullString `json:"device_name"`
	DevicePrice sql.NullInt32  `json:"device_price"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
}

type SaleLeads struct {
	ID             uuid.UUID     `json:"id"`
	LeadBy         uuid.UUID     `json:"lead_by"`
	ReferalName    string        `json:"referal_name"`
	ReferalContact string        `json:"referal_contact"`
	Status         string        `json:"status"`
	QuatationCount sql.NullInt32 `json:"quatation_count"`
	CreatedAt      time.Time     `json:"created_at"`
	UpdatedAt      time.Time     `json:"updated_at"`
	IsLeadInfo     sql.NullBool  `json:"is_lead_info"`
	IsOrderInfo    sql.NullBool  `json:"is_order_info"`
}
