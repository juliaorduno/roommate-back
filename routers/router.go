package routers

import (
	"../controllers"

	"github.com/gin-gonic/gin"
	"github.com/appleboy/gin-jwt"
	"github.com/gin-contrib/cors"
	"log"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.Use(cors.Default())

	authMiddleware := SetupMiddleware()

	member := new(controllers.MemberController)
	rGroup := new(controllers.GroupController)
	announcement := new(controllers.AnnouncementController)
	meeting := new(controllers.MeetingController)
	shoppingItem := new(controllers.ShoppingItemController)
	task := new(controllers.TaskController)
	user := new(controllers.UserController)

	api := router.Group("/api")
	{
		api.POST("/login", authMiddleware.LoginHandler)
		api.POST("/users", user.Create)
		api.GET("/members", member.Find)

		api.GET("/members/:email", member.GetByEmail)
		api.POST("/members", member.Create)
		//api.GET("/current/:email", member.GetByEmail)
		api.POST("/members/joinGroup", member.JoinGroup)

		api.GET("/groups", rGroup.Find)
		api.GET("/groups/:id", rGroup.Get)
		api.POST("/groups", rGroup.CreateGroup)
		api.GET("/groups/:id/tasks", task.GetToDos)
		api.GET("/groups/:id/meetings", meeting.GetMeetings)
		api.GET("/groups/:id/shoppingItems", shoppingItem.GetShoppingItems)
		api.GET("/groups/:id/announcements", announcement.GetAnnouncements)

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

	router.NoRoute(authMiddleware.MiddlewareFunc(), func(c *gin.Context) {
		claims := jwt.ExtractClaims(c)
		log.Printf("NoRoute claims: %#v\n", claims)
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})

	// Refresh time can be longer than token timeout
	api.GET("/refresh_token", authMiddleware.RefreshHandler)
	api.Use(authMiddleware.MiddlewareFunc())
	{
		api.GET("/users/currentUser", user.GetCurrentUser)

		//api.GET("/hello", helloHandler)
	}

	return router
}