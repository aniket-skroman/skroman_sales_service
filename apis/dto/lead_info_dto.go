package dto

import (
	"time"

	db "github.com/aniket-skroman/skroman_sales_service.git/sqlc_lib"
	"github.com/google/uuid"
)

type CreateLeadInfoRequestDTO struct {
	LeadID       string `json:"lead_id" binding:"required"`
	Name         string `json:"name" binding:"required"`
	Email        string `json:"email" binding:"required,email"`
	Contact      string `json:"contact" binding:"required,min=10"`
	AddressLine1 string `json:"address_line_1" binding:"required"`
	City         string `json:"city" binding:"required"`
	State        string `json:"state" binding:"required"`
	LeadType     string `json:"lead_type" binding:"required"`
}

type GetLeadInfoDTO struct {
	ID           uuid.UUID `json:"id"`
	LeadID       uuid.UUID `json:"lead_id"`
	Name         string    `json:"name"`
	Email        string    `json:"email"`
	Contact      string    `json:"contact"`
	AddressLine1 string    `json:"address_line_1"`
	City         string    `json:"city"`
	State        string    `json:"state"`
	LeadType     string    `json:"lead_type"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

func (lead_info *GetLeadInfoDTO) MakeGetLeadInfo(module_data ...db.LeadInfo) interface{} {
	if len(module_data) == 1 {
		return GetLeadInfoDTO{
			ID:           module_data[0].ID,
			LeadID:       module_data[0].LeadID.UUID,
			Name:         module_data[0].Name,
			Email:        module_data[0].Email.String,
			Contact:      module_data[0].Contact,
			AddressLine1: module_data[0].AddressLine1.String,
			City:         module_data[0].City.String,
			State:        module_data[0].State.String,
			LeadType:     module_data[0].LeadType.String,
			CreatedAt:    module_data[0].CreatedAt,
			UpdatedAt:    module_data[0].UpdatedAt,
		}
	}

	new_leads := make([]GetLeadInfoDTO, len(module_data))

	for i := range module_data {
		new_leads[i] = GetLeadInfoDTO{
			ID:           module_data[i].ID,
			LeadID:       module_data[i].LeadID.UUID,
			Name:         module_data[i].Name,
			Email:        module_data[i].Email.String,
			Contact:      module_data[i].Contact,
			AddressLine1: module_data[i].AddressLine1.String,
			City:         module_data[i].City.String,
			State:        module_data[i].State.String,
			LeadType:     module_data[i].LeadType.String,
			CreatedAt:    module_data[i].CreatedAt,
			UpdatedAt:    module_data[i].UpdatedAt,
		}
	}
	return new_leads
}
