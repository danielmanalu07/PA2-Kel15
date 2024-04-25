package database

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connection() {
	godotenv.Load()

	dbhost := os.Getenv("DB_HOST")
	dbuser := os.Getenv("DB_USER")
	dbname := os.Getenv("DB_NAME")

	connect := fmt.Sprintf("%s:@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc-Local", dbuser, dbhost, dbname)

	db, err := gorm.Open(mysql.Open(connect), &gorm.Config{})

	if err != nil {
		panic("Failed to connection Database")
	}

	DB = db

	fmt.Println("Connected Successfully")
}
