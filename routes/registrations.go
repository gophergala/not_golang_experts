package routes

import (
	"encoding/json"
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
	var json_response []byte
	var status int

	email, password, password_confirmation, err := parseRequest(req.Body)

	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}

	if passwordsMatch(password, password_confirmation) {
		message := map[string]string{"message": email}
		json_response, err = json.Marshal(message)
		status = 201
		PanicIf(err)
	} else {
		message := map[string]string{"error": "Invalid login information"}
		json_response, err = json.Marshal(message)
		status = 422
		PanicIf(err)
	}

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

func passwordsMatch(password string, password_confirmation string) bool {
	return password == password_confirmation && password != ""
}
