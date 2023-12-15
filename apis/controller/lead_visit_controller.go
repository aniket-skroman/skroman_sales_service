package controller

import (
	"net/http"

	"github.com/aniket-skroman/skroman_sales_service.git/apis/dto"
	"github.com/aniket-skroman/skroman_sales_service.git/apis/helper"
	"github.com/aniket-skroman/skroman_sales_service.git/apis/services"
	"github.com/aniket-skroman/skroman_sales_service.git/utils"
	"github.com/gin-gonic/gin"
)

type LeadVisitController interface {
	CreateLeadVisit(ctx *gin.Context)
	FetchAllVisitByLead(ctx *gin.Context)
}

type lead_visit_cont struct {
	lead_visit_ser services.LeadVisitService
	response       map[string]interface{}
}

func NewLeadVisitService(serv services.LeadVisitService) LeadVisitController {
	return &lead_visit_cont{
		lead_visit_ser: serv,
		response:       map[string]interface{}{},
	}
}

func (cont *lead_visit_cont) CreateLeadVisit(ctx *gin.Context) {
	var req dto.CreateLeadVisitRequestDTO

	if err := ctx.ShouldBindJSON(&req); err != nil {
		cont.response = utils.BuildFailedResponse(helper.Handle_required_param_error(err))
		ctx.JSON(http.StatusBadRequest, cont.response)
		return
	}

	result, err := cont.lead_visit_ser.CreateLeadVisit(req)

	if err != nil {
		cont.response = utils.BuildFailedResponse(err.Error())

		if err == helper.ERR_INVALID_ID {
			ctx.JSON(http.StatusBadRequest, cont.response)
			return
		}

		ctx.JSON(http.StatusInternalServerError, cont.response)
		return
	}

	cont.response = utils.BuildSuccessResponse(utils.DATA_INSERTED, utils.SALES_LEAD, result)
	ctx.JSON(http.StatusCreated, cont.response)
}

func (cont *lead_visit_cont) FetchAllVisitByLead(ctx *gin.Context) {
	lead_id := ctx.Param("lead_id")

	result, err := cont.lead_visit_ser.FetchAllVisitByLead(lead_id)

	if err != nil {
		cont.response = utils.BuildFailedResponse(err.Error())

		if err == helper.ERR_INVALID_ID {
			ctx.JSON(http.StatusBadRequest, cont.response)
			return
		} else if err == helper.Err_Data_Not_Found {
			ctx.JSON(http.StatusNotFound, cont.response)
			return
		}

		ctx.JSON(http.StatusInternalServerError, cont.response)
		return
	}

	cont.response = utils.BuildSuccessResponse(utils.FETCHED_SUCCESS, utils.SALES_LEAD, result)
	ctx.JSON(http.StatusOK, cont.response)
}
