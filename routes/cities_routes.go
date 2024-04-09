// repository/city_routes.go
package routes

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/vivekbnwork/bz-backend/bz-main/controllers"
)

func initializeCityRoutes(router *gin.Engine, cityController *controllers.CityController) {
	var baseAPIURL = os.Getenv("BASE_API_URL")

	group := router.Group(baseAPIURL + "/cities")
	{
		group.GET("", cityController.ListCities)
		// Add more city-related routes here
	}
}
