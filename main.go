package main

import (
	"net/http"
	
	"./controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	api := router.Group("/api/v1")
	{
		member := new(controllers.MemberController)
		api.GET("/members", member.Find)
		api.POST("/members", member.Create)
	}

	router.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "Not Found")
	})

	router.Run(":3030")
}