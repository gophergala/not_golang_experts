package main

import (
	"github.com/gophergala/not_golang_experts/conf"
	"github.com/gophergala/not_golang_experts/model"
	"github.com/gophergala/not_golang_experts/router"
	"github.com/gophergala/not_golang_experts/worker"
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "3000"
	}

	db := conf.SetupDB()
	db.AutoMigrate(&model.User{}, &model.Page{})

	model.DB = db

	stopchan := make(chan bool, 1)
	worker.StartObserving(stopchan)

	log.Println("Initializing application on port: " + port)

	http.ListenAndServe(":"+port, router.GetRoutes())

	<-stopchan
}
