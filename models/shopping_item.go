package models

import (
	"time"

	"../db"
)

type ShoppingItem struct {
	ID          uint       `json:"id"`
	Description string     `sql:"type:varchar(255);not null" json:"description"`
	CreatedBy   uint        `json:"created_by"`
	GroupID     uint        `json:"group_id"`
	Finished    int        `json:"finished"`
	FinishedAt  time.Time  `json:"finished_at"`
	FinishedBy  uint        `json:"finished_by"`
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

func (m *ShoppingItemModel) GetShoppingItems(groupID string, list *[]ShoppingItem) (err error) {
	if err := db.DB.Where("group_id = ? AND finished = ?", groupID, 0).Find(list).Error; err != nil {
		return err
	}
	return nil
}

func (m *ShoppingItemModel) FinishItem(id int, finished_by uint, shoppingItem *ShoppingItem) (err error) {
	if err := db.DB.First(shoppingItem, id).Error; err != nil {
		return err
	}

	shoppingItem.Finished = 1
	shoppingItem.FinishedBy = finished_by
	shoppingItem.FinishedAt = time.Now()

	if err := db.DB.Save(shoppingItem).Error; err != nil {
		return err
	}
	return nil
}
