package main

import (
	"github.com/gophergala/not_golang_experts/conf"
	"github.com/gophergala/not_golang_experts/model"
	"github.com/gophergala/not_golang_experts/routes"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	db := conf.SetupDB()
	db.AutoMigrate(&model.User{})

	router := mux.NewRouter()
	router.HandleFunc("/", routes.Index)
	http.ListenAndServe(":3000", router)
}
