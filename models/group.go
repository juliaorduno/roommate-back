package models

import "time"

type RGroup struct {
	ID 	        uint     	`json:"id"`
	Code		string		`sql:"type:varchar(50);not null;unique" json:"code"`
	Size		int			`json:"size"`
	CreatedBy	int			`sql:"not null" json:"created_by"`	
	Admin		int			`sql:"not null" json:"admin"`
	CreatedAt 	time.Time 	`json:"created_at"`
	UpdatedAt 	time.Time 	`json:"updated_at"`
	DeletedAt 	*time.Time 	`json:"deleted_at"`	
}