package routes

import (
	"social-app/backend/internal/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterTimelines(r *gin.Engine) {
	timeline := r.Group("/api/v1/timeline")
	{
		timeline.GET("/home", controllers.GetFeed)
		timeline.GET("/user/:id", controllers.GetUserFeed)
		timeline.GET("/trending", controllers.GetTrendingFeed)
		timeline.GET("/explore", controllers.GetExploreFeed)
	}
}
