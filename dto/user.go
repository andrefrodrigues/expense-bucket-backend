package dto

import "time"

type User struct {
	Name      string
	Email     string
	CreatedAt time.Time
}
