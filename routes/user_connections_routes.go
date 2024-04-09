package routes

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/vivekbnwork/bz-backend/bz-main/controllers"
)

func initializeUserConnectionsRoutes(router *gin.Engine, userConnectionController *controllers.UserConnectionController) {
	var baseAPIURL = os.Getenv("BASE_API_URL")

	group := router.Group(baseAPIURL + "/user-connections")
	{
		group.POST("", userConnectionController.CreateUserConnection)
		group.PUT("/:id", userConnectionController.UpdateUserConnection)
		group.DELETE("/:id", userConnectionController.DeleteUserConnection)

	}
}
