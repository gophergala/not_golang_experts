package router

import (
	"github.com/gorilla/mux"
	"net/http"
)

func GetRoutes() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", BaseHandler(Index))

	r.Path("/registrations").Subrouter().Methods("POST").HandlerFunc(BaseHandler(RegisterSession))

	r.Path("/sessions").Subrouter().Methods("POST").HandlerFunc(BaseHandler(CreateSession))
	r.Path("/sessions").Subrouter().Methods("DELETE").HandlerFunc(BaseHandler(DestroySession))

	// Serve static assets

	r.Handle("/public/javascripts/{rest}", http.StripPrefix("/public/", http.FileServer(http.Dir("public/"))))
	r.Handle("/public/stylesheets/{rest}", http.StripPrefix("/public/", http.FileServer(http.Dir("public/"))))
	r.Handle("/public/images/{rest}", http.StripPrefix("/public/", http.FileServer(http.Dir("public/"))))

	return r
}

func BaseHandler(fn func(w http.ResponseWriter, r *http.Request)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		fn(w, r)
	}
}
