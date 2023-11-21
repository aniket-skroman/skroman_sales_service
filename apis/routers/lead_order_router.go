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
	lead_order_repo repositories.LeadOrderRepository
	lead_order_ser  services.LeadOrderService
	lead_order_cont controller.LeadOrderController
)

func LeadOrderRouter(route *gin.Engine, db *apis.Store) {
	lead_order_repo = repositories.NewLeadOrderRepository(db)
	lead_order_ser = services.NewLeadOrderService(lead_order_repo)
	lead_order_cont = controller.NewLeadOrderService(lead_order_ser)

	lead_order := route.Group("/api", middleware.AuthorizeJWT(jwt_servive))
	{
		lead_order.POST("/lead_order", lead_order_cont.CreateLeadOrder)
		lead_order.GET("/lead_order/:lead_id", lead_order_cont.FetchOrdersByLeadId)
		lead_order.GET("/lead_order/order/:order_id", lead_order_cont.FetchLeadOrderByOrderId)
		lead_order.DELETE("/lead_order", lead_order_cont.DeleteLeadOrder)
		lead_order.PUT("/lead_order/:order_id", lead_order_cont.UpdateLeadOrder)
	}

	order_quatation := route.Group("/api", middleware.AuthorizeJWT(jwt_servive))
	{
		order_quatation.POST("/order_quatation", lead_order_cont.UploadOrderQuatation)
	}

}
