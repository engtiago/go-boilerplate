package main

import (
	"fmt"
	"log"

	"github.com/engtiago/go-boilerplate/src/initializers"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {

	for _, model := range RegisterModels() {
		err := initializers.DB.AutoMigrate(model.Model)
		if err != nil {
			log.Fatal("Migrate error")
		}
	}

	fmt.Println("Migrate success")
}
