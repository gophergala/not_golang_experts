package model

import "time"

type Page struct {
	Url						string
	LastCheckedAt time.Time
	CreatedAt			time.Time
	UpdatedAt			time.Time
}
