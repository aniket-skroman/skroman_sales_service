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
	visit_repo repositories.LeadVisitRepository
	visit_serv services.LeadVisitService
	visit_cont controller.LeadVisitController
)

func LeadVisitRouter(router *gin.Engine, db *apis.Store) {
	visit_repo = repositories.NewLeadVisitRepository(db)
	visit_serv = services.NewLeadVisitService(visit_repo, sales_service)
	visit_cont = controller.NewLeadVisitService(visit_serv)

	lead_visit := router.Group("/api", middleware.AuthorizeJWT(jwt_servive))
	{
		lead_visit.POST("/lead-visit", visit_cont.CreateLeadVisit)
		lead_visit.GET("/lead-visit/:lead_id", visit_cont.FetchAllVisitByLead)
	}
}
