package migration

import (
	"admin/database"
	"admin/models/entity"
	"fmt"
	"log"
)

func Migration() {

	err := database.DB.AutoMigrate(&entity.Admin{})

	if err != nil {
		log.Println(err)
	}
	fmt.Println("Database Migrated")
}
