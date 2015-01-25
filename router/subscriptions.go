package router

import (
	"encoding/json"
	"github.com/gophergala/not_golang_experts/model"
	"io"
	"net/http"
)

type SubscriptionParams struct {
	Url string `json:"url"`
}

func SubscriptionsIndex(res http.ResponseWriter, req *http.Request) {
	token := getToken(req)

	model.GetSubscriptionsForUser(token, func(subscriptions []model.Subscription) {
		respondWith(subscriptions, 200, res)
	}, func(message string) {
		respondWith(map[string]string{"error": message}, 401, res)
	})
}

func SubscriptionsCreate(res http.ResponseWriter, req *http.Request) {
	token := getToken(req)
	url, err := parseSubscriptionsRequest(req.Body)
	PanicIf(err)
	subscription := model.SubscribeUser(url, token)

	respondWith(map[string]interface{}{"Subscription": subscription}, 201, res)
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
	subscription := SubscriptionParams{}

	decoder := json.NewDecoder(body)
	err := decoder.Decode(&subscription)

	return subscription.Url, err
}

func getToken(req *http.Request) string {
	params := req.URL.Query()
	return params["token"][0]
}
