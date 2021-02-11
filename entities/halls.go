package entities

import (
	"database/sql"
	"time"
)

type Hall struct {
	Id        int           `json:"id" db:"id"`
	Name      string        `json:"name" db:"name"`
	Seats     int           `json:"seats" db:"seats"`
	TheatreId sql.NullInt64 `json:"theatre_id" db:"theatre_id"`
	CreatedAt time.Time     `json:"created_at" db:"created_at"`
	UpdatedAt time.Time     `json:"updated_at" db:"updated_at"`
}
