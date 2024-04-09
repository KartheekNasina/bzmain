package routes

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/vivekbnwork/bz-backend/bz-main/controllers"
)

func initializeBrewClassesRoutes(router *gin.Engine, brewClassController *controllers.BrewClassController) {
	var baseAPIURL = os.Getenv("BASE_API_URL")

	group := router.Group(baseAPIURL + "/BrewClasses")
	{
		group.GET("/:userID/:classID", brewClassController.GetUserBrewClassRegistration)
		group.PUT("/:userID/:classID", brewClassController.UpdateUserBrewClassRegistration)
		group.DELETE("/:registrationID", brewClassController.DeleteUserBrewClassRegistration)
		group.POST("", brewClassController.CreateUserBrewClassRegistration)

	}
}
