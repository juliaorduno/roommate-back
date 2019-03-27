package controllers

import (
	"../models"
	"github.com/gin-gonic/gin"
)

var taskModel = new(models.TaskModel)

type TaskController struct{}

func (task *TaskController) Find(c *gin.Context) {
	var list []models.Task
	err := taskModel.Find(&list)
	if err != nil {
		c.JSON(404, gin.H{"message": "Find Error", "error": err.Error()})
		c.Abort()
	} else {
		c.JSON(200, gin.H{"data": list})
	}
}

func (task *TaskController) Get(c *gin.Context) {
	id := c.Param("id")
	var data models.Task
	err := taskModel.Get(id, &data)
	if err != nil {
		c.JSON(404, gin.H{"message": "Task not found", "error": err.Error()})
		c.Abort()
	} else {
		c.JSON(200, gin.H{"data": data})
	}
}

func (task *TaskController) Create(c *gin.Context) {
	var data models.Task
	if c.BindJSON(&data) != nil {
		c.JSON(400, gin.H{"message": "Invalid form", "form": data})
		c.Abort()
		return
	}

	err := taskModel.Create(&data)
	if err != nil {
		c.JSON(500, gin.H{"message": "Task could not be created", "error": err.Error()})
		c.Abort()
		return
	}

	c.JSON(200, gin.H{"message": "Task Created", "new_task": data})
}

func (task *TaskController) GetToDos(c *gin.Context) {
	var data struct {
		GroupID		int		`json:"group_id"`
	}
	if c.BindJSON(&data) != nil {
		c.JSON(400, gin.H{"message": "Invalid form", "form": data})
		c.Abort()
		return
	}

	groupID := data.GroupID
	var list []models.Task
	err := taskModel.GetToDos(groupID, &list)
	
	if err != nil {
		c.JSON(500, gin.H{"message": "To Dos could not be retrieved", "error": err.Error()})
		c.Abort()
		return
	}

	c.JSON(200, gin.H{"message": "To dos retrieved", "todos": list})
}

func (task *TaskController) FinishTask(c *gin.Context) {
	var data struct {
		ID int `json:"id"`
		FinishedBy uint `json:"finished_by"`
	}

	if c.BindJSON(&data) != nil {
		c.JSON(400, gin.H{"message": "Invalid form", "form": data})
		c.Abort()
		return
	}

	if data.ID <= 0 {
		c.JSON(400, gin.H{"message": "Invalid form", "error": "You must add an ID"})
		c.Abort()
		return
	}

	if data.FinishedBy <= 0 {
		c.JSON(400, gin.H{"message": "Invalid form", "error": "You must add a memberID"})
		c.Abort()
		return
	}
	
	var updatedTask models.Task

	err := taskModel.FinishTask(data.ID, data.FinishedBy, &updatedTask)
	if err != nil {
		c.JSON(500, gin.H{"message": "Task could not be updated", "error": err.Error()})
		c.Abort()
		return
	}

	c.JSON(200, gin.H{"message": "Task Finished", "new_shopping_item": updatedTask})
} 