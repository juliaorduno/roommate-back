package models

import (
	"time"
	"../db"
)

type Task struct {
	ID     			uint		`json:"id"`
	Description		string		`sql:"type:varchar(255);not null" json:"description"`
	GroupID			int			`sql:"type:int(10)" json:"group_id"`
	Asignee			int			`sql:"type:int(10)" json:"asignee"`
	DueDate			time.Time	`sql:"not null" json:"due_date"`
	Finished		int			`json:"finished"`
	FinishedAt		time.Time	`json:"finished_at"`
	CreatedBy		int 		`sql:"type:int(10); not null" json:"created_by"`
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

func (m *TaskModel) GetToDos(groupID int, list *[]Task) ( err error) {
	if err := db.DB.Where("group_id = ?", groupID).Find(list).Error; err != nil {
		return err
	}
	return nil
}