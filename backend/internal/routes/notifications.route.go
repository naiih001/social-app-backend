package routes

import (
	"social-app/backend/internal/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterNotifications(r *gin.Engine) {
	notifications := r.Group("/api/v1/notifications")
	{
		notifications.GET("/", controllers.GetNotfications)
		notifications.PATCH("/:id/read", controllers.UpdateNotifications)
	}
}
