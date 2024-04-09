package routes

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/vivekbnwork/bz-backend/bz-main/controllers"
)

func initializeAwsRoutes(router *gin.Engine, awsController *controllers.AWSController) {
	var baseAPIURL = os.Getenv("BASE_API_URL")

	userGroup := router.Group(baseAPIURL + "/aws/")
	{
		userGroup.POST("image/upload/get-presigned-url", awsController.GetPresignedURLsForUpload)

		userGroup.POST("image/download/get-presigned-url", awsController.GetPresignedURLsForDownload)

		userGroup.POST("image/delete", awsController.DeleteFilesFromS3)

		// Add more user-related routes here
	}
}
