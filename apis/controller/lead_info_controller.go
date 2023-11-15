package controller

import (
	"errors"
	"net/http"

	"github.com/aniket-skroman/skroman_sales_service.git/apis/dto"
	"github.com/aniket-skroman/skroman_sales_service.git/apis/helper"
	"github.com/aniket-skroman/skroman_sales_service.git/apis/services"
	"github.com/aniket-skroman/skroman_sales_service.git/utils"
	"github.com/gin-gonic/gin"
)

type LeadInfoController interface {
	CreateLeadInfo(ctx *gin.Context)
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
		}
		cont.response = utils.BuildFailedResponse(err.Error())
		ctx.JSON(http.StatusInternalServerError, cont.response)
		return
	}

	cont.response = utils.BuildSuccessResponse(utils.DATA_INSERTED, utils.SALES_LEAD, lead_info)
	ctx.JSON(http.StatusOK, cont.response)
}
