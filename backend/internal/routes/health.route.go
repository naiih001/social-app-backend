package routes

import (
	"social-app/backend/internal/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterHealth(r *gin.Engine) {
	health := r.Group("api/v1/")
	{
		health.GET("/health", controllers.GetHealth)
		health.GET("/metrics", controllers.GetMetrics)
	}
}
