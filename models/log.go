package models

import (
	"time"
	
	"../db"
)

/*
	1 - announcement
	2 - meeting
	3 - shopping item
	4 - task
*/

type Log struct {
	ID     			uint					`json:"id"`
	GroupID			uint					`json:"group_id"`
	CreatedBy		uint					`sql:"not null" json:"created_by"`
	Member			Member				`gorm:"foreignkey:CreatedBy;association_foreignkey:ID" json:"member"`
	ItemType		int						`json:"type"`
	Item				string				`json:"item"`
	CreatedAt		time.Time			`json:"created_at"`
	UpdatedAt 	time.Time 		`json:"updated_at"`
	DeletedAt 	*time.Time 		`json:"deleted_at"`	
}

type LogModel struct{}

var logModel = new(LogModel)

func (m *LogModel) Create(log *Log) (err error) {
	if err := db.DB.Create(log).Error; err != nil {
		return err
	}

	return nil
}

func (m *LogModel) GetLogs (groupID string)(err error, list []Log){
	if err := db.DB.Where("group_id = ?", groupID).Order("created_at desc").Find(&list).Error; err != nil {
		return err, nil
	}

	for i := 0; i < len(list); i++ {
		var member Member

		if err = db.DB.First(&member, list[i].CreatedBy).Error; err !=  nil {
			return err, nil
		}

		list[i].Member = member
	}

	return nil, list
}