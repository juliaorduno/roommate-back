package models

import "time"

type Task struct {
	ID     			uint		`json:"id"`
	Description		string		`sql:"type:varchar(255);not null" json:"description"`
	Asignee			int			`sql:"type:int(10)" json:"asignee"`
	DueDate			time.Time	`sql:"not null" json:"due_date"`
	Finished		int			`json:"finished"`
	FinishedAt		time.Time	`json:"finished_at"`
	CreatedBy		int 		`sql:"type:int(10); not null" json:"created_by"`
	CreatedAt 		time.Time 	`json:"created_at"`
	UpdatedAt 		time.Time 	`json:"updated_at"`
	DeletedAt 		*time.Time 	`json:"deleted_at"`	
}