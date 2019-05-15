package controllers

import (
	"../models"
	"github.com/gin-gonic/gin"
)

var logModel = new(models.LogModel)

type LogController struct{}

func(log *LogController) GetLogs(c *gin.Context) {
	groupID := c.Param("id")

	err, list := logModel.GetLogs(groupID)

	if err != nil{
		c.JSON(500, gin.H{"message": "logs could not be retrieved", "error": err.Error()})
		c.Abort()
		return
	}

	c.JSON(200, gin.H{"message": "logs retrieved", "logs": list})
}