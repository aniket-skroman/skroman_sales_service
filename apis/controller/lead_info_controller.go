package controller

import (
	"database/sql"
	"errors"
	"net/http"

	"github.com/aniket-skroman/skroman_sales_service.git/apis/dto"
	"github.com/aniket-skroman/skroman_sales_service.git/apis/helper"
	"github.com/aniket-skroman/skroman_sales_service.git/apis/services"
	"github.com/aniket-skroman/skroman_sales_service.git/utils"
	"github.com/gin-gonic/gin"
)

type LeadInfoController interface {
	CreateLeadInfo(*gin.Context)
	FetchLeadInfoByLeadID(*gin.Context)
	UpdateLeadInfo(*gin.Context)
	DeleteLeadInfo(*gin.Context)
}

type lead_info_cont struct {
	serv     services.LeadInfoService
	response map[string]interface{}
}

func NewLeadInfoController(service services.LeadInfoService) LeadInfoController {
	return &lead_info_cont{
		serv:     service,
		response: make(map[string]interface{}),
	}
}

func (cont *lead_info_cont) CreateLeadInfo(ctx *gin.Context) {
	var req dto.CreateLeadInfoRequestDTO

	if err := ctx.ShouldBindJSON(&req); err != nil {
		cont.response = utils.RequestParamsMissingResponse(helper.Handle_required_param_error(err))
		ctx.JSON(http.StatusBadRequest, cont.response)
		return
	}

	lead_info, err := cont.serv.CreateLeadInfo(req)

	if err != nil {
		if errors.Is(err, helper.ERR_INVALID_ID) {
			cont.response = utils.BuildFailedResponse(helper.ERR_INVALID_ID.Error())
			ctx.JSON(http.StatusForbidden, cont.response)
			return
		} else if errors.Is(err, helper.ERR_REQUIRED_PARAMS) {
			cont.response = utils.BuildFailedResponse("contact - " + helper.ERR_REQUIRED_PARAMS.Error())
			ctx.JSON(http.StatusForbidden, cont.response)
			return
		} else if errors.Is(err, helper.Err_Lead_Exists) {
			cont.response = utils.BuildFailedResponse(err.Error())
			ctx.JSON(http.StatusForbidden, cont.response)
			return
		}
		cont.response = utils.BuildFailedResponse(err.Error())
		ctx.JSON(http.StatusInternalServerError, cont.response)
		return
	}

	cont.response = utils.BuildSuccessResponse(utils.DATA_INSERTED, utils.SALES_LEAD, lead_info)
	ctx.JSON(http.StatusOK, cont.response)
}

func (cont *lead_info_cont) FetchLeadInfoByLeadID(ctx *gin.Context) {
	lead_id := ctx.Param("lead_id")

	result, err := cont.serv.FetchLeadInfoByLeadID(lead_id)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			cont.response = utils.BuildFailedResponse(helper.Err_Data_Not_Found.Error())
			ctx.JSON(http.StatusNotFound, cont.response)
			return
		} else if errors.Is(err, helper.ERR_INVALID_ID) {
			cont.response = utils.BuildFailedResponse(err.Error())
			ctx.JSON(http.StatusUnprocessableEntity, cont.response)
			return
		}

		cont.response = utils.BuildFailedResponse(err.Error())
		ctx.JSON(http.StatusInternalServerError, cont.response)
		return
	}

	cont.response = utils.BuildSuccessResponse(utils.FETCHED_SUCCESS, utils.SALES_LEAD, result)
	ctx.JSON(http.StatusOK, cont.response)
}

func (cont *lead_info_cont) UpdateLeadInfo(ctx *gin.Context) {
	lead_info_id := ctx.Param("lead_id")

	var req dto.UpdateLeadInfoRequestDTO

	if err := ctx.ShouldBindJSON(&req); err != nil {
		cont.response = utils.BuildFailedResponse(helper.Handle_required_param_error(err))
		ctx.JSON(http.StatusBadRequest, cont.response)
		return
	}

	result, err := cont.serv.UpdateLeadInfo(req, lead_info_id)

	if err != nil {
		cont.response = utils.BuildFailedResponse(err.Error())

		if errors.Is(err, helper.ERR_INVALID_ID) {
			ctx.JSON(http.StatusBadRequest, cont.response)
			return
		} else if errors.Is(err, helper.Err_Update_Failed) {
			ctx.JSON(http.StatusUnprocessableEntity, cont.response)
			return
		}
		ctx.JSON(http.StatusInternalServerError, cont.response)
		return
	}

	cont.response = utils.BuildSuccessResponse(utils.UPDATE_SUCCESS, utils.SALES_LEAD, result)
	ctx.JSON(http.StatusOK, cont.response)
}

func (cont *lead_info_cont) DeleteLeadInfo(ctx *gin.Context) {
	lead_id := ctx.Param("lead_id")

	err := cont.serv.DeleteLeadInfo(lead_id)

	if err != nil {
		cont.response = utils.BuildFailedResponse(err.Error())
		if errors.Is(err, helper.ERR_INVALID_ID) {
			ctx.JSON(http.StatusBadRequest, cont.response)
			return
		}
		ctx.JSON(http.StatusInternalServerError, cont.response)
		return
	}

	cont.response = utils.BuildSuccessResponse(utils.DELETE_SUCCESS, utils.SALES_LEAD, nil)
	ctx.JSON(http.StatusOK, cont.response)
}
