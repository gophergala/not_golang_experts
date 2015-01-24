package main

import (
	"github.com/gophergala/not_golang_experts/conf"
	"github.com/gophergala/not_golang_experts/model"
	"github.com/gophergala/not_golang_experts/router"
	"github.com/gophergala/not_golang_experts/worker"
	"net/http"
)

func main() {
	db := conf.SetupDB()
	db.AutoMigrate(&model.User{}, &model.Page{})

	model.Db = db

	stopped := make(chan bool, 1)
	worker.StartObserving(stopped)

	http.ListenAndServe(":3000", router.GetRoutes())

	<-stopped
}
