package services

import (
	"database/sql"
	"reflect"
	"strconv"

	"github.com/aniket-skroman/skroman_sales_service.git/apis/dto"
	"github.com/aniket-skroman/skroman_sales_service.git/apis/helper"
	"github.com/aniket-skroman/skroman_sales_service.git/apis/repositories"
	db "github.com/aniket-skroman/skroman_sales_service.git/sqlc_lib"
	"github.com/google/uuid"
)

type LeadInfoService interface {
	CreateLeadInfo(req dto.CreateLeadInfoRequestDTO) (dto.GetLeadInfoDTO, error)
	FetchLeadInfoByLeadID(lead_id string) (dto.GetLeadInfoDTO, error)
	UpdateLeadInfo(req dto.UpdateLeadInfoRequestDTO, lead_info_id string) (dto.GetLeadInfoDTO, error)
	DeleteLeadInfo(lead_id string) error
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

func (ser *lead_info_service) FetchLeadInfoByLeadID(lead_id string) (dto.GetLeadInfoDTO, error) {
	lead_obj_id, err := uuid.Parse(lead_id)

	if err != nil {
		return dto.GetLeadInfoDTO{}, helper.ERR_INVALID_ID
	}

	args := uuid.NullUUID{UUID: lead_obj_id, Valid: true}

	result, err := ser.lead_info_repo.FetchLeadInfoByLead(args)

	if err != nil {
		return dto.GetLeadInfoDTO{}, err
	}

	// check if result is not null
	if (reflect.DeepEqual(result, db.LeadInfo{})) {
		return dto.GetLeadInfoDTO{}, helper.Err_Data_Not_Found
	}

	return new(dto.GetLeadInfoDTO).MakeGetLeadInfo(result).(dto.GetLeadInfoDTO), nil
}

func (ser *lead_info_service) UpdateLeadInfo(req dto.UpdateLeadInfoRequestDTO, lead_info_id string) (dto.GetLeadInfoDTO, error) {
	info_obj_id, err := uuid.Parse(lead_info_id)

	if err != nil {
		return dto.GetLeadInfoDTO{}, helper.ERR_INVALID_ID
	}

	args := db.UpdateLeadInfoParams{
		ID:           info_obj_id,
		Name:         req.Name,
		Email:        sql.NullString{String: req.Email, Valid: true},
		Contact:      req.Contact,
		AddressLine1: sql.NullString{String: req.AddressLine1, Valid: true},
		City:         sql.NullString{String: req.City, Valid: true},
		State:        sql.NullString{String: req.State, Valid: true},
		LeadType:     sql.NullString{String: req.LeadType, Valid: true},
	}

	result, err := ser.lead_info_repo.UpdateLeadInfo(args)

	err = helper.Handle_db_err(err)

	if err != nil {
		return dto.GetLeadInfoDTO{}, err
	}

	// check for result should not be empty struct
	if (reflect.DeepEqual(result, db.LeadInfo{})) {
		return dto.GetLeadInfoDTO{}, helper.Err_Update_Failed
	}

	return new(dto.GetLeadInfoDTO).MakeGetLeadInfo(result).(dto.GetLeadInfoDTO), nil
}

func (ser *lead_info_service) DeleteLeadInfo(lead_id string) error {
	leat_obj_id, err := uuid.Parse(lead_id)

	if err != nil {
		return helper.ERR_INVALID_ID
	}

	args := uuid.NullUUID{UUID: leat_obj_id, Valid: true}

	result, err := ser.lead_info_repo.DeleteLeadInfo(args)

	if err != nil {
		return err
	}

	affected_row, _ := result.RowsAffected()

	if affected_row == 0 {
		return helper.Err_Delete_Failed
	}

	return nil
}
