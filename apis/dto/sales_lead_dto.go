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
	PageId   int    `uri:"page_id"`
	PageSize int    `uri:"page_size"`
	Status   string `uri:"status" binding:"required,oneof=INIT PLACED CANCEL"`
}

type FetchLeadCountsDTO struct {
	LeadCount       db.FetchLeadCountsRow         `json:"lead_count"`
	LeadMonthCounts []db.FetchLeadCountByMonthRow `json:"lead_month_counts"`
}

type SaleLeadsDetailsDTO struct {
	LeadID         uuid.UUID      `json:"lead_id"`
	LeadBy         uuid.UUID      `json:"lead_by"`
	ReferalName    string         `json:"referal_name"`
	ReferalContact string         `json:"referal_contact"`
	Status         string         `json:"status"`
	LeadCreatedAt  time.Time      `json:"lead_created_at"`
	LeadUpdatedAt  time.Time      `json:"lead_updated_at"`
	QuatationCount int32          `json:"quatation_count"`
	LeadInfo       GetLeadInfoDTO `json:"lead_info"`
}

type OrderQuatation struct {
	ID            uuid.UUID `json:"id"`
	QuotationLink string    `json:"quotation_link"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

type SaleLeadsDTO struct {
	ID              uuid.UUID         `json:"id"`
	LeadBy          uuid.UUID         `json:"lead_by"`
	ReferalName     string            `json:"referal_name"`
	ReferalContact  string            `json:"referal_contact"`
	Status          string            `json:"status"`
	QuatationCount  int32             `json:"quatation_count"`
	CreatedAt       time.Time         `json:"created_at"`
	UpdatedAt       time.Time         `json:"updated_at"`
	IsLeadInfo      bool              `json:"is_lead_info"`
	IsOrderInfo     bool              `json:"is_order_info"`
	LeadInfo        *GetLeadInfoDTO   `json:"lead_info,omitempty"`
	LeadOrders      *[]LeadOrderDTO   `json:"lead_orders,omitempty"`
	OrderQuatations *[]OrderQuatation `json:"order_quotation,omitempty"`
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
			IsLeadInfo:     modeule_data[0].IsLeadInfo.Bool,
			IsOrderInfo:    modeule_data[0].IsOrderInfo.Bool,
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
			IsLeadInfo:     modeule_data[i].IsLeadInfo.Bool,
			IsOrderInfo:    modeule_data[i].IsOrderInfo.Bool,
		}
	}
	return sales_leads
}
