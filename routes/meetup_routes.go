package routes

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/vivekbnwork/bz-backend/bz-main/controllers"
)

func initializeMeetupRoutes(router *gin.Engine, MeetupController *controllers.MeetupController) {
	var baseAPIURL = os.Getenv("BASE_API_URL")

	group := router.Group(baseAPIURL + "/Meetup")
	{
		log.Println(group)

		// Define your routes and handlers for Meetup here
	}
}
