package model

import (
	"time"
)

type Subscription struct {
	Id        int64
	UserId    int64
	PageId    int64
	CreatedAt time.Time
	UpdatedAt time.Time
}

func SubscribeUser(url string, token string) Subscription {
	page := FindOrCreatePageByUrl(url)
	user := FindUserByAuthToken(token)
	subscription := Subscription{}
	DB.Where(Subscription{UserId: user.Id, PageId: page.Id}).FirstOrCreate(&subscription)
	return subscription
}

func GetSubscriptionsForUser(token string, success func(subscriptions []Subscription), not_success func(message string)) {
	user := FindUserByAuthToken(token)
	if user.Id != 0 {
		subscriptions := []Subscription{}
		DB.Where(Subscription{UserId: user.Id}).Find(&subscriptions)
		success(subscriptions)
	} else {
		not_success("Invalid session token")
	}
}
