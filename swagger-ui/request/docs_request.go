package request

import (
	"database/sql"
	"time"
)

// swagger:model
type NewBooking struct {
	UserId     int     `json:"user_id"`
	ShowId     int     `json:"show_id"`
	Seats      int     `json:"seats"`
	TotalPrice float64 `json:"total_price"`
	MovieId    int     `json:"movie_id"`
}

// swagger:model
type AddHall struct {
	Name      string        `json:"name"`
	Seats     int           `json:"seats"`
	TheatreId sql.NullInt64 `json:"theatre_id"`
}

// swagger:model
type AddMovie struct {
	Name         string    `json:"name"`
	DirectorName string    `json:"director_name"`
	ReleaseDate  time.Time `json:"release_date"`
	IsActive     bool      `json:"is_active"`
}

// swagger:model
type AddRegion struct {
	Id         int           `json:"id"`
	Name       string        `json:"name"`
	RegionType int           `json:"region_type"`
	ParentId   sql.NullInt64 `json:"parent_id"`
}

// swagger:model
type AddShow struct {
	MovieId        sql.NullInt64 `json:"movie_id"`
	HallId         sql.NullInt64 `json:"hall_id"`
	ShowDate       time.Time     `json:"show_date"`
	TimingId       timingRequest `json:"timing_id"`
	SeatPrice      float64       `json:"seat_price"`
	AvailableSeats int           `json:"available_seats"`
}

// swagger:model
type AddTheatre struct {
	Name     string        `json:"name"`
	Address  string        `json:"address"`
	RegionId sql.NullInt64 `json:"region_id"`
}

// swagger:model
type timingRequest struct {
	Name      string    `json:"name"`
	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time"`
}

// swagger:model
type AddUser struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	EmailId   string `json:"email_id"`
	Password  string `json:"password"`
}

// swagger:model
type AddBooking struct {
	UserId int `json:"user_id"`
	ShowId int `json:"show_id"`
	Seats  int `json:"seats"`
}

// swagger:model
type LoginRequest struct {
	EmailId  string `json:"email_id"`
	Password string `json:"password"`
}
