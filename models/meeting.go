package models

import (
	"time"

	"../db"
)

type Meeting struct {
	ID        uint       `json:"id"`
	Event     string     `sql:"type:varchar(255);not null" json:"event"`
	GroupID   uint        `json:"group_id"`
	StartDate time.Time  `sql:"not null" json:"start_date"`
	EndDate   time.Time  `json:"end_date"`
	CreatedBy uint        `sql:"not null" json:"created_by"`
	Member		Member			`gorm:"foreignkey:CreatedBy;association_foreignkey:ID" json:"organizer"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

type MeetingModel struct{}

func (m *MeetingModel) Find(list *[]Meeting) (err error) {
	if err := db.DB.Find(list).Error; err != nil {
		return err
	}
	return nil
}

func (m *MeetingModel) Get(id string, meeting *Meeting) (err error) {
	if err := db.DB.First(meeting, id).Error; err != nil {
		return err
	}
	return nil
}

func (m *MeetingModel) Create(meeting *Meeting) (err error) {
	if err := db.DB.Create(meeting).Error; err != nil {
		return err
	}

	log := Log{
		GroupID: 	 meeting.GroupID,
		CreatedBy: meeting.CreatedBy,
		ItemType:  2,
		Item:			 meeting.Event,
	}

	err = logModel.Create(&log)

	if err != nil{
		return err
	}

	return nil
}

func (m *MeetingModel) GetMeetings(groupID string) (err error, list []Meeting) {
	if err := db.DB.Where("group_id = ? AND start_date > ?", groupID, time.Now()).Find(&list).Error; err != nil {
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
