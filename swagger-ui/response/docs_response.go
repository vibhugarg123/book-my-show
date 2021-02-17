package response

import (
	"database/sql"
	"time"
)

// swagger:model
type AddBookingResponse struct {
	Id         int       `json:"id"`
	UserId     int       `json:"user_id"`
	ShowId     int       `json:"show_id"`
	Seats      int       `json:"seats"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	TotalPrice float64   `json:"total_price"`
	MovieId    int       `json:"movie_id"`
}

// swagger:model
type AddHallResponse struct {
	Id        int           `json:"id"`
	Name      string        `json:"name"`
	Seats     int           `json:"seats"`
	TheatreId sql.NullInt64 `json:"theatre_id"`
	CreatedAt time.Time     `json:"created_at"`
	UpdatedAt time.Time     `json:"updated_at"`
}

// swagger:model
type AddMovieResponse struct {
	Id           int       `json:"id"`
	Name         string    `json:"name"`
	DirectorName string    `json:"director_name"`
	ReleaseDate  time.Time `json:"release_date"`
	IsActive     bool      `json:"is_active"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

// swagger:model
type AddRegionResponse struct {
	Id         int           `json:"id"`
	Name       string        `json:"name"`
	RegionType int           `json:"region_type"`
	ParentId   sql.NullInt64 `json:"parent_id"`
	CreatedAt  time.Time     `json:"created_at"`
	UpdatedAt  time.Time     `json:"updated_at"`
}

// swagger:model
type AddShowResponse struct {
	Id             int            `json:"id"`
	MovieId        sql.NullInt64  `json:"movie_id"`
	HallId         sql.NullInt64  `json:"hall_id"`
	ShowDate       time.Time      `json:"show_date"`
	TimingId       timingResponse `json:"timing_id"`
	SeatPrice      float64        `json:"seat_price"`
	AvailableSeats int            `json:"available_seats"`
	CreatedAt      time.Time      `json:"created_at"`
	UpdatedAt      time.Time      `json:"updated_at"`
}

// swagger:model
type AddTheatreResponse struct {
	Id        int           `json:"id"`
	Name      string        `json:"name"`
	Address   string        `json:"address"`
	RegionId  sql.NullInt64 `json:"region_id"`
	CreatedAt time.Time     `json:"created_at"`
	UpdatedAt time.Time     `json:"updated_at"`
}

// swagger:model
type timingResponse struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// swagger:model
type AddUserResponse struct {
	Id        int       `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	EmailId   string    `json:"email_id"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// swagger:model
type LoginResponse struct {
	LoginStatus string `json:"login_status"`
}
