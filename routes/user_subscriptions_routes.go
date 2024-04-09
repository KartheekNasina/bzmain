package routes

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/vivekbnwork/bz-backend/bz-main/controllers"
)

func initializeUserSubscriptionRoutes(router *gin.Engine, userSubscriptionController *controllers.UserSubscriptionController) {
	var baseAPIURL = os.Getenv("BASE_API_URL")

	group := router.Group(baseAPIURL + "/user-subscription")
	{
		group.POST("", userSubscriptionController.CreateUserSubscription)
		group.GET("/:userID", userSubscriptionController.GetUserSubscriptionByUserID)
		group.PUT("/:userID", userSubscriptionController.UpdateUserSubscription)
		group.DELETE("/:userID", userSubscriptionController.DeleteUserSubscription)

	}
}
