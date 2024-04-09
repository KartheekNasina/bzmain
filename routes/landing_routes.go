package routes

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/vivekbnwork/bz-backend/bz-main/controllers"
)

func initializeLandingRoutes(router *gin.Engine, landingController *controllers.LandingController) {
	var baseAPIURL = os.Getenv("BASE_API_URL")

	userGroup := router.Group(baseAPIURL + "/landing")
	{
		userGroup.GET("", landingController.FetchLandingData)
		// Add more user-related routes here
	}
}
