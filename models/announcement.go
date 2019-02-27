package models

import "time"

type Announcement struct {
	ID     			uint			`json:"id"`
	Content			string			`sql:"type:varchar(255);not null" json:"content"`
	CreatedBy		int			    `sql:"type:int(10);not null" json:"created_by"`
	CreatedAt		time.Time		`json:"created_at"`
	UpdatedAt 		time.Time 	`json:"updated_at"`
	DeletedAt 		*time.Time 	`json:"deleted_at"`	
}