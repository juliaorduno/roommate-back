package controllers

import (
	"../models"
	"github.com/gin-gonic/gin"
)

var userModel = new(models.UserModel)

type UserController struct{}

func (user *UserController) Create(c *gin.Context) {
	var data models.User
	if c.BindJSON(&data) != nil {
		c.JSON(400, gin.H{"message": "Invalid form user", "form": data})
		c.Abort()
		return
	}

	err := userModel.Create(&data)
	if err != nil {
		c.JSON(500, gin.H{"message": "User could not be created", "error": err.Error()})
		c.Abort()
		return
	}

	c.JSON(200, gin.H{"message": "User Created", "user": data})
}