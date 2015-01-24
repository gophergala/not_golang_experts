package routes

import (
	"encoding/json"
	"github.com/gophergala/not_golang_experts/model"
	"io"
	"io/ioutil"
	"net/http"
)

func CreateSession(res http.ResponseWriter, req *http.Request) {
	user := model.User{}
	body, err := ioutil.ReadAll(io.LimitReader(req.Body, 1048576))

	PanicIf(err)

	if err := json.Unmarshal(body, &user); err != nil {
		res.Header().Set("Content-Type", "application/json; charset=UTF-8")
		res.WriteHeader(422)
		if err := json.NewEncoder(res).Encode(err); err != nil {
			panic(err)
		}
	}

	message := map[string]string{"token": user.Token}
	json, err := json.Marshal(message)
	PanicIf(err)
	res.Write(json)
}

func DestroySession(res http.ResponseWriter, req *http.Request) {
	message := map[string]string{"message": "Hello world!"}
	json, err := json.Marshal(message)
	PanicIf(err)
	res.Write(json)
}
