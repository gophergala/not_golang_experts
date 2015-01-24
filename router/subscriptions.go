package router

import (
	"encoding/json"
	"net/http"
)

func SubscriptionsIndex(res http.ResponseWriter, req *http.Request) {
	message := map[string]string{"message": "Subscriptions Index"}
	json, err := json.Marshal(message)
	if err != nil {
		panic(err)
	}
	res.Write(json)
}

func SubscriptionsCreate(res http.ResponseWriter, req *http.Request) {
	message := map[string]string{"message": "Subscriptions Create"}
	json, err := json.Marshal(message)
	if err != nil {
		panic(err)
	}
	res.Write(json)
}

func SubscriptionsDestroy(res http.ResponseWriter, req *http.Request) {
	message := map[string]string{"message": "Subscriptions Destroy"}
	json, err := json.Marshal(message)
	if err != nil {
		panic(err)
	}
	res.Write(json)
}
