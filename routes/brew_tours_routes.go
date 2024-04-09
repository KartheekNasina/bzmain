package routes

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/vivekbnwork/bz-backend/bz-main/controllers"
)

func initializeBrewToursRoutes(router *gin.Engine, brewToursController *controllers.BrewTourController) {
	var baseAPIURL = os.Getenv("BASE_API_URL")

	group := router.Group(baseAPIURL + "/brew-tours")
	{
		group.GET("", brewToursController.GetTours)
		group.PUT("/registrations/:id", brewToursController.UpdateUserBrewTourRegistration)
		group.DELETE("/registrations/:registrationID", brewToursController.DeleteUserBrewTourRegistration)
		group.POST("/registrations", brewToursController.CreateUserBrewTourRegistration)

	}
}
