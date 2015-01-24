package main

import (
	"github.com/gophergala/not_golang_experts/conf"
	"github.com/gophergala/not_golang_experts/model"
	"github.com/gophergala/not_golang_experts/router"
	"net/http"
)

func main() {
	db := conf.SetupDB()
	db.AutoMigrate(&model.User{}, &model.Page{})

	http.ListenAndServe(":3000", router.GetRoutes())
}
