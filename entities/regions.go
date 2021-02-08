package entities

import (
	"database/sql"
	"time"
)

type Region struct {
	Id         int           `json:"id" db:"id"`
	Name       string        `json:"name" db:"name"`
	RegionType int           `json:"region_type" db:"region_type"`
	ParentId   sql.NullInt64 `json:"parent_id" db:"parent_id"`
	CreatedAt  time.Time     `json:"created_at" db:"created_at"`
	UpdatedAt  time.Time     `json:"updated_at" db:"updated_at"`
}
