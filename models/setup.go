package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	//database, err := gorm.Open("sqlite3", "test.db")

	dsn := "root:admin123@tcp(192.168.16.2:3306)/admin?charset=utf8mb4&parseTime=True&loc=Local"
	//dsn := os.Getenv("DB_CONN_STR")
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database!")
	}

	DB = database
}
