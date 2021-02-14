package entities

import (
	"database/sql"
	"time"
)

type Show struct {
	Id             int           `json:"id" db:"id"`
	MovieId        sql.NullInt64 `json:"movie_id" db:"movie_id"`
	HallId         sql.NullInt64 `json:"hall_id" db:"hall_id"`
	ShowDate       time.Time     `json:"show_date" db:"show_date"`
	TimingId       Timing        `json:"timing_id" db:"timing_id"`
	SeatPrice      float64       `json:"seat_price" db:"seat_price"`
	AvailableSeats int           `json:"available_seats" db:"available_seats"`
	CreatedAt      time.Time     `json:"created_at" db:"created_at"`
	UpdatedAt      time.Time     `json:"updated_at" db:"updated_at"`
}
