// repository/brewery_routes.go
package routes

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/vivekbnwork/bz-backend/bz-main/controllers"
)

func initializeBreweryRoutes(router *gin.Engine, breweryController *controllers.BreweryController) {
	var baseAPIURL = os.Getenv("BASE_API_URL")

	group := router.Group(baseAPIURL + "/breweries")
	{
		group.GET("/:id", breweryController.GetBrewery)
		group.GET("", breweryController.GetBreweries)
		// Add more brewery-related routes here
	}
}
