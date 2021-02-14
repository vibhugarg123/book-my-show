package entities

import (
	"time"
)

type Timing struct {
	Id        int       `json:"id" db:"id"`
	Name      string    `json:"name" db:"name"`
	StartTime time.Time `json:"start_time" db:"start_time"`
	EndTime   time.Time `json:"end_time" db:"end_time"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}
