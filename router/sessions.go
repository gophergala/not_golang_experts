package router

import (
	"encoding/json"
	"github.com/gophergala/not_golang_experts/model"
	"io"
	"net/http"
)

type Session struct {
	User UserRegistration `json:"user"`
}

type UserSession struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func CreateSession(res http.ResponseWriter, req *http.Request) {
	email, password, err := parseSessionsRequest(req.Body)

	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}

	model.RegisterUserSession(email, password, func(token string) {
		respondWith(map[string]string{"token": token}, 201, res)
	}, func(message string) {
		respondWith(map[string]string{"error": message}, 422, res)
	})
}

func DestroySession(res http.ResponseWriter, req *http.Request) {
}

func parseSessionsRequest(body io.Reader) (string, string, error) {
	registration := Registration{}

	decoder := json.NewDecoder(body)
	err := decoder.Decode(&registration)

	return registration.User.Email, registration.User.Password, err
}
