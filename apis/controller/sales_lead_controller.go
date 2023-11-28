package controller

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/aniket-skroman/skroman_sales_service.git/apis/dto"
	"github.com/aniket-skroman/skroman_sales_service.git/apis/helper"
	"github.com/aniket-skroman/skroman_sales_service.git/apis/services"
	"github.com/aniket-skroman/skroman_sales_service.git/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type SaleLeadsController interface {
	CreateNewLead(*gin.Context)
	FetchAllLeads(*gin.Context)
	FetchLeadByLeadId(*gin.Context)
	FetchLeadCounts(*gin.Context)
	FetchLeadsByStatus(ctx *gin.Context)
	CancelLead(ctx *gin.Context)
	FetchCancelLead(ctx *gin.Context)
}

type sale_controller struct {
	sale_serv services.SalesLeadService
	response  map[string]interface{}
}

func NewSaleLeadsController(sale_serv services.SalesLeadService) SaleLeadsController {
	return &sale_controller{
		sale_serv: sale_serv,
		response:  make(map[string]interface{}),
	}
}

func (cont *sale_controller) CreateNewLead(ctx *gin.Context) {
	var req dto.CreateNewLeadDTO

	if err := ctx.ShouldBindJSON(&req); err != nil {
		fmt.Println("Errors", err)
		cont.response = utils.RequestParamsMissingResponse(helper.Handle_required_param_error(err))
		ctx.JSON(http.StatusBadRequest, cont.response)
		return
	}

	result, err := cont.sale_serv.CreateNewLead(req)
	if err != nil {
		cont.response = utils.BuildFailedResponse(err.Error())
		ctx.JSON(http.StatusInternalServerError, cont.response)
		return
	}

	cont.response = utils.BuildSuccessResponse(utils.DATA_INSERTED, utils.SALES_LEAD, result)
	ctx.JSON(http.StatusCreated, cont.response)
}

func (cont *sale_controller) FetchAllLeads(ctx *gin.Context) {
	var req dto.FetchAllLeadsRequestDTO

	if err := ctx.ShouldBindUri(&req); err != nil {
		cont.response = utils.RequestParamsMissingResponse(helper.Handle_required_param_error(err))
		ctx.JSON(http.StatusBadRequest, cont.response)
		return
	}

	result, err := cont.sale_serv.FetchAllLeads(req)

	if err != nil {

		if err == sql.ErrNoRows {
			cont.response = utils.BuildFailedResponse("data not found")
			ctx.JSON(http.StatusNotFound, cont.response)
			return
		} else if err == context.DeadlineExceeded {
			cont.response = utils.BuildFailedResponse(helper.Err_Something_Wents_Wrong.Error())
			ctx.JSON(http.StatusInternalServerError, cont.response)
			return
		}
		cont.response = utils.BuildFailedResponse(err.Error())
		ctx.JSON(http.StatusInternalServerError, cont.response)
		return
	}

	cont.response = utils.BuildResponseWithPagination(utils.FETCHED_SUCCESS, "", utils.SALES_LEAD, result)
	ctx.JSON(http.StatusOK, cont.response)
}

func (cont *sale_controller) FetchLeadByLeadId(ctx *gin.Context) {
	m_id, _ := ctx.Get("lead_id")

	lead, err := cont.sale_serv.FetchLeadByLeadId(m_id.(uuid.UUID))

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			cont.response = utils.BuildFailedResponse("lead not found")
			ctx.JSON(http.StatusNotFound, cont.response)
			return
		} else if errors.Is(err, helper.ERR_INVALID_ID) {
			cont.response = utils.BuildFailedResponse(helper.ERR_INVALID_ID.Error())
			ctx.JSON(http.StatusBadRequest, cont.response)
			return
		}
		cont.response = utils.BuildFailedResponse(err.Error())
		ctx.JSON(http.StatusInternalServerError, cont.response)
		return
	}

	cont.response = utils.BuildSuccessResponse(utils.FETCHED_SUCCESS, utils.SALES_LEAD, lead)
	ctx.JSON(http.StatusOK, cont.response)
}

