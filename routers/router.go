package routers

import (
	"net/http"
	"../controllers"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	api := router.Group("/api/v1")
	{
		member := new(controllers.MemberController)
		api.GET("/members", member.Find)
		api.GET("/members/:id", member.Get)
		api.POST("/members", member.Create)
		api.POST("/members/joinGroup", member.JoinGroup)

		rGroup := new(controllers.GroupController)
		api.GET("/groups", rGroup.Find)
		api.GET("/groups/:id", rGroup.Get)
		api.POST("/groups", rGroup.Create)
		api.POST("/groups/createGroup", rGroup.CreateGroup)

		announcement := new(controllers.AnnouncementController)
		api.GET("/announcements", announcement.Find)
		api.GET("/announcements/:id", announcement.Get)
		api.POST("/announcements", announcement.Create)

		meeting := new(controllers.MeetingController)
		api.GET("/meetings", meeting.Find)
		api.GET("/meetings/:id", meeting.Get)
		api.POST("/meetings", meeting.Create)

		shoppingItem := new(controllers.ShoppingItemController)
		api.GET("/shoppingItems", shoppingItem.Find)
		api.GET("/shoppingItems/:id", shoppingItem.Get)
		api.POST("/shoppingItems", shoppingItem.Create)

		task := new(controllers.TaskController)
		api.GET("/tasks", task.Find)
		api.GET("/tasks/:id", task.Get)
		api.POST("/tasks", task.Create)
	}

	router.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "Not Found")
	})

	return router
}