package dto

import (
	"fmt"
	"time"

	db "github.com/aniket-skroman/skroman_sales_service.git/sqlc_lib"
	"github.com/google/uuid"
)

type CreateLeadOrderRequestDTO struct {
	LeadID      string `json:"lead_id" binding:"required"`
	DeviceType  string `json:"device_type" binding:"required"`
	DeviceModel string `json:"device_model" binding:"required"`
	DevicePrice int32  `json:"device_price" binding:"required"`
	DeviceName  string `json:"device_name"`
}

type DeleteLeadOrderRequestDTO struct {
	OrderId string `form:"order_id" binding:"required"`
	LeadId  string `form:"lead_id" binding:"required"`
}

type UpdateLeadOrderRequestDTO struct {
	LeadID      string `json:"lead_id"`
	DeviceType  string `json:"device_type"`
	DeviceModel string `json:"device_model"`
	DevicePrice int32  `json:"device_price"`
}

type LeadOrderDTO struct {
	ID          uuid.UUID     `json:"id"`
	LeadID      uuid.NullUUID `json:"lead_id"`
	DeviceType  string        `json:"device_type"`
	DeviceModel string        `json:"device_model"`
	DeviceName  string        `json:"device_name"`
	DevicePrice int32         `json:"device_price"`
	CreatedAt   time.Time     `json:"created_at"`
	UpdatedAt   time.Time     `json:"updated_at"`
}

func (order *LeadOrderDTO) MakeLeadOrderDTO(module_data ...db.LeadOrder) interface{} {
	if len(module_data) == 1 {
		return LeadOrderDTO{
			ID:          module_data[0].ID,
			LeadID:      module_data[0].LeadID,
			DeviceType:  module_data[0].DeviceType.String,
			DeviceModel: module_data[0].DeviceModel.String,
			DeviceName:  module_data[0].DeviceName.String,
			DevicePrice: module_data[0].DevicePrice.Int32,
			CreatedAt:   module_data[0].CreatedAt,
			UpdatedAt:   module_data[0].UpdatedAt,
		}
	}

	orders := make([]LeadOrderDTO, len(module_data))

	for i := range module_data {
		orders[i] = LeadOrderDTO{
			ID:          module_data[i].ID,
			LeadID:      module_data[i].LeadID,
			DeviceType:  module_data[i].DeviceType.String,
			DeviceModel: module_data[i].DeviceModel.String,
			DeviceName:  module_data[i].DeviceName.String,
			DevicePrice: module_data[i].DevicePrice.Int32,
			CreatedAt:   module_data[i].CreatedAt,
			UpdatedAt:   module_data[i].UpdatedAt,
		}
	}
	fmt.Println("Returning orders....")
	return orders
}
