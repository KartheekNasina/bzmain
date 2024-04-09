package routes

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/vivekbnwork/bz-backend/bz-main/controllers"
)

func initializeUserFavoriteBreweryRoutes(router *gin.Engine, userFavoriteBreweryController *controllers.UserFavoriteBreweryController) {
	var baseAPIURL = os.Getenv("BASE_API_URL")

	group := router.Group(baseAPIURL + "/UserFavoriteBrewery")
	{
		group.POST("", userFavoriteBreweryController.CreateUserFavoriteBrewery)
		group.GET("/:id", userFavoriteBreweryController.GetUserFavoriteBrewery)
		group.DELETE("/:id", userFavoriteBreweryController.DeleteUserFavoriteBrewery)
	}
}
