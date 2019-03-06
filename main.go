package main

import (
	"./db"
	"./models"
	"./routers"
)

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
	)

	//db.DB.Model(&models.Member{}).AddForeignKey("group_id", "r_groups(id)","RESTRICT", "RESTRICT")

	router := routers.SetupRouter()

	router.Run(":3000")
}