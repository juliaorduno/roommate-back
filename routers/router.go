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
		rGroup := new(controllers.GroupController)
		announcement := new(controllers.AnnouncementController)
		meeting := new(controllers.MeetingController)
		shoppingItem := new(controllers.ShoppingItemController)
		task := new(controllers.TaskController)

		api.GET("/members", member.Find)
		api.GET("/members/:id", member.Get)
		api.POST("/members", member.Create)
		api.POST("/members/joinGroup", member.JoinGroup)

		api.GET("/groups", rGroup.Find)
		api.GET("/groups/:id", rGroup.Get)
		api.POST("/groups", rGroup.Create)
		api.POST("/groups/createGroup", rGroup.CreateGroup)
		api.POST("/groups/getTodos", task.GetToDos)
		api.POST("/groups/getMeetings", meeting.GetMeetings)
		api.POST("/groups/getShoppingItems", shoppingItem.GetShoppingItems)
		api.POST("/groups/getAnnouncements", announcement.GetAnnouncements)
		
		api.GET("/announcements", announcement.Find)
		api.GET("/announcements/:id", announcement.Get)
		api.POST("/announcements", announcement.Create)
		
		api.GET("/meetings", meeting.Find)
		api.GET("/meetings/:id", meeting.Get)
		api.POST("/meetings", meeting.Create)

		api.GET("/shoppingItems", shoppingItem.Find)
		api.GET("/shoppingItems/:id", shoppingItem.Get)
		api.POST("/shoppingItems", shoppingItem.Create)
		api.POST("/shoppingItems/finish", shoppingItem.FinishItem)

		api.GET("/tasks", task.Find)
		api.GET("/tasks/:id", task.Get)
		api.POST("/tasks", task.Create)
		api.POST("/tasks/finish", task.FinishTask)
	}

	router.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "Not Found")
	})

	return router
}