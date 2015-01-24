package main

import (
	"fmt"
	"github.com/gophergala/not_golang_experts/conf"
	"github.com/gophergala/not_golang_experts/models"
)

func main() {
	db := conf.SetupDB()
	db.AutoMigrate(&models.User{})

	fmt.Println("Hello world!!!")
}
