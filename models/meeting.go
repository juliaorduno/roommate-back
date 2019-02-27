package models

import "time"

type Meeting struct {
	ID     			uint		`json:"id"`
	Event			string		`sql:"type:varchar(255);not null" json:"event"`
	StartDate		time.Time	`sql:"not null" json:"start_date"`
	EndDate			time.Time	`json:"end_date"`
	CreatedBy		int 		`sql:"type:int(10);not null" json:"created_by"`
	CreatedAt 		time.Time 	`json:"created_at"`
	UpdatedAt 		time.Time 	`json:"updated_at"`
	DeletedAt 		*time.Time 	`json:"deleted_at"`	
}