package controllers

import (
	"../models"
	"github.com/gin-gonic/gin"
)

var groupModel = new(models.GroupModel)

type GroupController struct{}

// ListGroups godoc
// @Summary List groups
// @Description get groups
// @Tags groups
// @Accept  json
// @Produce  json
// @Success 200 {array} models.RGroup
// @Router /groups [get]
func (RGroup *GroupController) Find(c *gin.Context) {
	var list []models.RGroup
	err := groupModel.Find(&list)
	if err != nil {
		c.JSON(404, gin.H{"message": "Find Error", "error": err.Error()})
		c.Abort()
	} else {
		c.JSON(200, gin.H{"data": list})
	}
}

// GetGroup godoc
// @Summary Get a group
// @Description get group by ID
// @Tags groups
// @Accept  json
// @Produce  json
// @Param id path int true "Group ID"
// @Success 200 {object} models.RGroup
// @Router /groups/{id} [get]
func (RGroup *GroupController) Get(c *gin.Context) {
	id := c.Param("id")
	var data models.RGroup
	err := groupModel.Get(id, &data)
	if err != nil {
		c.JSON(404, gin.H{"message": "Group not found", "error": err.Error()})
		c.Abort()
	} else {
		c.JSON(200, gin.H{"data": data})
	}
}

func (RGroup *GroupController) Create(c *gin.Context) {
	var data models.RGroup
	if c.BindJSON(&data) != nil {
		c.JSON(400, gin.H{"message": "Invalid form", "form": data})
		c.Abort()
		return
	}

	err := groupModel.Create(&data)
	if err != nil {
		c.JSON(500, gin.H{"message": "Group could not be created", "error": err.Error()})
		c.Abort()
		return
	}

	c.JSON(200, gin.H{"message": "Group Created", "new_rgroup": data})
}

// CreateGroup godoc
// @Summary Create a group
// @Description create by json group
// @Tags groups
// @Accept  json
// @Produce  json
// @Param group body models.RGroup true "Add group"
// @Success 200 {object} models.RGroup
// @Router /groups [post]
func (RGroup *GroupController) CreateGroup(c *gin.Context) {
	var data models.RGroup
	if c.BindJSON(&data) != nil {
		c.JSON(400, gin.H{"message": "Invalid form", "form": data})
		c.Abort()
		return
	}

	if data.Name == "" {
		c.JSON(400, gin.H{"message": "Invalid form", "error": "You must add a group name"})
		c.Abort()
		return
	}

	if data.CreatedBy <= 0 {
		c.JSON(400, gin.H{"message": "Invalid form", "error": "You must add a memberID"})
		c.Abort()
		return
	}

	if data.Size <= 0 {
		data.Size = 1
	}
	data.Admin = data.CreatedBy

	err := groupModel.Create(&data)
	if err != nil {
		c.JSON(500, gin.H{"message": "Group could not be created", "error": err.Error()})
		c.Abort()
		return
	}

	var member models.Member
	err = memberModel.AddToGroup(data.CreatedBy, uint(data.ID), &member)
	if err != nil {
		c.JSON(500, gin.H{"message": "Group could not be assigned to member", "error": err.Error()})
		c.Abort()
		return
	}

	c.JSON(200, gin.H{"message": "Group Created", "new_rgroup": data})
}