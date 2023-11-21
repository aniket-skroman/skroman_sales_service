package dto

import (
	"mime/multipart"
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
	Quantity    int32  `json:"quantity" binding:"required"`
}

type DeleteLeadOrderRequestDTO struct {
	OrderId string `form:"order_id" binding:"required"`
	LeadId  string `form:"lead_id" binding:"required"`
}

type UpdateLeadOrderRequestDTO struct {
	LeadID      string `json:"lead_id" binding:"required"`
	DeviceType  string `json:"device_type" binding:"required"`
	DeviceModel string `json:"device_model" binding:"required"`
	DevicePrice int32  `json:"device_price" binding:"required"`
	Quantity    int32  `json:"quantity" binding:"required"`
	DeviceName  string `json:"device_name"`
}

// struct will be used only for internal, this will not communicate to request
type UploadOrderQuatationRequestDTO struct {
	LeadId        uuid.UUID
	GeneratedBy   uuid.UUID
	QuatationLink string
	QuatationFile multipart.File
	FileHandler   multipart.FileHeader
}

// delete order quotation
type DeleteOrderQuotationRequestDTO struct {
	LeadId      string `form:"lead_id" binding:"required"`
	QuotationId string `form:"quotation_id" binding:"required"`
}

type LeadOrderDTO struct {
	ID          uuid.UUID     `json:"id"`
	LeadID      uuid.NullUUID `json:"lead_id"`
	DeviceType  string        `json:"device_type"`
	DeviceModel string        `json:"device_model"`
	DeviceName  string        `json:"device_name"`
	DevicePrice int32         `json:"device_price"`
	Quantity    int32         `json:"quantity"`
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
			Quantity:    module_data[0].Quantity.Int32,
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
			Quantity:    module_data[i].Quantity.Int32,
			CreatedAt:   module_data[i].CreatedAt,
			UpdatedAt:   module_data[i].UpdatedAt,
		}
	}
	return orders
}
