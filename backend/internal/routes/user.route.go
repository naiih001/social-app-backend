package routes

import (
	"social-app/backend/internal/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterUser(r *gin.Engine) {
	users := r.Group("api/v1/users")
	{
		users.GET("/:id", controllers.GetUserProfile)
		users.GET("/username/:username", controllers.GetUserResolveByUsername)
		users.PATCH("/me", controllers.UpdateUserProfile)
		users.DELETE("/me", controllers.DeactivateUserProfile)
		users.GET("/:id/followers", controllers.GetUserFollowers)
		users.GET("/:id/following", controllers.GetUserFollowing)
		users.GET(":id/posts", controllers.GetUserPosts)
		users.GET("/search", controllers.Search)

		// Likes
		users.GET("/:id/likes", controllers.GetLikes)
	}
}
