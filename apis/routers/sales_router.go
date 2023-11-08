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
	sales_service = services.NewSalesLeadService(sales_repository)
	sales_controller = controller.NewSaleLeadsController(sales_service)
	jwt_servive = services.NewJWTService()

	sales_lead := router.Group("/api", middleware.AuthorizeJWT(jwt_servive))
	{
		sales_lead.POST("/sales-lead", sales_controller.CreateNewLead)
	}

}
