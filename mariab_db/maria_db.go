package mariab_db

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const DB_USERNAME = "root"
const DB_PASSWORD = "root"
const DB_NAME = "kvshipping_prelive"
const DB_HOST = "host.docker.internal"
const DB_PORT = "3312"

var Db *gorm.DB

func InitDb() *gorm.DB {
	return connectDB()
}

func connectDB() *gorm.DB {
	var err error
	dsn := DB_USERNAME + ":" + DB_PASSWORD + "@tcp" + "(" + DB_HOST + ":" + DB_PORT + ")/" + DB_NAME + "?" + "parseTime=true&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("Error connecting to database : error=%v", err)
		return nil
	}

	//err = db.AutoMigrate(&models.User{})
	//if err != nil {
	//	return nil
	//}

	return db
}
