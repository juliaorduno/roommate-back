package models

import (
	"time"
	"../db"
)

type Task struct {
	ID     			uint		`json:"id"`
	Description		string		`sql:"type:varchar(255);not null" json:"description"`
	GroupID			uint		`json:"group_id"`
	Asignee			uint		`json:"asignee"`
	DueDate			time.Time	`sql:"not null" json:"due_date"`
	Finished		int			`json:"finished"`
	FinishedBy		uint		`sql:"default:null" json:"finished_by"`
	FinishedAt		time.Time	`json:"finished_at"`
	CreatedBy		uint 		`sql:"not null" json:"created_by"`
	CreatedAt 		time.Time 	`json:"created_at"`
	UpdatedAt 		time.Time 	`json:"updated_at"`
	DeletedAt 		*time.Time 	`json:"deleted_at"`	
}

type TaskModel struct{}

func (m *TaskModel) Find(list *[]Task) ( err error) {
	if err := db.DB.Find(list).Error; err != nil {
		return err
	}
	return nil
}

func (m *TaskModel) Get(id string, task *Task) (err error) {
	if err := db.DB.First(task, id).Error; err != nil {
		return err
	}
	return nil
}

func (m *TaskModel) Create(task *Task) (err error) {
	if err := db.DB.Create(task).Error; err != nil {
		return err	
	}
	return nil
}

func (m *TaskModel) GetToDos(groupID string, list *[]Task) ( err error) {
	if err := db.DB.Where("group_id = ? AND finished = ?", groupID, 0).Find(list).Error; err != nil {
		return err
	}
	return nil
}

func (m *TaskModel) FinishTask(id int, finished_by uint, task *Task) (err error) {
	if err := db.DB.First(task, id).Error; err != nil {
		return err
	}

	task.Finished = 1
	task.FinishedBy = finished_by
	task.FinishedAt = time.Now()

	if err := db.DB.Save(task).Error; err != nil {
		return err
	}
	return nil
}