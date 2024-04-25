package migration

import (
	"fmt"
	"log"
	"pa2/database"
	"pa2/models/entity"
)

func Migration() {
	err := database.DB.AutoMigrate(&entity.Admin{}, &entity.Category{}, &entity.Product{})

	if err != nil {
		log.Println(err)
	}

	fmt.Println("Database migrated successfully")
}
