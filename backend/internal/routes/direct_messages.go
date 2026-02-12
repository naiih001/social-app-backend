package routes

import (
	"social-app/backend/internal/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterDirectMessage(r *gin.Engine) {
	directMessage := r.Group("/api/v1/messages")
	{
		directMessage.POST("/", controllers.CreateDirectMessage)
		directMessage.GET("/conversations", controllers.GetConversations)
		directMessage.GET("/conversations/:id", controllers.GetSingleConversations)
	}

}
