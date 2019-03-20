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
		c.JSON(404, gin.H{"message": "ShoppingItem not found", "error": err.Error()})
		c.Abort()
	} else {
		c.JSON(200, gin.H{"data": data})
	}
}

func (shoppingItem *ShoppingItemController) Create(c *gin.Context) {
	var data models.ShoppingItem
	if c.BindJSON(&data) != nil {
		c.JSON(406, gin.H{"message": "Invalid form", "form": data})
		c.Abort()
		return
	}

	err := shoppingItemModel.Create(&data)
	if err != nil {
		c.JSON(406, gin.H{"message": "ShoppingItem could not be created", "error": err.Error()})
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
		c.JSON(406, gin.H{"message": "Invalid form", "form": data})
		c.Abort()
		return
	}

	groupID := data.GroupID
	var list []models.ShoppingItem
	err := shoppingItemModel.GetShoppingItems(groupID, &list)

	if err != nil {
		c.JSON(406, gin.H{"message": "Shopping items could not be retrieved", "error": err.Error()})
		c.Abort()
		return
	}

	c.JSON(200, gin.H{"message": "Shopping items retrieved", "shopping items": list})
}