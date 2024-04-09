package routes

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/vivekbnwork/bz-backend/bz-main/controllers"
)

func initializeUserBlockedListRoutes(router *gin.Engine, userBlockedListController *controllers.UserBlockedListController) {
	var baseAPIURL = os.Getenv("BASE_API_URL")

	group := router.Group(baseAPIURL + "/user-blocked")
	{
		group.POST("/", userBlockedListController.CreateUserBlocked)
		group.GET("/:userId", userBlockedListController.GetUserBlockedByUserID)
		group.DELETE("/:userId", userBlockedListController.DeleteUserBlocked)

	}
}
