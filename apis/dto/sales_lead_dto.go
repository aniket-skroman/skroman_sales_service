package dto

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
