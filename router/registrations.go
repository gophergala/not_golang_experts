package router

import (
	"encoding/json"
	"github.com/gophergala/not_golang_experts/model"
	"io"
	"net/http"
)

type Registration struct {
	User User `json:"user"`
}

type User struct {
	Email                string `json:"email"`
	Password             string `json:"password"`
	PasswordConfirmation string `json:"password_confirmation"`
}

func RegisterSession(res http.ResponseWriter, req *http.Request) {
	email, password, password_confirmation, err := parseRequest(req.Body)

	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}

	model.RegisterUser(email, password, password_confirmation, func(token string) {
		respondWith(map[string]string{"token": token}, 201, res)
	}, func(message string) {
		respondWith(map[string]string{"error": message}, 422, res)
	})
}

func respondWith(json_map map[string]string, status int, res http.ResponseWriter) {
	json_response, err := json.Marshal(json_map)
	PanicIf(err)
	res.Header().Set("Content-Type", "application/json; charset=UTF-8")
	res.WriteHeader(status)
	res.Write(json_response)
}

func parseRequest(body io.Reader) (string, string, string, error) {
	registration := Registration{}

	decoder := json.NewDecoder(body)
	err := decoder.Decode(&registration)

	return registration.User.Email, registration.User.Password, registration.User.PasswordConfirmation, err
}
