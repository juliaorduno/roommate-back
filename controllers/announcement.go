package controllers

import (
	"../models"
	"github.com/gin-gonic/gin"
)

var announcementModel = new(models.AnnouncementModel)

type AnnouncementController struct{}

// ListAnnouncements godoc
// @Summary List announcements
// @Description get announcements
// @Tags announcements
// @Accept  json
// @Produce  json
// @Success 200 {array} models.Announcement
// @Router /announcements [get]
func (announcement *AnnouncementController) Find(c *gin.Context) {
	var list []models.Announcement
	err := announcementModel.Find(&list)
	if err != nil {
		c.JSON(404, gin.H{"message": "Find error", "error": err.Error()})
		c.Abort()
	} else {
		c.JSON(200, gin.H{"data": list})
	}
}

// GetAnnouncement godoc
// @Summary Get an announcement
// @Description get announcement by ID
// @Tags announcements
// @Accept  json
// @Produce  json
// @Param id path int true "Announcement ID"
// @Success 200 {object} models.Announcement
// @Router /announcements/{id} [get]
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

// CreateAnnouncement godoc
// @Summary Create an announcement
// @Description create by json announcement
// @Tags announcements
// @Accept  json
// @Produce  json
// @Param announcement body models.Announcement true "Add announcement"
// @Success 200 {object} models.Announcement
// @Router /announcements [post]
func (announcement *AnnouncementController) Create(c *gin.Context) {
	var data models.Announcement
	if c.BindJSON(&data) != nil {
		c.JSON(400, gin.H{"message": "Invalid form", "form": data})
		c.Abort()
		return
	}

	err := announcementModel.Create(&data)
	if err != nil {
		c.JSON(500, gin.H{"message": "Announcement could not be created", "error": err.Error()})
		c.Abort()
		return
	}

	c.JSON(200, gin.H{"message": "Announcement Created", "new_announcement": data})

}

// ListGroupAnnouncements godoc
// @Summary List group announcements
// @Description get group announcements
// @Tags groups
// @Accept  json
// @Produce  json
// @Param id path int true "Group ID"
// @Success 200 {array} models.Announcement
// @Router /group/{id}/announcements [get]
func(announcement *AnnouncementController) GetAnnouncements(c *gin.Context) {
	groupID := c.Param("id")
	var list []models.Announcement
	err := announcementModel.GetAnnouncements(groupID, &list)

	if err != nil{
		c.JSON(500, gin.H{"message": "Announcements could not be retrieved", "error": err.Error()})
		c.Abort()
		return
	}

	c.JSON(200, gin.H{"message": "Announcements retrieved", "announcements": list})
}