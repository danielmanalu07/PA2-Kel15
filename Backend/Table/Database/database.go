package database

import (
	"fmt"
	"service/table/Models/entity"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connection() {
	const connect = "root@tcp(localhost)/service_table?charset=utf8&parseTime=True&loc=Local"
	dsn := connect
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Couldn't connect to database")
	}

	DB = db

	fmt.Println("Successfully Connect to Database")

	DB.AutoMigrate(&entity.Table{})
}
