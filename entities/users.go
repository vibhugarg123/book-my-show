package entities

import "time"

type User struct {
	Id        int       `json:"id" db:"id"`
	FirstName string    `json:"first_name" db:"first_name"`
	LastName  string    `json:"last_name" db:"last_name"`
	EmailId   string    `json:"email_id" db:"email_id"`
	Password  string    `json:"password" db:"password"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

var UsersSchema = `
CREATE TABLE users ( 
		id integer,
		first_name text,
  		last_name text,
    	email_id text,
    	password text,
   	    created_at DATETIME,
        updated_at DATETIME
   );
`
