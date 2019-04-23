package models

import (
	"time"
	
	"../db"
)

type Announcement struct {
	ID     			uint			`json:"id"`
	GroupID			uint			`json:"group_id"`
	Content			string			`sql:"type:varchar(255);not null" json:"content"`
	CreatedBy		uint			`sql:"not null" json:"created_by"`
	CreatedAt		time.Time		`json:"created_at"`
	UpdatedAt 		time.Time 		`json:"updated_at"`
	DeletedAt 		*time.Time 		`json:"deleted_at"`	
}

type AnnouncementModel struct{}

func (m *AnnouncementModel) Find(list *[]Announcement) ( err error) {
	if err := db.DB.Find(list).Error; err != nil {
		return err
	}
	return nil
}

func (m *AnnouncementModel) Get(id string, announcement *Announcement) (err error) {
	if err := db.DB.First(announcement, id).Error; err != nil {
		return err
	}
	return nil
}

func (m *AnnouncementModel) Create(announcement *Announcement) (err error) {
	if err := db.DB.Create(announcement).Error; err != nil {
		return err
	}
	return nil
}

func (m *AnnouncementModel) GetAnnouncements (groupID string, list *[]Announcement)(err error){
	if err := db.DB.Where("group_id = ?", groupID).Find(list).Error; err != nil {
		return err
	}
	return nil
}