func (cont *sale_controller) IncreaeQuatationCount(ctx *gin.Context) {
	lead_id, _ := ctx.Get("lead_id")
	err := cont.sale_serv.IncreaeQuatationCount(lead_id.(uuid.UUID))

	if err != nil {
		if errors.Is(err, helper.ERR_INVALID_ID) {
			cont.response = utils.BuildFailedResponse(helper.ERR_INVALID_ID.Error())
			ctx.JSON(http.StatusForbidden, cont.response)
			return
		}
		cont.response = utils.BuildFailedResponse(err.Error())
		ctx.JSON(http.StatusInternalServerError, cont.response)
		return
	}
	cont.response = utils.BuildSuccessResponse("quatation count has been increased", utils.SALES_LEAD, utils.EmptyObj{})
	ctx.JSON(http.StatusOK, cont.response)
}

func (cont *sale_controller) FetchLeadCounts(ctx *gin.Context) {
	result, err := cont.sale_serv.FetchLeadCounts()

	if err != nil {
		cont.response = utils.BuildFailedResponse(err.Error())
		ctx.JSON(http.StatusInternalServerError, cont.response)
		return
	}

	cont.response = utils.BuildSuccessResponse(utils.FETCHED_SUCCESS, utils.SALES_LEAD, result)
	ctx.JSON(http.StatusOK, cont.response)
}

func (cont *sale_controller) FetchLeadsByStatus(ctx *gin.Context) {
	var req dto.FetchLeadsByStatusRequestDTO

	if err := ctx.ShouldBindUri(&req); err != nil {
		cont.response = utils.RequestParamsMissingResponse(helper.Handle_required_param_error(err))
		ctx.JSON(http.StatusBadRequest, cont.response)
		return
	}

	result, err := cont.sale_serv.FetchLeadsByStatus(req)

	if err != nil {

		if errors.Is(err, sql.ErrNoRows) {
			cont.response = utils.BuildFailedResponse("data not found")
			ctx.JSON(http.StatusNotFound, cont.response)
			return
		} else if err == context.DeadlineExceeded {
			cont.response = utils.BuildFailedResponse(helper.Err_Something_Wents_Wrong.Error())
			ctx.JSON(http.StatusInternalServerError, cont.response)
			return
		}
		cont.response = utils.BuildFailedResponse(err.Error())
		ctx.JSON(http.StatusInternalServerError, cont.response)
		return
	}

	cont.response = utils.BuildResponseWithPagination(utils.FETCHED_SUCCESS, "", utils.SALES_LEAD, result)
	ctx.JSON(http.StatusOK, cont.response)
}

func (cont *sale_controller) CancelLead(ctx *gin.Context) {
	var req dto.CancelLeadRequestDTO

	if err := ctx.ShouldBindJSON(&req); err != nil {
		cont.response = utils.BuildFailedResponse(helper.Handle_required_param_error(err))
		ctx.JSON(http.StatusBadRequest, cont.response)
		return
	}

	err := cont.sale_serv.CancelLead(req)

	if err != nil {
		cont.response = utils.BuildFailedResponse(err.Error())
		if err == helper.ERR_INVALID_ID {
			ctx.JSON(http.StatusBadRequest, cont.response)
			return
		} else if strings.Contains(err.Error(), "already canceld") {
			ctx.JSON(http.StatusConflict, cont.response)
			return
		}
		ctx.JSON(http.StatusInternalServerError, cont.response)
		return
	}

	cont.response = utils.BuildSuccessResponse("lead has been cancel successfully", utils.SALES_LEAD, nil)
	ctx.JSON(http.StatusOK, cont.response)
}

func (cont *sale_controller) FetchCancelLead(ctx *gin.Context) {
	lead_id := ctx.Param("lead_id")

	if lead_id == "" {
		cont.response = utils.RequestParamsMissingResponse(helper.ERR_REQUIRED_PARAMS)
		ctx.JSON(http.StatusBadRequest, cont.response)
		return
	}

	result, err := cont.sale_serv.FetchCancelLeadByLeadId(lead_id)

	if err != nil {
		if err == sql.ErrNoRows {
			cont.response = utils.BuildFailedResponse(helper.Err_Data_Not_Found.Error())
			ctx.JSON(http.StatusNotFound, cont.response)
			return
		}

		cont.response = utils.BuildFailedResponse(err.Error())
		ctx.JSON(http.StatusInternalServerError, cont.response)
		return
	}

	cont.response = utils.BuildSuccessResponse(utils.FETCHED_SUCCESS, utils.SALES_LEAD, result)
	ctx.JSON(http.StatusOK, cont.response)
}
