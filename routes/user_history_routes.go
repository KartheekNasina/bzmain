package routes

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/vivekbnwork/bz-backend/bz-main/controllers"
)

func initializeUserHistoryRoutes(router *gin.Engine, UserHistoryController *controllers.UserHistoryController) {
	var baseAPIURL = os.Getenv("BASE_API_URL")

	group := router.Group(baseAPIURL + "/UserHistory")
	{
		log.Println(group)

		// Define your routes and handlers for UserHistory here
	}
}
