package controllers

import (
	"../models"
	"github.com/gin-gonic/gin"
)

var shoppingItemModel = new(models.ShoppingItemModel)

type ShoppingItemController struct{}

// ListShoppingItems godoc
// @Summary List shopping items
// @Description get shopping items
// @Tags shoppingItems
// @Accept  json
// @Produce  json
// @Success 200 {array} models.ShoppingItem
// @Router /shoppingItems [get]
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

// GetShoppingItem godoc
// @Summary Get a shopping item
// @Description get shopping item by ID
// @Tags shoppingItems
// @Accept  json
// @Produce  json
// @Param id path int true "ShoppingItem ID"
// @Success 200 {object} models.ShoppingItem
// @Router /shoppingItems/{id} [get]
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

// CreateShoppingItem godoc
// @Summary Create a shopping item
// @Description create by json shopping item
// @Tags shoppingItems
// @Accept  json
// @Produce  json
// @Param shoppingItem body models.ShoppingItem true "Add ShoppingItem"
// @Success 200 {object} models.ShoppingItem
// @Router /shoppingItems [post]
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

// ListGroupShoppingItems godoc
// @Summary List group pending shopping items
// @Description get group pending shopping items
// @Tags groups
// @Accept  json
// @Produce  json
// @Param id path int true "Group ID"
// @Success 200 {array} models.ShoppingItem
// @Router /group/{id}/shoppingItems [get]
func (shoppingItem *ShoppingItemController) GetShoppingItems(c *gin.Context) {
	groupID := c.Param("id")
	err, list := shoppingItemModel.GetShoppingItems(groupID)

	if err != nil {
		c.JSON(500, gin.H{"message": "Shopping items could not be retrieved", "error": err.Error()})
		c.Abort()
		return
	}

	c.JSON(200, gin.H{"message": "Shopping items retrieved", "shopping_items": list})
}

// FinishShoppingItem godoc
// @Summary Finish pending shopping item
// @Description finish a pending shopping item
// @Tags shoppingItems
// @Accept  json
// @Produce  json
// @Success 200 {object} models.ShoppingItem
// @Router /shoppingItems/finish [post]
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