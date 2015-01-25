package router

import (
	"encoding/json"
	"github.com/gophergala/not_golang_experts/model"
	"io"
	"net/http"
)

type Subscription struct {
	Url string `json:"url"`
}

func SubscriptionsIndex(res http.ResponseWriter, req *http.Request) {
	message := map[string]string{"message": "Subscriptions Index"}
	json, err := json.Marshal(message)
	if err != nil {
		panic(err)
	}
	res.Write(json)
}

func SubscriptionsCreate(res http.ResponseWriter, req *http.Request) {
	params := req.URL.Query()
	url, err := parseSubscriptionsRequest(req.Body)
	token := params["token"][0]

	subscription := model.SubscribeUser(url, token)
	message := map[string]interface{}{"Subscription": subscription}
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

func parseSubscriptionsRequest(body io.Reader) (string, error) {
	subscription := Subscription{}

	decoder := json.NewDecoder(body)
	err := decoder.Decode(&subscription)

	return subscription.Url, err
}
