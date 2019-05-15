package controllers

import (
	"../models"
	"github.com/gin-gonic/gin"
)

var meetingModel = new(models.MeetingModel)

type MeetingController struct{}

// ListMeetings godoc
// @Summary List meetings
// @Description get meetings
// @Tags meetings
// @Accept  json
// @Produce  json
// @Success 200 {array} models.Meeting
// @Router /meetings [get]
func (meeting *MeetingController) Find(c *gin.Context) {
	var list []models.Meeting
	err := meetingModel.Find(&list)
	if err != nil {
		c.JSON(404, gin.H{"message": "Find Error", "error": err.Error()})
		c.Abort()
	} else {
		c.JSON(200, gin.H{"data": list})
	}
}

// GetMeeting godoc
// @Summary Get a meeting
// @Description get meeting by ID
// @Tags meetings
// @Accept  json
// @Produce  json
// @Param id path int true "Meeting ID"
// @Success 200 {object} models.Meeting
// @Router /meetings/{id} [get]
func (meeting *MeetingController) Get(c *gin.Context) {
	id := c.Param("id")
	var data models.Meeting
	err := meetingModel.Get(id, &data)
	if err != nil {
		c.JSON(404, gin.H{"message": "Meeting not found", "error": err.Error()})
		c.Abort()
	} else {
		c.JSON(200, gin.H{"data": data})
	}
}

// CreateMeeting godoc
// @Summary Create a meeting
// @Description create by json meeting
// @Tags meetings
// @Accept  json
// @Produce  json
// @Param meeting body models.Meeting true "Add Meeting"
// @Success 200 {object} models.Meeting
// @Router /meetings [post]
func (meeting *MeetingController) Create(c *gin.Context) {
	var data models.Meeting
	if c.BindJSON(&data) != nil {
		c.JSON(400, gin.H{"message": "Invalid form", "form": data})
		c.Abort()
		return
	}

	err := meetingModel.Create(&data)
	if err != nil {
		c.JSON(500, gin.H{"message": "Meeting could not be created", "error": err.Error()})
		c.Abort()
		return
	}

	c.JSON(200, gin.H{"message": "Meeting Created", "new_meeting": data})
}

// ListGroupMeetings godoc
// @Summary List group upcoming meetings
// @Description get group upcoming meetings
// @Tags groups
// @Accept  json
// @Produce  json
// @Param id path int true "Group ID"
// @Success 200 {array} models.Meeting
// @Router /group/{id}/meetings [get]
func(meeting *MeetingController) GetMeetings(c *gin.Context){
		groupID := c.Param("id")
	
		err, list := meetingModel.GetMeetings(groupID)

		if err != nil {
			c.JSON(400, gin.H{"message": "Meetings could not be retrieved", "error": err.Error()})
			c.Abort()
			return
		}

		c.JSON(200, gin.H{"message": "Meetings retrieved", "meetings": list})

	}