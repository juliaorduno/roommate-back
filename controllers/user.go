package controllers

import (
	"../models"
	"github.com/gin-gonic/gin"
)

var userModel = new(models.UserModel)

type UserController struct{}

func (user *UserController) Create(c *gin.Context) {
	var data models.UserForm
	if c.BindJSON(&data) != nil {
		c.JSON(400, gin.H{"message": "Invalid form user", "form": data})
		c.Abort()
		return
	}

	var memberData models.Member
	memberData.Username = data.Username
	memberData.FullName = data.FullName
	memberData.Email = data.Email

	var userData models.User
	userData.Email = data.Email
	userData.Password = data.Password

	err := userModel.Create(&userData)
	if err != nil {
		c.JSON(500, gin.H{"message": "User could not be created", "error": err.Error()})
		c.Abort()
		return
	}


	memberData.UserID = userData.ID
	err = memberModel.Create(&memberData)
	if err != nil {
		c.JSON(500, gin.H{"message": "Member could not be created", "error": err.Error()})
		c.Abort()
		return
	}

	c.JSON(200, gin.H{"message": "User Created", "user": userData, "member": memberData})
}