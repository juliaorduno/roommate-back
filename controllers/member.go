package controllers

import (
	"../models"

	"github.com/gin-gonic/gin"
)

var memberModel = new(models.MemberModel)

type MemberController struct{}

func (member *MemberController) Find(c *gin.Context) {
	list, err := memberModel.Find()
	if err != nil {
		c.JSON(404, gin.H{"message": "Find Error", "error": err.Error()})
		c.Abort()
	} else {
		c.JSON(200, gin.H{"data": list})
	}
}

func (member *MemberController) Create(c *gin.Context) {
	var data models.Member
	if c.BindJSON(&data) != nil {
		c.JSON(406, gin.H{"message": "Invalid form", "form": data})
		c.Abort()
		return
	}

	newMember, err := memberModel.Create(data)
	if err != nil {
		c.JSON(406, gin.H{"message": "Member could not be created", "error": err.Error()})
		c.Abort()
		return
	}

	c.JSON(200, gin.H{"message": "Member Created", "new_member": newMember})
}
