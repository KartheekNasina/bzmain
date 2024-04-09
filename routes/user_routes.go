package routes

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/vivekbnwork/bz-backend/bz-main/controllers"
)

func initializeUserRoutes(router *gin.Engine, userController *controllers.UserController) {
	var baseAPIURL = os.Getenv("BASE_API_URL")

	group := router.Group(baseAPIURL + "/users")
	{
		group.GET("/all", userController.ListUsers)
		group.POST("", userController.CreateUser)
		group.PATCH("/:id", userController.UpdateUser)
		group.DELETE("/:id", userController.DeleteUser)
		group.GET("", userController.GetUsers)

		// Add more user-related routes here
	}
}
