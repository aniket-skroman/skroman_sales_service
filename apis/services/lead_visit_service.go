package services

import (
	"sync"

	"github.com/aniket-skroman/skroman_sales_service.git/apis/dto"
	"github.com/aniket-skroman/skroman_sales_service.git/apis/helper"
	"github.com/aniket-skroman/skroman_sales_service.git/apis/repositories"
	db "github.com/aniket-skroman/skroman_sales_service.git/sqlc_lib"
	"github.com/google/uuid"
)

type LeadVisitService interface {
	CreateLeadVisit(req dto.CreateLeadVisitRequestDTO) (db.LeadVisit, error)
	FetchAllVisitByLead(lead_id string) (dto.LeadVisit, error)
}

type lead_visit_serv struct {
	lead_visit_repo repositories.LeadVisitRepository
	lead_serv       SalesLeadService
	jwt_serv        JWTService
}

func NewLeadVisitService(repo repositories.LeadVisitRepository, lead_serv SalesLeadService,
	jwt_serv JWTService,
) LeadVisitService {
	return &lead_visit_serv{
		lead_visit_repo: repo,
		lead_serv:       lead_serv,
		jwt_serv:        jwt_serv,
	}
}

func (serv *lead_visit_serv) CreateLeadVisit(req dto.CreateLeadVisitRequestDTO) (db.LeadVisit, error) {
	lead_obj_id, err := uuid.Parse(req.LeadId)

	if err != nil {
		return db.LeadVisit{}, helper.ERR_INVALID_ID
	}

	visit_by_obj, err := uuid.Parse(req.VisitBy)

	if err != nil {
		return db.LeadVisit{}, helper.ERR_INVALID_ID
	}

	args := db.CreateLeadVisitParams{
		LeadID:          lead_obj_id,
		VisitBy:         visit_by_obj,
		VisitDiscussion: req.VisitDiscussion,
	}

	result, err := serv.lead_visit_repo.CreateLeadVisit(args)

	err = helper.Handle_db_err(err)

	if err != nil {
		return db.LeadVisit{}, err
	}

	return result, nil
}

func (serv *lead_visit_serv) FetchAllVisitByLead(lead_id string) (dto.LeadVisit, error) {
	lead_obj_id, err := uuid.Parse(lead_id)

	if err != nil {
		return dto.LeadVisit{}, helper.ERR_INVALID_ID
	}

	result, err := serv.lead_visit_repo.FetchAllVisitByLead(lead_obj_id)

	if err != nil {
		return dto.LeadVisit{}, err
	}

	if len(result) == 0 {
		return dto.LeadVisit{}, helper.Err_Data_Not_Found
	}

	var lead_visits dto.LeadVisit
	lead_visits.LeadVistDTO = make([]dto.LeadVistDTO, len(result))
	wg := sync.WaitGroup{}
	wg.Add(1)

	go func(lead_id uuid.UUID) {
		// defer wg.Done()
		lead_info, _ := serv.set_lead_info_data(&wg, lead_id)
		lead_visits.LeadInfo = lead_info
	}(lead_obj_id)

	for i, visit := range result {
		wg.Add(2)
		go serv.set_lead_visit_data(&wg, &lead_visits.LeadVistDTO[i], &visit)
		go serv.set_visit_by_user_data(&wg, &lead_visits.LeadVistDTO[i], visit.VisitBy)

	}

	wg.Wait()

	return lead_visits, nil
}

func (serv *lead_visit_serv) set_lead_visit_data(wg *sync.WaitGroup, result *dto.LeadVistDTO, data *db.LeadVisit) {
	defer wg.Done()
	result.ID = data.ID
	result.LeadID = data.LeadID
	result.VisitDiscussion = data.VisitDiscussion
	result.CreatedAt = data.CreatedAt
	result.UpdatedAt = data.UpdatedAt
}

func (serv *lead_visit_serv) set_visit_by_user_data(wg *sync.WaitGroup, result *dto.LeadVistDTO, visit_by uuid.UUID) {
	defer wg.Done()

	user_data, _ := serv.lead_serv.Fetch_users_data(visit_by.String())
	result.VisitBy = user_data
}

func (serv *lead_visit_serv) set_lead_info_data(wg *sync.WaitGroup, lead_id uuid.UUID) (interface{}, error) {
	defer wg.Done()
	lead_info, err := serv.lead_serv.FetchLeadByLeadId(lead_id)
	return lead_info, err
}
