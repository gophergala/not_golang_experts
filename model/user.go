package model

import "time"

type User struct {
	Id                int64
	Email             string
	EncryptedPassword []byte
	Token             string
	CreatedAt         time.Time
	UpdatedAt         time.Time
}
