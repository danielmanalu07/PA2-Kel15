package migrations

import (
	database "api/the_deck/Database"
	"api/the_deck/Models/entity"
	"fmt"
	"log"
)

func Migration() {
	err := database.DB.AutoMigrate(&entity.Admin{}, &entity.Category{}, &entity.Product{}, &entity.Table{})

	if err != nil {
		log.Println(err)
	}
	fmt.Println("Database Migrated")
}
