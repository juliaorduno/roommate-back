package controllers

import (
	"../models"
	"github.com/gin-gonic/gin"
)

var meetingModel = new(models.MeetingModel)

type MeetingController struct{}

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

func (meeting *MeetingController) Create(c *gin.Context) {
	var data models.Meeting
	if c.BindJSON(&data) != nil {
		c.JSON(406, gin.H{"message": "Invalid form", "form": data})
		c.Abort()
		return
	}

	err := meetingModel.Create(&data)
	if err != nil {
		c.JSON(406, gin.H{"message": "Meeting could not be created", "error": err.Error()})
		c.Abort()
		return
	}

	c.JSON(200, gin.H{"message": "Meeting Created", "new_meeting": data})
}

func(meeting *MeetingController) GetMeetings(c *gin.Context){
		var data struct {
			GroupID 		int			`json:"group_id"`
		}

		if c.BindJSON(&data) != nil {
			c.JSON(406, gin.H{"message": "Invalid form", "form": data})
			c.Abort()
			return
		}

		groupID := data.GroupID
		var list []models.Meeting
		err := meetingModel.GetMeetings(groupID, &list)

		if err != nil {
			c.JSON(406, gin.H{"message": "Meetings could not be retrieved", "error": err.Error()})
			c.Abort()
			return
		}

		c.JSON(200, gin.H{"message": "Meetings retrieved", "meetings": list})

	}