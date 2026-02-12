package routes

import (
	"social-app/backend/internal/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterAdmin(r *gin.Engine) {
	admin := r.Group("/api/v1/")
	{
		admin.GET("/admin/users", controllers.GetUsersByAdminPrivilege)
		admin.GET("/admin/posts/:id", controllers.DeletePostByAdminPrivilege)
		admin.GET("/admin/ban/:id", controllers.BanUserByAdminPrivilege)
	}
}
