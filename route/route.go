package route

import (
	"suranect_api/controller"
	"suranect_api/middleware"

	"github.com/gin-gonic/gin"
)

func Init(r *gin.Engine) {
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"Status": "Success 200",
		})
	})

	// Auth Group
	auth := r.Group("/auth")
	auth.POST("/register", controller.Register)
	auth.POST("/login", controller.Login)
	auth.GET("/send_verify_email", middleware.Auth, middleware.VerifiedEmail, controller.SendVerifyEmail)
	auth.POST("/verify_email", middleware.Auth, middleware.VerifiedEmail, controller.VerifyEmail)
}
