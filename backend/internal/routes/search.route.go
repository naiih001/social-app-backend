package routes

import (
	"social-app/backend/internal/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterSearch(r *gin.Engine) {
	search := r.Group("/api/v1/search")
	{
		search.GET("/users", controllers.SearchUsers)
		search.GET("/users", controllers.SearchPosts)
	}
}
