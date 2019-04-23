package controllers

import (
	"../models"
	"github.com/gin-gonic/gin"
)

var taskModel = new(models.TaskModel)

type TaskController struct{}

// ListTasks godoc
// @Summary List tasks
// @Description get tasks
// @Tags tasks
// @Accept  json
// @Produce  json
// @Success 200 {array} models.Task
// @Router /tasks [get]
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

// GetTask godoc
// @Summary Get a task
// @Description get task by ID
// @Tags tasks
// @Accept  json
// @Produce  json
// @Param id path int true "Task ID"
// @Success 200 {object} models.Task
// @Router /tasks/{id} [get]
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

// CreateTask godoc
// @Summary Create a task
// @Description create by json task
// @Tags tasks
// @Accept  json
// @Produce  json
// @Param task body models.Task true "Add Task"
// @Success 200 {object} models.Task
// @Router /tasks [post]
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

// ListGroupTasks godoc
// @Summary List group pending tasks
// @Description get group pending tasks
// @Tags groups
// @Accept  json
// @Produce  json
// @Param id path int true "Group ID"
// @Success 200 {array} models.Task
// @Router /group/{id}/tasks [get]
func (task *TaskController) GetToDos(c *gin.Context) {
	groupID := c.Param("id")
	var list []models.Task
	err := taskModel.GetToDos(groupID, &list)
	
	if err != nil {
		c.JSON(500, gin.H{"message": "To Dos could not be retrieved", "error": err.Error()})
		c.Abort()
		return
	}

	c.JSON(200, gin.H{"message": "To dos retrieved", "todos": list})
}

// FinishTask godoc
// @Summary Finish pending task
// @Description finish a pending task
// @Tags tasks
// @Accept  json
// @Produce  json
// @Success 200 {object} models.Task
// @Router /tasks/finish [post]
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