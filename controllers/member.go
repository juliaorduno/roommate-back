package controllers

import (
	"../models"
	"github.com/gin-gonic/gin"
)

var memberModel = new(models.MemberModel)

type MemberController struct{}

// ListMembers godoc
// @Summary List members
// @Description get members
// @Tags members
// @Accept  json
// @Produce  json
// @Success 200 {array} models.Member
// @Router /members [get]
func (member *MemberController) Find(c *gin.Context) {
	var list []models.Member
	err := memberModel.Find(&list)
	if err != nil {
		c.JSON(404, gin.H{"message": "Find Error", "error": err.Error()})
		c.Abort()
	} else {
		c.JSON(200, gin.H{"data": list})
	}
}

// GetMember godoc
// @Summary Get a member
// @Description get member by ID
// @Tags members
// @Accept  json
// @Produce  json
// @Param id path int true "Member ID"
// @Success 200 {object} models.Member
// @Router /members/{id} [get]
func (member *MemberController) Get(c *gin.Context) {
	id := c.Param("id")
	var data models.Member
	err := memberModel.Get(id, &data)
	if err != nil {
		c.JSON(404, gin.H{"message": "Member not found", "error": err.Error()})
		c.Abort()
	} else {
		c.JSON(200, gin.H{"data": data})
	}
}

func (member *MemberController) GetByEmail(c *gin.Context) {
	email := c.Param("email")
	var data models.Member
	err := memberModel.GetCurrentUser(email, &data)
	if err != nil {
		c.JSON(404, gin.H{"message": "Member not found", "error": err.Error()})
		c.Abort()
	} else {
		c.JSON(200, gin.H{"data": data})
	}
}

// CreateMember godoc
// @Summary Create a member
// @Description create by json member
// @Tags members
// @Accept  json
// @Produce  json
// @Param member body models.Member true "Add member"
// @Success 200 {object} models.Member
// @Router /members [post]
func (member *MemberController) Create(c *gin.Context) {
	var data models.Member
	if c.BindJSON(&data) != nil {
		c.JSON(400, gin.H{"message": "Invalid form", "form": data})
		c.Abort()
		return
	}

	err := memberModel.Create(&data)
	if err != nil {
		c.JSON(500, gin.H{"message": "Member could not be created", "error": err.Error()})
		c.Abort()
		return
	}

	c.JSON(200, gin.H{"message": "Member Created", "new_member": data})
}

// JoinGroup godoc
// @Summary Join a group
// @Description join group by groupcode and memberID
// @Tags members
// @Accept  json
// @Produce  json
// @Success 200 {object} models.Member
// @Router /members [post]
func (member *MemberController) JoinGroup(c *gin.Context) {
	var data struct {
		GroupCode string `json:"group_code"`
		MemberID  int    `json:"member_id"`
	}
	var updatedMember models.Member

	if c.BindJSON(&data) != nil {
		c.JSON(400, gin.H{"message": "Invalid form", "form": data})
		c.Abort()
		return
	}

	if data.MemberID <= 0 {
		c.JSON(400, gin.H{"message": "Invalid form", "error": "You must add a memberID"})
		c.Abort()
		return
	}

	if data.GroupCode == "" {
		c.JSON(400, gin.H{"message": "Invalid form", "error": "You must add a group code"})
		c.Abort()
		return
	}

	err := memberModel.JoinGroup(data.MemberID, data.GroupCode, &updatedMember)
	if err != nil {
		c.JSON(500, gin.H{"message": "Member could not join group", "error": err.Error()})
		c.Abort()
		return
	}

	c.JSON(200, gin.H{"message": "Group Joined", "data": updatedMember})
}