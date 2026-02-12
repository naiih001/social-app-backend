package routes

import (
	"social-app/backend/internal/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterAuth(r *gin.Engine) {
	auth := r.Group("/api/v1/auth/")
	{
		auth.POST("register", controllers.Register)
		auth.POST("login", controllers.Login)
		auth.POST("refresh", controllers.Refresh)
		auth.POST("logout", controllers.Logout)
		auth.POST("verify-email", controllers.VerifyEmail)
		auth.POST("forgot-password", controllers.ForgotPassword)
		auth.POST("reset-password", controllers.ResetPassword)
	}
}
