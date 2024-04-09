package routes

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/vivekbnwork/bz-backend/bz-main/controllers"
)

func initializeDrinkPurchaseRequestRoutes(router *gin.Engine, DrinkPurchaseRequestController *controllers.DrinkPurchaseRequestController) {
	var baseAPIURL = os.Getenv("BASE_API_URL")

	group := router.Group(baseAPIURL + "/DrinkPurchaseRequest")
	{
		log.Println(group)

		// Define your routes and handlers for DrinkPurchaseRequest here
	}
}
