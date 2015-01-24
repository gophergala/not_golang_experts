package models

import "time"

type User struct {
	Email     string
	CreatedAt time.Time
	UpdatedAt time.Time
}
