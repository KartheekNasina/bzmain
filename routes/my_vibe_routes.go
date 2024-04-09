package routes

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/vivekbnwork/bz-backend/bz-main/controllers"
)

func initializeMyVibeRoutes(router *gin.Engine, myVibeController *controllers.MyVibeController) {
	var baseAPIURL = os.Getenv("BASE_API_URL")

	group := router.Group(baseAPIURL + "/MyVibe")
	{
		group.POST("", myVibeController.CreateFoodDrinkRating)
		group.GET("/:id", myVibeController.GetFoodDrinkRatingByID)
		// ... (additional routes for the other functions)

	}
}
