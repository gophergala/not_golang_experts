package router

import (
	"github.com/gorilla/mux"
	"net/http"
)

func GetRoutes() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", BaseHandler(Index))

	return r
}

func BaseHandler(fn func(w http.ResponseWriter, r *http.Request)) http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		fn(w, r)
	}
}
