package routes

import (
	"social-app/backend/internal/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterMedia(r *gin.Engine) {
	media := r.Group("/api/v1/media")
	{
		media.POST("/upload", controllers.UploadMedia)
		media.POST("/confirm", controllers.ConfirmUpload)
		media.GET("/:id", controllers.GetUpload)
	}
}
