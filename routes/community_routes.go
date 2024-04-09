package routes

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/vivekbnwork/bz-backend/bz-main/controllers"
)

func initializeCommunityRoutes(router *gin.Engine, CommunityController *controllers.CommunityController) {
	var baseAPIURL = os.Getenv("BASE_API_URL")

	group := router.Group(baseAPIURL + "/Community")
	{
		log.Println(group)

		// Define your routes and handlers for Community here
	}
}
