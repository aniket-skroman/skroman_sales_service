package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/aniket-skroman/skroman_sales_service.git/apis"
	"github.com/aniket-skroman/skroman_sales_service.git/apis/database"
	"github.com/aniket-skroman/skroman_sales_service.git/apis/middleware"
	"github.com/aniket-skroman/skroman_sales_service.git/apis/routers"
	"github.com/aniket-skroman/skroman_sales_service.git/apis/services"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	_ "github.com/lib/pq"
)

var (
	PORT = "9001"
)

func CORSConfig() cors.Config {
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"http://localhost:3000", "http://3.109.133.20:3000", "http://13.234.110.115:3000"}
	corsConfig.AllowCredentials = true
	corsConfig.AddAllowHeaders("Access-Control-Allow-Headers", "access-control-allow-origin, access-control-allow-headers", "Content-Type", "X-XSRF-TOKEN", "Accept", "Origin", "X-Requested-With", "Authorization")
	corsConfig.AddAllowMethods("GET", "POST", "PUT", "DELETE")
	return corsConfig
}

const (
	ContentTypeBinary = "application/octet-stream"
	ContentTypeForm   = "application/x-www-form-urlencoded"
	ContentTypeJSON   = "application/json"
	ContentTypeHTML   = "text/html; charset=utf-8"
	ContentTypeText   = "text/plain; charset=utf-8"
)

func init_routers() *gin.Engine {
	router := gin.New()
	jwt_serv := services.NewJWTService()
	router.Use(cors.New(CORSConfig()))
	router.Use(middleware.AuthorizeJWT(jwt_serv))
	router.Use(middleware.ValidateRequest)
	return router
}

func routing(route *gin.Engine, store *apis.Store) {
	routers.SalesRouter(route, store)
	routers.LeadInfoRouter(route, store)
	routers.LeadOrderRouter(route, store)
	routers.LeadVisitRouter(route, store)
}

func main() {
	db := database.DB_Instance
	defer func(db *sql.DB) {
		if err := database.CloseDB(db); err != nil {
			log.Fatal("connection closed issued : ", err)
		}
	}(db)

	store := apis.NewStore(db)
	router := init_routers()

	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"Data": "Service working..."})
	})

	routing(router, store)
	router.Run(":" + PORT)
}
