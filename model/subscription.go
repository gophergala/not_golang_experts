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
	page := Page{}
	user := FindUserByAuthToken(token)
	DB.FirstOrCreate(&page, &Page{Url: url})

	subscription := Subscription{
		PageId: page.Id,
		UserId: user.Id,
	}
	DB.Create(&subscription)
	return subscription
}
