package controller

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"

	"github.com/aniket-skroman/skroman_sales_service.git/apis/dto"
	"github.com/aniket-skroman/skroman_sales_service.git/apis/helper"
	"github.com/aniket-skroman/skroman_sales_service.git/apis/services"
	"github.com/aniket-skroman/skroman_sales_service.git/utils"
	"github.com/gin-gonic/gin"
)

type SaleLeadsController interface {
	CreateNewLead(ctx *gin.Context)
	FetchAllLeads(ctx *gin.Context)
	FetchLeadByLeadId(ctx *gin.Context)
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

		if errors.Is(err, sql.ErrNoRows) {
			cont.response = utils.BuildFailedResponse("data not found")
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

func (cont *sale_controller) FetchLeadByLeadId(ctx *gin.Context) {
	var lead_id = ctx.Param("lead_id")

	lead, err := cont.sale_serv.FetchLeadByLeadId(lead_id)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			cont.response = utils.BuildFailedResponse("lead not found")
			ctx.JSON(http.StatusNotFound, cont.response)
			return
		}
		cont.response = utils.BuildFailedResponse(err.Error())
		ctx.JSON(http.StatusInternalServerError, cont.response)
		return
	}

	cont.response = utils.BuildSuccessResponse(utils.FETCHED_SUCCESS, utils.SALES_LEAD, lead)
	ctx.JSON(http.StatusOK, cont.response)
}
