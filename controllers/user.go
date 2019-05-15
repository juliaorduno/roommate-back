package controllers

import (
	"../models"
	"github.com/gin-gonic/gin"
	"github.com/appleboy/gin-jwt"
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

	c.JSON(200, gin.H{"message": "User Created", "user": userData, "data": memberData})
}

func (user *UserController) GetCurrentUser(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	userID := claims["id"].(string)

	var member models.Member
	err := memberModel.GetCurrentUser(userID, &member)
	if err != nil {
		c.JSON(500, gin.H{"message": "Member could not be found", "error": err.Error()})
		c.Abort()
		return
	}

	c.JSON(200, gin.H{"message": "Current User", "member": member})
}