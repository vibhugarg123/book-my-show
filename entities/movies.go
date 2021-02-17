package entities

import (
	"time"
)

type Movie struct {
	Id           int       `json:"id" db:"id"`
	Name         string    `json:"name" db:"name"`
	DirectorName string    `json:"director_name" db:"director_name"`
	ReleaseDate  time.Time `json:"release_date" db:"release_date"`
	IsActive     bool      `json:"is_active" db:"is_active"`
	CreatedAt    time.Time `json:"created_at" db:"created_at"`
	UpdatedAt    time.Time `json:"updated_at" db:"updated_at"`
}

