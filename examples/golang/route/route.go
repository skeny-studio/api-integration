package route

import (
	"api_integrasi/controller"
	"api_integrasi/middleware"

	"github.com/gin-gonic/gin"
)

func Route(router *gin.Engine) {

	//  Public route
	router.GET("/", func(c *gin.Context) {
		c.String(200, "Welcome To This Api")
	})

	

	attendance := router.Group("/")
	attendance.Use(middleware.APIKeyMiddleware())
	{
		attendance.POST("/check-in", controller.CheckIn)
		attendance.PUT("/check-out", controller.CheckOut)
	}


}
