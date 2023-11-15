package services

import (
	"database/sql"
	"strconv"

	"github.com/aniket-skroman/skroman_sales_service.git/apis/dto"
	"github.com/aniket-skroman/skroman_sales_service.git/apis/helper"
	"github.com/aniket-skroman/skroman_sales_service.git/apis/repositories"
	db "github.com/aniket-skroman/skroman_sales_service.git/sqlc_lib"
	"github.com/google/uuid"
)

type LeadInfoService interface {
	CreateLeadInfo(req dto.CreateLeadInfoRequestDTO) (dto.GetLeadInfoDTO, error)
}

type lead_info_service struct {
	lead_info_repo repositories.LeadInfoRepository
}

func NewLeadInfoService(repo repositories.LeadInfoRepository) LeadInfoService {
	return &lead_info_service{
		lead_info_repo: repo,
	}
}

func (ser *lead_info_service) CreateLeadInfo(req dto.CreateLeadInfoRequestDTO) (dto.GetLeadInfoDTO, error) {
	// check  for valid lead id
	lead_obj_id, err := uuid.Parse(req.LeadID)

	if err != nil {
		return dto.GetLeadInfoDTO{}, helper.ERR_INVALID_ID
	}

	// check for valid contact number
	_, err = strconv.Atoi(req.Contact)

	if err != nil {
		return dto.GetLeadInfoDTO{}, helper.ERR_REQUIRED_PARAMS
	}

	args := db.CreateLeadInfoParams{
		LeadID:       uuid.NullUUID{UUID: lead_obj_id, Valid: true},
		Name:         req.Name,
		Email:        sql.NullString{String: req.Email, Valid: true},
		Contact:      req.Contact,
		AddressLine1: sql.NullString{String: req.AddressLine1, Valid: true},
		City:         sql.NullString{String: req.City, Valid: true},
		State:        sql.NullString{String: req.State, Valid: true},
		LeadType:     sql.NullString{String: req.LeadType, Valid: true},
	}

	lead_info, err := ser.lead_info_repo.CreateLeadInfo(args)
	err = helper.Handle_db_err(err)

	if err != nil {
		return dto.GetLeadInfoDTO{}, err
	}

	result := new(dto.GetLeadInfoDTO).MakeGetLeadInfo(lead_info)

	return result.(dto.GetLeadInfoDTO), nil
}
