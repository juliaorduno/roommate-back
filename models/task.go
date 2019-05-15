package models

import (
	"time"
	"../db"
)

type Task struct {
	ID     			uint		`json:"id"`
	Description		string		`sql:"type:varchar(255);not null" json:"description"`
	GroupID			uint		`json:"group_id"`
	Asignee			uint		`json:"asignee_id"`
	Member			Member	`gorm:"foreignkey:Asignee;association_foreignkey:ID" json:"asignee"`
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

	log := Log{
		GroupID: 	 task.GroupID,
		CreatedBy: task.CreatedBy,
		ItemType:  4,
		Item:			 task.Description,
	}

	err = logModel.Create(&log)

	if err != nil{
		return err
	}
	
	return nil
}

func (m *TaskModel) GetToDos(groupID string) ( err error, list []Task) {
	if err := db.DB.Where("group_id = ? AND finished = ?", groupID, 0).Find(&list).Error; err != nil {
		return err, nil
	}

	for i := 0; i < len(list); i++ {
		var member Member

		if err = db.DB.First(&member, list[i].Asignee).Error; err !=  nil {
			return err, nil
		}

		list[i].Member = member
	}
	return nil, list
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