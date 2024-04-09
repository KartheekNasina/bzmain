package routes

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/vivekbnwork/bz-backend/bz-main/controllers"
)

func initializeLeaderboardRoutes(router *gin.Engine, LeaderboardController *controllers.LeaderboardController) {
	var baseAPIURL = os.Getenv("BASE_API_URL")

	group := router.Group(baseAPIURL + "/Leaderboard")
	{
		log.Println(group)

		// Define your routes and handlers for Leaderboard here
	}
}
