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
	auth.GET("/me", middleware.Auth, controller.Me)

	// Laporan Api
	laporan := new(controller.LaporanController)
	r.GET("/laporan", middleware.Auth, laporan.Index)
	r.POST("/laporan", middleware.Auth, laporan.Store)
	r.GET("/laporan/:id", middleware.Auth, laporan.Show)
	r.PATCH("/laporan/:id", middleware.Auth, laporan.Update)
	r.DELETE("/laporan/:id", middleware.Auth, laporan.Delete)
}
