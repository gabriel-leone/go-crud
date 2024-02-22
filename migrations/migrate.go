package main

import (
	"github.com/gabriel-leone/go-crud/initializers"
	"github.com/gabriel-leone/go-crud/models"
)

func init() {
	initializers.ConnectToDatabase()
	initializers.LoadEnvVariables()
}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
}
