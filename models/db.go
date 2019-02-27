package models

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func DbConn() (*gorm.DB) {
	var myEnv map[string]string
	myEnv, err := godotenv.Read()

	mysqlCredentials := fmt.Sprintf(
		"%s:%s@/%s?charset=utf8&parseTime=True&loc=Local",
		myEnv["DB_USER"],
		myEnv["DB_PASSWORD"],
		myEnv["DB_NAME"],
	)

	db, err := gorm.Open(myEnv["DB_DRIVER"], mysqlCredentials)

    if err != nil {
        panic(err.Error())
	}

	db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(
		&Member{},
		&RGroup{},
	)
	
	//db.Model(&Member{}).AddForeignKey("group_id", "r_groups(id)","RESTRICT", "RESTRICT")

	return db
}