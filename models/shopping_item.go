package models

import (
	"time"

	"../db"
)

type ShoppingItem struct {
	ID          uint       `json:"id"`
	Description string     `sql:"type:varchar(255);not null" json:"description"`
	CreatedBy   uint        `json:"created_by"`
	Member			Member			`gorm:"foreignkey:CreatedBy;association_foreignkey:ID" json:"added_by"`
	GroupID     uint        `json:"group_id"`
	Finished    int        `json:"finished"`
	FinishedAt  time.Time  `json:"finished_at"`
	FinishedBy  uint        `sql:"default:null" json:"finished_by"`
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

	log := Log{
		GroupID: 	 shoppingItem.GroupID,
		CreatedBy: shoppingItem.CreatedBy,
		ItemType:  3,
		Item:			 shoppingItem.Description,
	}

	err = logModel.Create(&log)

	if err != nil{
		return err
	}

	return nil
}

func (m *ShoppingItemModel) GetShoppingItems(groupID string,) (err error, list []ShoppingItem) {
	if err := db.DB.Where("group_id = ? AND finished = ?", groupID, 0).Find(&list).Error; err != nil {
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
