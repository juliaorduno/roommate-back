package routers

import (
	"time"
	"log"
	"../controllers"
	"../models"
	"github.com/gin-gonic/gin"
	"github.com/appleboy/gin-jwt"
)

type login struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

var identityKey = "id"

func helloHandler(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	user, _ := c.Get(identityKey)
	c.JSON(200, gin.H{
		"userID":   claims["id"],
		"userName": user.(*User).UserName,
		"text":     "Hello World.",
	})
}

// User demo
type User struct {
	UserName  string
	FirstName string
	LastName  string
}

func SetupRouter() *gin.Engine {
	router := gin.Default()

	// the jwt middleware
	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:       "test zone",
		Key:         []byte("secret key"),
		Timeout:     time.Hour,
		MaxRefresh:  time.Hour,
		IdentityKey: identityKey,
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*User); ok {
				return jwt.MapClaims{
					identityKey: v.UserName,
				}
			}
			return jwt.MapClaims{}
		},
		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c)
			return &User{
				UserName: claims["id"].(string),
			}
		},
		Authenticator: func(c *gin.Context) (interface{}, error) {
			var loginVals login
			if err := c.ShouldBind(&loginVals); err != nil {
				return "", jwt.ErrMissingLoginValues
			}
			userID := loginVals.Username
			password := loginVals.Password

			isValid := models.Login(userID, password)

			if (userID == "admin" && password == "adminn") || (userID == "test" && password == "test" || isValid == true) {
				return &User{
					UserName:  userID,
					LastName:  "Bo-Yi",
					FirstName: "Wu",
				}, nil
			}

			return nil, jwt.ErrFailedAuthentication
		},
		Authorizator: func(data interface{}, c *gin.Context) bool {
			if v, ok := data.(*User); ok && v.UserName == "admin" {
				return true
			}

			return true
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{
				"code":    code,
				"message": message,
			})
		},
		// TokenLookup is a string in the form of "<source>:<name>" that is used
		// to extract token from the request.
		// Optional. Default value "header:Authorization".
		// Possible values:
		// - "header:<name>"
		// - "query:<name>"
		// - "cookie:<name>"
		// - "param:<name>"
		TokenLookup: "header: Authorization, query: token, cookie: jwt",
		// TokenLookup: "query:token",
		// TokenLookup: "cookie:token",

		// TokenHeadName is a string in the header. Default value is "Bearer"
		TokenHeadName: "Bearer",

		// TimeFunc provides the current time. You can override it to use another time value. This is useful for testing or if your server uses a different time zone than your tokens.
		TimeFunc: time.Now,
	})

	if err != nil {
		log.Fatal("JWT Error:" + err.Error())
	}

	router.NoRoute(authMiddleware.MiddlewareFunc(), func(c *gin.Context) {
		claims := jwt.ExtractClaims(c)
		log.Printf("NoRoute claims: %#v\n", claims)
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})

	api := router.Group("/api/v1")
	{
		member := new(controllers.MemberController)
		rGroup := new(controllers.GroupController)
		announcement := new(controllers.AnnouncementController)
		meeting := new(controllers.MeetingController)
		shoppingItem := new(controllers.ShoppingItemController)
		task := new(controllers.TaskController)
		user := new(controllers.UserController)

		api.POST("/login", authMiddleware.LoginHandler)
		api.POST("/users", user.Create)

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

	auth := router.Group("/auth")
	// Refresh time can be longer than token timeout
	auth.GET("/refresh_token", authMiddleware.RefreshHandler)
	auth.Use(authMiddleware.MiddlewareFunc())
	{
		auth.GET("/hello", helloHandler)
	}

	return router
}