package database

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connection() {
	const con = "root@tcp(localhost)/the_deck?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := con
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Couldn't connect to database")
	}

	DB = db

	fmt.Println("Successfully Connect to database")
}
