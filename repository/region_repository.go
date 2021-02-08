package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/vibhugarg123/book-my-show/appcontext"
	"github.com/vibhugarg123/book-my-show/entities"
	"time"
)

type RegionRepository interface {
	InsertRegion(region entities.Region) error
	FetchRegionById(int) ([]entities.Region, error)
}

type regionRepository struct {
	db *sqlx.DB
}

func (r regionRepository) InsertRegion(region entities.Region) error {
	tx := r.db.MustBegin()
	r.db.MustExec("INSERT INTO regions (id,name,region_type,parent_id,created_at,updated_at) VALUES (?,?,?,?,?,?)", region.Id, region.Name, region.RegionType, region.ParentId, time.Now(), time.Now())
	return tx.Commit()
}

func (r regionRepository) FetchRegionById(regionId int) ([]entities.Region, error) {
	var regions []entities.Region
	err := r.db.Select(&regions, "SELECT * FROM regions WHERE id=?", regionId)
	return regions, err
}

func NewRegionRepository() RegionRepository {
	return &regionRepository{db: appcontext.MySqlConnection()}
}
