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

type LeadOrderController interface {
	CreateLeadOrder(*gin.Context)
	FetchOrdersByLeadId(*gin.Context)
	DeleteLeadOrder(*gin.Context)
	UpdateLeadOrder(*gin.Context)
	FetchLeadOrderByOrderId(*gin.Context)
}

type lead_order_cont struct {
	serv     services.LeadOrderService
	response map[string]interface{}
}

func NewLeadOrderService(serv services.LeadOrderService) LeadOrderController {
	return &lead_order_cont{
		serv:     serv,
		response: make(map[string]interface{}),
	}
}

func (cont *lead_order_cont) CreateLeadOrder(ctx *gin.Context) {
	var req dto.CreateLeadOrderRequestDTO

	if err := ctx.ShouldBindJSON(&req); err != nil {
		cont.response = utils.BuildFailedResponse(helper.Handle_required_param_error(err))
		ctx.JSON(http.StatusBadRequest, cont.response)
		return
	}

	order, err := cont.serv.CreateLeadOrder(req)

	if err != nil {
		cont.response = utils.BuildFailedResponse(err.Error())
		if errors.Is(err, helper.ERR_INVALID_ID) {
			ctx.JSON(http.StatusBadRequest, cont.response)
			return
		}
		ctx.JSON(http.StatusInternalServerError, cont.response)
		return
	}

	cont.response = utils.BuildSuccessResponse(utils.DATA_INSERTED, utils.SALES_LEAD, order)
	ctx.JSON(http.StatusCreated, cont.response)
}

func (cont *lead_order_cont) FetchOrdersByLeadId(ctx *gin.Context) {
	lead_id := ctx.Param("lead_id")

	orders, err := cont.serv.FetchOrdersByLeadId(lead_id)

	if err != nil {
		cont.response = utils.BuildFailedResponse(err.Error())
		if errors.Is(err, helper.ERR_INVALID_ID) {
			ctx.JSON(http.StatusBadRequest, cont.response)
			return
		} else if errors.Is(err, sql.ErrNoRows) {
			cont.response = utils.BuildFailedResponse(helper.Err_Data_Not_Found.Error())
			ctx.JSON(http.StatusNotFound, cont.response)
			return
		}

		ctx.JSON(http.StatusInternalServerError, cont.response)
		return
	}

	cont.response = utils.BuildSuccessResponse(utils.FETCHED_SUCCESS, utils.SALES_LEAD, orders)
	ctx.JSON(http.StatusOK, cont.response)
}

func (cont *lead_order_cont) DeleteLeadOrder(ctx *gin.Context) {
	var req dto.DeleteLeadOrderRequestDTO

	if err := ctx.ShouldBindQuery(&req); err != nil {
		cont.response = utils.BuildFailedResponse(helper.Handle_required_param_error(err))
		ctx.JSON(http.StatusBadRequest, cont.response)
		return
	}

	err := cont.serv.DeleteLeadOrder(req)

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

func (cont *lead_order_cont) UpdateLeadOrder(ctx *gin.Context) {
	var req dto.UpdateLeadOrderRequestDTO

	if err := ctx.ShouldBindJSON(&req); err != nil {
		cont.response = utils.BuildFailedResponse(helper.Handle_required_param_error(err))
		ctx.JSON(http.StatusBadRequest, cont.response)
		return
	}

	order_id := ctx.Param("order_id")

	result, err := cont.serv.UpdateLeadOrder(req, order_id)

	if err != nil {
		cont.response = utils.BuildFailedResponse(err.Error())

		if errors.Is(err, helper.ERR_INVALID_ID) {
			ctx.JSON(http.StatusBadRequest, cont.response)
			return
		} else if errors.Is(err, sql.ErrNoRows) {
			cont.response = utils.BuildFailedResponse(helper.Err_Update_Failed.Error())
			ctx.JSON(http.StatusForbidden, cont.response)
			return
		}

		ctx.JSON(http.StatusInternalServerError, cont.response)
		return
	}

	cont.response = utils.BuildSuccessResponse(utils.UPDATE_SUCCESS, utils.SALES_LEAD, result)
	ctx.JSON(http.StatusOK, cont.response)
}

func (cont *lead_order_cont) FetchLeadOrderByOrderId(ctx *gin.Context) {
	order_id := ctx.Param("order_id")

	result, err := cont.serv.FetchOrdersByOrderId(order_id)

	if err != nil {
		cont.response = utils.BuildFailedResponse(err.Error())
		if errors.Is(err, helper.ERR_INVALID_ID) {
			ctx.JSON(http.StatusBadRequest, cont.response)
			return
		} else if errors.Is(err, sql.ErrNoRows) {
			cont.response = utils.BuildFailedResponse(helper.Err_Data_Not_Found.Error())
			ctx.JSON(http.StatusNotFound, cont.response)
			return
		}

		ctx.JSON(http.StatusInternalServerError, cont.response)
		return
	}

	cont.response = utils.BuildSuccessResponse(utils.FETCHED_SUCCESS, utils.SALES_LEAD, result)
	ctx.JSON(http.StatusOK, cont.response)
}
