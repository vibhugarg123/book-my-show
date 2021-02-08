package entities

import (
	"database/sql"
	"time"
)

type Theatre struct {
	Id        int           `json:"id" db:"id"`
	Name      string        `json:"name" db:"name"`
	Address   string        `json:"address" db:"address"`
	RegionId  sql.NullInt64 `json:"region_id" db:"region_id"`
	CreatedAt time.Time     `json:"created_at" db:"created_at"`
	UpdatedAt time.Time     `json:"updated_at" db:"updated_at"`
}
