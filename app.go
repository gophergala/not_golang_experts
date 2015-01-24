package main

import (
	"fmt"
	"github.com/gophergala/not_golang_experts/conf"
	"github.com/gophergala/not_golang_experts/model"
)

func main() {
	db := conf.SetupDB()
	db.AutoMigrate(&model.User{})

	fmt.Println("Hello world!!!")
}
