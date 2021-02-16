package entities

import "time"

// swagger:model
// User defines the structure for a user in the application
type User struct {
	Id        int       `json:"-" db:"id"`
	FirstName string    `json:"first_name" db:"first_name"`
	LastName  string    `json:"last_name" db:"last_name"`
	EmailId   string    `json:"email_id" db:"email_id"`
	Password  string    `json:"password" db:"password"`
	CreatedAt time.Time `json:"-" db:"created_at"`
	UpdatedAt time.Time `json:"-" db:"updated_at"`
}
