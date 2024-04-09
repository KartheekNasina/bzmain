package routes

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/vivekbnwork/bz-backend/bz-main/controllers"
)

func initializeEventRoutes(router *gin.Engine, eventController *controllers.EventController) {
	var baseAPIURL = os.Getenv("BASE_API_URL")

	group := router.Group(baseAPIURL + "/event")
	{
		group.GET("", eventController.GetEvents)
		group.GET("/schedule/:id", eventController.GetEventAndSchedule)
		group.POST("/join", eventController.CreateEventJoinRequest)
		// group.PUT("/event-join-requests/:id", eventController.UpdateEventJoinRequest)
		// group.DELETE("/event-join-requests/:requestID", eventController.DeleteEventJoinRequest)

	}
}
