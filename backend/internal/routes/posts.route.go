package routes

import (
	"social-app/backend/internal/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterPosts(r *gin.Engine) {
	posts := r.Group("/api/v1/posts")
	{
		posts.POST("/", controllers.CreatePost)
		posts.GET("/:id", controllers.GetSinglePost)
		posts.DELETE("/:id", controllers.DeletePost)
		posts.PATCH("/:id", controllers.UpdatePost)
		posts.GET("/:id/likes", controllers.GetPostLikes)
		posts.GET("/:id/reposts", controllers.GetPostReposts)

		// Comments
		posts.POST("/:id/comments", controllers.CreateComment)
		posts.GET("/:id/comments", controllers.GetComments)
		r.DELETE("api/v1/comments/:id", controllers.DeleteComment)

		// Likes
		posts.POST("/:id/like", controllers.CreateLike)
		posts.DELETE("/:id/like", controllers.DeleteLike)

		// Reposts
		posts.POST("/:id/repost", controllers.CreateRepost)
		posts.DELETE("/:id/repost", controllers.DeleteRepost)
	}
}
