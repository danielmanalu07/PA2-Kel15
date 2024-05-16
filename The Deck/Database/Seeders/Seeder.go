package seeders

import (
	database "api/the_deck/Database"
	"api/the_deck/Models/entity"
	utils "api/the_deck/Utils"
	"fmt"
	"log"
)

func SeederAdmin() {
	password, err := utils.GeneratePassword("admin12345")
	if err != nil {
		log.Fatalf(err.Error())
	}

	admin := &entity.Admin{
		Username: "admin",
		Password: password,
	}

	if err := database.DB.Create(&admin); err != nil {
		log.Fatalf("Failed to create admin: %v", err)
	}

	fmt.Println("Data Seeded Successfully")
}
