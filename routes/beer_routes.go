package routes

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/vivekbnwork/bz-backend/bz-main/controllers"
)

func initializeBeerRoutes(router *gin.Engine, beerController *controllers.BeerController) {
	var baseAPIURL = os.Getenv("BASE_API_URL")

	group := router.Group(baseAPIURL + "/beers")
	{
		group.GET("/:id", beerController.GetBeer)
		group.GET("/type/:type", beerController.GetBeersBasedOnType)
		group.GET("/brewery/:id", beerController.GetBeersBasedOnBreweryID)

	}
}
