package routers

import (
	"github.com/aniket-skroman/skroman_sales_service.git/apis"
	"github.com/aniket-skroman/skroman_sales_service.git/apis/controller"
	"github.com/aniket-skroman/skroman_sales_service.git/apis/middleware"
	"github.com/aniket-skroman/skroman_sales_service.git/apis/repositories"
	"github.com/aniket-skroman/skroman_sales_service.git/apis/services"
	"github.com/gin-gonic/gin"
)

var (
	lead_info_repo    repositories.LeadInfoRepository
	lead_info_service services.LeadInfoService
	lead_info_cont    controller.LeadInfoController
)

func LeadInfoRouter(router *gin.Engine, api *apis.Store) {
	lead_info_repo = repositories.NewLeadInfoRepository(api)
	lead_info_service = services.NewLeadInfoService(lead_info_repo)
	lead_info_cont = controller.NewLeadInfoController(lead_info_service)

	lead := router.Group("/api", middleware.AuthorizeJWT(jwt_servive))
	{
		lead.POST("/lead_info", lead_info_cont.CreateLeadInfo)
	}
}
