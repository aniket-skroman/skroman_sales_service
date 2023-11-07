package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/aniket-skroman/skroman_sales_service.git/apis/database"
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
	fmt.Println("ENV has been loaded..")
}

func init_routers() *gin.Engine {
	router := gin.Default()
	return router
}

func routing(route *gin.Engine) {

}

func main() {
	db := database.DB_Instance
	defer func(db *sql.DB) {
		if err := database.CloseDB(db); err != nil {
			panic(err)
		}
	}(db)
	router := init_routers()
	routing(router)
	router.Run(":" + PORT)
}
