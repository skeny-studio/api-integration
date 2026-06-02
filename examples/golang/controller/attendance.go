package controller

import (
	config "api_integrasi/database"
	"api_integrasi/model"

	"github.com/gin-gonic/gin"
)

func ReceiveAttendance(c *gin.Context) {

	var req model.AttendanceRecord

	if err := c.ShouldBindJSON(&req); err != nil {

		c.JSON(400, gin.H{
			"success": false,
			"message": err.Error(),
		})

		return
	}

	if err := config.Database.Create(&req).Error; err != nil {

		c.JSON(500, gin.H{
			"success": false,
			"message": "Failed save attendance",
		})

		return
	}

	c.JSON(200, gin.H{
		"success": true,
		"message": "OK",
	})
}