package main

import (
	"fmt"
	"log"

	"github.com/engtiago/go-boilerplate/initializers"
	"github.com/engtiago/go-boilerplate/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	err := initializers.DB.AutoMigrate(&models.User{})

	if err != nil {
		log.Fatal("Migrate error")
	}
	fmt.Println("Migrate success")
}
