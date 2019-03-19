package controllers

import (
	"../models"
	"github.com/gin-gonic/gin"
)

var announcementModel = new(models.AnnouncementModel)

type AnnouncementController struct{}

func (announcement *AnnouncementController) Find(c *gin.Context) {
	var list []models.Announcement
	err := announcementModel.Find(&list)
	if err != nil {
		c.JSON(404, gin.H{"message": "Find Error", "error": err.Error()})
		c.Abort()
	} else {
		c.JSON(200, gin.H{"data": list})
	}
}

func (announcement *AnnouncementController) Get(c *gin.Context) {
	id := c.Param("id")
	var data models.Announcement
	err := announcementModel.Get(id, &data)
	if err != nil {
		c.JSON(404, gin.H{"message": "Announcement not found", "error": err.Error()})
		c.Abort()
	} else {
		c.JSON(200, gin.H{"data": data})
	}
}

func (announcement *AnnouncementController) Create(c *gin.Context) {
	var data models.Announcement
	if c.BindJSON(&data) != nil {
		c.JSON(406, gin.H{"message": "Invalid form", "form": data})
		c.Abort()
		return
	}

	err := announcementModel.Create(&data)
	if err != nil {
		c.JSON(406, gin.H{"message": "Announcement could not be created", "error": err.Error()})
		c.Abort()
		return
	}

	c.JSON(200, gin.H{"message": "Announcement Created", "new_announcement": data})

	func(announcement *AnnouncementController) GetAnnouncements(c *gin.Context) {
		var data struct {
			GroupID 		int 		`json:"group_id"`

		}
		if c.JSON(406, gin.H{"message": "Invalid form", "form": data})
		c.Abort()
		return
	}

	groupID = data.GroupID
	var list []models.Announcement
	err := announcementModel.Create(groupID, &list)

	if err != nil{
		c.JSON(406, gin.H{"message": "Announcements could not be retrieved", "error": err.Error()})
		c.Abort()
		return
	}

	c.JSON(200, gin.H{"message": "Announcements retrieved", "announcements": data}


}