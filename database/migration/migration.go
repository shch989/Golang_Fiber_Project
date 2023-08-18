package migration

import (
	"fmt"
	"log"

	"github.com/shch989/Golang_Fiber_Project/database"
	"github.com/shch989/Golang_Fiber_Project/models/entity"
)

func RunMigration() {
	err := database.DB.AutoMigrate(&entity.User{})
	if err != nil {
		log.Println(err)
	}
	fmt.Println("Database Migration")
}
