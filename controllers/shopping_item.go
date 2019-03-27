package controllers

import (
	"../models"
	"github.com/gin-gonic/gin"
)

var shoppingItemModel = new(models.ShoppingItemModel)

type ShoppingItemController struct{}

func (shoppingItem *ShoppingItemController) Find(c *gin.Context) {
	var list []models.ShoppingItem
	err := shoppingItemModel.Find(&list)
	if err != nil {
		c.JSON(404, gin.H{"message": "Find Error", "error": err.Error()})
		c.Abort()
	} else {
		c.JSON(200, gin.H{"data": list})
	}
}

func (shoppingItem *ShoppingItemController) Get(c *gin.Context) {
	id := c.Param("id")
	var data models.ShoppingItem
	err := shoppingItemModel.Get(id, &data)
	if err != nil {
		c.JSON(404, gin.H{"message": "Shopping Item not found", "error": err.Error()})
		c.Abort()
	} else {
		c.JSON(200, gin.H{"data": data})
	}
}

func (shoppingItem *ShoppingItemController) Create(c *gin.Context) {
	var data models.ShoppingItem
	if c.BindJSON(&data) != nil {
		c.JSON(400, gin.H{"message": "Invalid form", "form": data})
		c.Abort()
		return
	}

	err := shoppingItemModel.Create(&data)
	if err != nil {
		c.JSON(500, gin.H{"message": "ShoppingItem could not be created", "error": err.Error()})
		c.Abort()
		return
	}

	c.JSON(200, gin.H{"message": "ShoppingItem Created", "new_shopping_item": data})
}

func (shoppingItem *ShoppingItemController) GetShoppingItems(c *gin.Context) {
	var data struct {
		GroupID int `json:"group_id"`
	}
	if c.BindJSON(&data) != nil {
		c.JSON(400, gin.H{"message": "Invalid form", "form": data})
		c.Abort()
		return
	}

	groupID := data.GroupID
	var list []models.ShoppingItem
	err := shoppingItemModel.GetShoppingItems(groupID, &list)

	if err != nil {
		c.JSON(500, gin.H{"message": "Shopping items could not be retrieved", "error": err.Error()})
		c.Abort()
		return
	}

	c.JSON(200, gin.H{"message": "Shopping items retrieved", "shopping items": list})
}

func (shoppingItem *ShoppingItemController) FinishItem(c *gin.Context) {
	var data struct {
		ID int `json:"id"`
		FinishedBy uint `json:"finished_by"`
	}

	if c.BindJSON(&data) != nil {
		c.JSON(400, gin.H{"message": "Invalid form", "form": data})
		c.Abort()
		return
	}

	if data.ID <= 0 {
		c.JSON(400, gin.H{"message": "Invalid form", "error": "You must add an ID"})
		c.Abort()
		return
	}

	if data.FinishedBy <= 0 {
		c.JSON(400, gin.H{"message": "Invalid form", "error": "You must add a memberID"})
		c.Abort()
		return
	}
	
	var updatedItem models.ShoppingItem

	err := shoppingItemModel.FinishItem(data.ID, data.FinishedBy, &updatedItem)
	if err != nil {
		c.JSON(500, gin.H{"message": "ShoppingItem could not be updated", "error": err.Error()})
		c.Abort()
		return
	}

	c.JSON(200, gin.H{"message": "Shopping Item Created", "new_shopping_item": updatedItem})
} 