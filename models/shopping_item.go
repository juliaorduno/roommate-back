package models

import "time"

type ShoppingItem struct {
	ID     			uint			`json:"id"`
	Description		string			`sql:"type:varchar(255);not null" json:"description"`
	CreatedBy		int				`sql:"type:int(10)" json:"created_by"`
	Finished		int 			`json:"finished"`
	FinishedAt		time.Time		`json:"finished_at"`
	FinishedBy		int				`sql:"type:int(10)" json:"finished_by"`
	CreatedAt 		time.Time 		`json:"created_at"`
	UpdatedAt 		time.Time 		`json:"updated_at"`
	DeletedAt 		*time.Time 		`json:"deleted_at"`	
}