package router

import (
	"encoding/json"
	"net/http"
)

func PanicIf(err error) {
	if err != nil {
		panic(err)
	}
}

func respondWith(json_map map[string]interface{}, status int, res http.ResponseWriter) {
	json_response, err := json.Marshal(json_map)
	PanicIf(err)
	res.WriteHeader(status)
	res.Write(json_response)
}
