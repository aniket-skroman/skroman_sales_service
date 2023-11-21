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
	sales_repository repositories.SalesRepository
	sales_service    services.SalesLeadService
	jwt_servive      services.JWTService
	sales_controller controller.SaleLeadsController
)

func SalesRouter(router *gin.Engine, store *apis.Store) {
	sales_repository = repositories.NewSalesRepository(store)
	lead_order_repo = repositories.NewLeadOrderRepository(store)
	lead_order_ser = services.NewLeadOrderService(lead_order_repo)
	sales_service = services.NewSalesLeadService(sales_repository, lead_order_ser)
	sales_controller = controller.NewSaleLeadsController(sales_service)
	jwt_servive = services.NewJWTService()

	sales_lead := router.Group("/api", middleware.AuthorizeJWT(jwt_servive))
	{
		sales_lead.POST("/sales-lead", sales_controller.CreateNewLead)
		sales_lead.GET("/sales-lead/:lead_id", sales_controller.FetchLeadByLeadId)
		sales_lead.GET("/sales-leads/:page_id/:page_size/:status", sales_controller.FetchLeadsByStatus)
		sales_lead.GET("/sales-leads/:page_id/:page_size", sales_controller.FetchAllLeads)
		sales_lead.GET("/sale_lead_counts", sales_controller.FetchLeadCounts)
	}

}
