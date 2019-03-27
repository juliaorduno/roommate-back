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

	router := routers.SetupRouter()

	router.Run(":3030")
}