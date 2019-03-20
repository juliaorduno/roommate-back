package models

import (
	"time"

	"../db"
)

type ShoppingItem struct {
	ID          uint       `json:"id"`
	Description string     `sql:"type:varchar(255);not null" json:"description"`
	CreatedBy   int        `sql:"type:int(10)" json:"created_by"`
	GroupID     int        `sql:"type:int(10)" json:"group_id"`
	Finished    int        `json:"finished"`
	FinishedAt  time.Time  `json:"finished_at"`
	FinishedBy  int        `sql:"type:int(10)" json:"finished_by"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at"`
}

type ShoppingItemModel struct{}

func (m *ShoppingItemModel) Find(list *[]ShoppingItem) (err error) {
	if err := db.DB.Find(list).Error; err != nil {
		return err
	}
	return nil
}

func (m *ShoppingItemModel) Get(id string, shoppingItem *ShoppingItem) (err error) {
	if err := db.DB.First(shoppingItem, id).Error; err != nil {
		return err
	}
	return nil
}

func (m *ShoppingItemModel) Create(shoppingItem *ShoppingItem) (err error) {
	if err := db.DB.Create(shoppingItem).Error; err != nil {
		return err
	}
	return nil
}

func (m *ShoppingItemModel) GetShoppingItems(groupID int, list *[]ShoppingItem) (err error) {
	if err := db.DB.Where("group_id = ?", groupID).Find(list).Error; err != nil {
		return err
	}
	return nil
}
