package main

import (
	"database/sql"
	"net/http"
	"os"

	"github.com/aniket-skroman/skroman_sales_service.git/apis"
	"github.com/aniket-skroman/skroman_sales_service.git/apis/database"
	"github.com/aniket-skroman/skroman_sales_service.git/apis/routers"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var (
	PORT = ""
)

func init() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}
	PORT = os.Getenv("PORT")
}

func init_routers() *gin.Engine {
	router := gin.Default()
	return router
}

func routing(route *gin.Engine, store *apis.Store) {
	routers.SalesRouter(route, store)
}

func main() {
	db := database.DB_Instance
	defer func(db *sql.DB) {
		if err := database.CloseDB(db); err != nil {
			panic(err)
		}
	}(db)

	store := apis.NewStore(db)

	router := init_routers()

	router.GET("/", func(ctx *gin.Context) {
		ctx.Redirect(http.StatusSeeOther, "http://3.109.133.20:3000/")
	})

	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	http.Redirect(w, r, "http://3.109.133.20:3000/login", http.StatusSeeOther)
	// })

	routing(router, store)
	router.Run(":" + PORT)
}
