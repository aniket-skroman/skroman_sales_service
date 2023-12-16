package dto

import (
	"time"

	"github.com/google/uuid"
)

type CreateLeadVisitRequestDTO struct {
	LeadId          string `json:"lead_id" binding:"required"`
	VisitBy         string `json:"visit_by" binding:"required"`
	VisitDiscussion string `json:"discussion" binding:"required"`
}

type LeadVistDTO struct {
	ID              uuid.UUID   `json:"id"`
	LeadID          uuid.UUID   `json:"lead_id"`
	VisitBy         interface{} `json:"visit_by"`
	VisitDiscussion string      `json:"visit_discussion"`
	CreatedAt       time.Time   `json:"created_at"`
	UpdatedAt       time.Time   `json:"updated_at"`
}

type LeadVisit struct {
	LeadVistDTO []LeadVistDTO `json:"visit_data"`
	LeadInfo    interface{}   `json:"lead_info"`
}
