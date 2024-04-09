package routes

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/vivekbnwork/bz-backend/bz-main/controllers"
)

func initializeUserReferralRoutes(router *gin.Engine, userReferralController *controllers.UserReferralController) {
	var baseAPIURL = os.Getenv("BASE_API_URL")

	group := router.Group(baseAPIURL + "/UserReferral")
	{
		group.GET("", userReferralController.GetUserReferral)
		group.GET("/:id", userReferralController.GetUserReferralByID)
		group.POST("", userReferralController.CreateUserReferral)
		group.PUT("/:id", userReferralController.UpdateUserReferral)
		group.DELETE("/:id", userReferralController.DeleteUserReferral)
	}
}
