package routes

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/vivekbnwork/bz-backend/bz-main/controllers"
)

func initializeBreweryReviewRoutes(router *gin.Engine, BreweryReviewController *controllers.BreweryReviewController) {
	var baseAPIURL = os.Getenv("BASE_API_URL")

	group := router.Group(baseAPIURL + "/BreweryReview")
	{
		log.Println(group)

		// Define your routes and handlers for BreweryReview here
	}
}
