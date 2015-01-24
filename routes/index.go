package routes

import (
	"encoding/json"
	"net/http"
)

func Index(res http.ResponseWriter, req *http.Request) {
	message := map[string]string{"message": "Hello world!"}
	json, err := json.Marshal(message)
	if err != nil {
		panic(err)
	}
	res.Write(json)
}
