package routes

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/vivekbnwork/bz-backend/bz-main/controllers"
)

func initializeUserBreweryCheckinsRoutes(router *gin.Engine, userBreweryCheckinsController *controllers.UserBreweryCheckinsController) {
	var baseAPIURL = os.Getenv("BASE_API_URL")

	group := router.Group(baseAPIURL + "/user-brewery-checkins")
	{
		group.POST("", userBreweryCheckinsController.CreateUserBreweryCheckin)
		group.PUT("/:id", userBreweryCheckinsController.UpdateUserBreweryCheckin)
		group.DELETE("/:id", userBreweryCheckinsController.DeleteUserBreweryCheckin)
		group.GET("", userBreweryCheckinsController.ListUserBreweryCheckins)

	}
}
