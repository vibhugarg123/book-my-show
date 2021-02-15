package entities

import (
	"time"
)

type Booking struct {
	Id         int       `json:"id" db:"id"`
	UserId     int       `json:"user_id" db:"user_id"`
	ShowId     int       `json:"show_id" db:"show_id"`
	Seats      int       `json:"seats" db:"seats"`
	CreatedAt  time.Time `json:"created_at" db:"created_at"`
	UpdatedAt  time.Time `json:"updated_at" db:"updated_at"`
	TotalPrice float64   `json:"total_price"`
	MovieId    int       `json:"movie_id"`
}
