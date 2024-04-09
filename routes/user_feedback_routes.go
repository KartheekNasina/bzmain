package routes

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/vivekbnwork/bz-backend/bz-main/controllers"
)

func initializeUserFeedbackRoutes(router *gin.Engine, userFeedbackController *controllers.UserFeedbackController) {
	var baseAPIURL = os.Getenv("BASE_API_URL")

	group := router.Group(baseAPIURL + "/UserFeedback")
	{
		group.POST("", userFeedbackController.CreateUserFeedback)
		group.GET("/:id", userFeedbackController.GetUserFeedbackByID)
		group.PUT("/:id", userFeedbackController.UpdateUserFeedback)
		group.DELETE("/:id", userFeedbackController.DeleteUserFeedback)
	}
}
