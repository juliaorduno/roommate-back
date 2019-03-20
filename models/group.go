package models

import (
	"time"

	"../db"
)

type RGroup struct {
	ID        uint       `json:"id"`
	Code      string     `sql:"type:varchar(50);not null;unique" json:"code"`
	Name      string     `sql:"type:varchar(50);not null" json:"name"`
	GroupID   int        `sql:"type:int(10)" json:"group_id"`
	Size      int        `json:"size"`
	CreatedBy int        `sql:"not null" json:"created_by"`
	Admin     int        `sql:"not null" json:"admin"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

type GroupModel struct{}

func (m *GroupModel) Find(list *[]RGroup) (err error) {
	if err := db.DB.Find(list).Error; err != nil {
		return err
	}
	return nil
}

func (m *GroupModel) Get(id string, rGroup *RGroup) (err error) {
	if err := db.DB.First(rGroup, id).Error; err != nil {
		return err
	}
	return nil
}

func (m *GroupModel) Create(rGroup *RGroup) (err error) {
	if err := db.DB.Create(rGroup).Error; err != nil {
		return err
	}
	return nil
}