package models

import (
	"time"

	"../db"
)

type Meeting struct {
	ID        uint       `json:"id"`
	Event     string     `sql:"type:varchar(255);not null" json:"event"`
	GroupID   uint        ` json:"group_id"`
	StartDate time.Time  `sql:"not null" json:"start_date"`
	EndDate   time.Time  `json:"end_date"`
	CreatedBy uint        `sql:"not null" json:"created_by"`
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
	return nil
}

func (m *MeetingModel) GetMeetings(groupID string, list *[]Meeting) (err error) {
	if err := db.DB.Where("group_id = ? AND start_date > ?", groupID, time.Now()).Find(list).Error; err != nil {
		return err
	}
	return nil
}
