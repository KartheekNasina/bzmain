package routes

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/vivekbnwork/bz-backend/bz-main/controllers"
)

func initializePaymentRoutes(router *gin.Engine, PaymentController *controllers.PaymentController) {
	var baseAPIURL = os.Getenv("BASE_API_URL")

	group := router.Group(baseAPIURL + "/Payment")
	{
		log.Println(group)

		// Define your routes and handlers for Payment here
	}
}
