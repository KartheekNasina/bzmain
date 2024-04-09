package routes

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/vivekbnwork/bz-backend/bz-main/controllers"
)

func initializeUserNotificationRoutes(router *gin.Engine, userNotificationController *controllers.UserNotificationController) {
	var baseAPIURL = os.Getenv("BASE_API_URL")

	group := router.Group(baseAPIURL + "/UserNotification")
	{
		group.POST("", userNotificationController.CreateUserNotification)
		group.GET("/:id", userNotificationController.GetUserNotificationByID)
		group.PUT("", userNotificationController.UpdateUserNotification)
		group.DELETE("/:id", userNotificationController.DeleteUserNotification)
	}
}
