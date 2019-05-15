package main

import (
	//"net/http"

	"./db"
	"./models"
	"./routers"
	_ "./docs"

	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles" // swagger embed files
)

// @title Roommate API
// @version 1.0
// @description This is the API corresponding to a Roommate server.

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:3030
// @BasePath /api
func main() {
	db.DB = db.DbConn()
	defer db.DB.Close()

	db.DB.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(
		&models.Member{},
		&models.RGroup{},
		&models.Announcement{},
		&models.Meeting{},
		&models.ShoppingItem{},
		&models.Task{},
		&models.User{},
		&models.Log{},
	)

	db.DB.Model(&models.Member{}).AddForeignKey("group_id", "r_groups(id)","RESTRICT", "RESTRICT")
	db.DB.Model(&models.Announcement{}).AddForeignKey("group_id", "r_groups(id)","RESTRICT", "RESTRICT")
	db.DB.Model(&models.Announcement{}).AddForeignKey("created_by", "members(id)","RESTRICT", "RESTRICT")
	db.DB.Model(&models.Meeting{}).AddForeignKey("group_id", "r_groups(id)","RESTRICT", "RESTRICT")
	db.DB.Model(&models.Meeting{}).AddForeignKey("created_by", "members(id)","RESTRICT", "RESTRICT")
	db.DB.Model(&models.ShoppingItem{}).AddForeignKey("group_id", "r_groups(id)","RESTRICT", "RESTRICT")
	db.DB.Model(&models.ShoppingItem{}).AddForeignKey("created_by", "members(id)","RESTRICT", "RESTRICT")
	db.DB.Model(&models.ShoppingItem{}).AddForeignKey("finished_by", "members(id)","RESTRICT", "RESTRICT")
	db.DB.Model(&models.Task{}).AddForeignKey("group_id", "r_groups(id)","RESTRICT", "RESTRICT")
	db.DB.Model(&models.Task{}).AddForeignKey("created_by", "members(id)","RESTRICT", "RESTRICT")
	db.DB.Model(&models.Task{}).AddForeignKey("asignee", "members(id)","RESTRICT", "RESTRICT")
	db.DB.Model(&models.Task{}).AddForeignKey("finished_by", "members(id)","RESTRICT", "RESTRICT")
	db.DB.Model(&models.Log{}).AddForeignKey("group_id", "r_groups(id)","RESTRICT", "RESTRICT")
	db.DB.Model(&models.Log{}).AddForeignKey("created_by", "members(id)","RESTRICT", "RESTRICT")

	router := routers.SetupRouter()

	config := &ginSwagger.Config{
		URL: "http://localhost:3030/swagger/doc.json", //The url pointing to API definition
	}
	// use ginSwagger middleware to 
	router.GET("/swagger/*any", ginSwagger.CustomWrapHandler(config, swaggerFiles.Handler))
	router.Run(":3030")
}