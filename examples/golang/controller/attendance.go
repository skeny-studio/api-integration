package controller

import (
	config "api_integrasi/database"
	"api_integrasi/model"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)


func CheckIn(c *gin.Context) {

	var req model.AttendanceRecord

	if err := c.ShouldBindJSON(&req); err != nil {

		c.JSON(400, gin.H{
			"success": false,
			"message": err.Error(),
		})

		return
	}

	var existing model.AttendanceRecord

	err := config.Database.
		Where("session_key = ?", req.SessionKey).
		First(&existing).
		Error

	if err == nil {

		c.JSON(200, gin.H{
			"success": true,
			"message": "Already synced",
		})

		return
	}

	if err != gorm.ErrRecordNotFound {

		c.JSON(500, gin.H{
			"success": false,
			"message": err.Error(),
		})

		return
	}

	if err := config.Database.Create(&req).Error; err != nil {

		c.JSON(500, gin.H{
			"success": false,
			"message": err.Error(),
		})

		return
	}

	c.JSON(200, gin.H{
		"success": true,
		"message": "Check-in saved",
	})
}

func CheckOut(c *gin.Context) {

	var req model.AttendanceRecord

	if err := c.ShouldBindJSON(&req); err != nil {

		c.JSON(400, gin.H{
			"success": false,
			"message": err.Error(),
		})

		return
	}

	result := config.Database.
		Model(&model.AttendanceRecord{}).
		Where("session_key = ?", req.SessionKey).
		Updates(map[string]interface{}{
			"check_out_time": req.CheckOutTime,
		})

	if result.Error != nil {

		c.JSON(500, gin.H{
			"success": false,
			"message": result.Error.Error(),
		})

		return
	}

	if result.RowsAffected == 0 {

		c.JSON(404, gin.H{
			"success": false,
			"message": "Session not found",
		})

		return
	}

	c.JSON(200, gin.H{
		"success": true,
		"message": "Check-out updated",
	})
}