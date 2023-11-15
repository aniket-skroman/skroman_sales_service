package dto

import (
	"time"

	db "github.com/aniket-skroman/skroman_sales_service.git/sqlc_lib"
	"github.com/google/uuid"
)

type CreateNewLeadDTO struct {
	LeadBy         string `json:"lead_by" binding:"required"`
	ReferalName    string `json:"referal_name" binding:"required"`
	ReferalContact string `json:"referal_contact" validate:"required,min=10"`
	Status         string `json:"status"`
	QuatationCount int    `json:"quatation_count"`
}

type FetchAllLeadsRequestDTO struct {
	PageId   int `uri:"page_id"`
	PageSize int `uri:"page_size"`
}

type SaleLeadsDTO struct {
	ID             uuid.UUID `json:"id"`
	LeadBy         uuid.UUID `json:"lead_by"`
	ReferalName    string    `json:"referal_name"`
	ReferalContact string    `json:"referal_contact"`
	Status         string    `json:"status"`
	QuatationCount int32     `json:"quatation_count"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

func (sale *SaleLeadsDTO) MakeSaleLeadsDTO(modeule_data ...db.SaleLeads) interface{} {
	if len(modeule_data) == 1 {
		return SaleLeadsDTO{
			ID:             modeule_data[0].ID,
			LeadBy:         modeule_data[0].LeadBy,
			ReferalName:    modeule_data[0].ReferalName,
			ReferalContact: modeule_data[0].ReferalContact,
			Status:         modeule_data[0].Status,
			QuatationCount: modeule_data[0].QuatationCount.Int32,
			CreatedAt:      modeule_data[0].CreatedAt,
			UpdatedAt:      modeule_data[0].UpdatedAt,
		}
	}
	sales_leads := make([]SaleLeadsDTO, len(modeule_data))

	for i := range modeule_data {
		sales_leads[i] = SaleLeadsDTO{
			ID:             modeule_data[i].ID,
			LeadBy:         modeule_data[i].LeadBy,
			ReferalName:    modeule_data[i].ReferalName,
			ReferalContact: modeule_data[i].ReferalContact,
			Status:         modeule_data[i].Status,
			QuatationCount: modeule_data[i].QuatationCount.Int32,
			CreatedAt:      modeule_data[i].CreatedAt,
			UpdatedAt:      modeule_data[i].UpdatedAt,
		}
	}
	return sales_leads
}
