package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/vibhugarg123/book-my-show/appcontext"
	"github.com/vibhugarg123/book-my-show/entities"
	"time"
)

type TheatreRepository interface {
	InsertTheatre(entities.Theatre) error
	FetchTheatreByName(string) ([]entities.Theatre, error)
	FetchTheatreByRegionId(int) ([]entities.Theatre, error)
	FetchTheatreByNameRegionIdAndAddress(entities.Theatre) ([]entities.Theatre, error)
}

type theatreRepository struct {
	db *sqlx.DB
}

func (t theatreRepository) FetchTheatreByNameRegionIdAndAddress(theatre entities.Theatre) ([]entities.Theatre, error) {
	var theatres []entities.Theatre
	err := t.db.Select(&theatres, "SELECT * FROM theatres WHERE name=? AND region_id=? AND address=?", theatre.Name, theatre.RegionId.Int64, theatre.Address)
	return theatres, err
}

func (t theatreRepository) InsertTheatre(theatre entities.Theatre) error {
	tx := t.db.MustBegin()
	_, err := tx.Exec("INSERT INTO theatres (name,address,region_id,created_at,updated_at) VALUES (?,?,?,?,?)", theatre.Name, theatre.Address, theatre.RegionId, time.Now(), time.Now())
	if err != nil {
		return err
	}
	return tx.Commit()
}

func (t theatreRepository) FetchTheatreByName(theatreName string) ([]entities.Theatre, error) {
	var theatres []entities.Theatre
	err := t.db.Select(&theatres, "SELECT * FROM theatres WHERE name=?", theatreName)
	return theatres, err
}

func (t theatreRepository) FetchTheatreByRegionId(regionId int) ([]entities.Theatre, error) {
	var theatres []entities.Theatre
	err := t.db.Select(&theatres, "SELECT * FROM theatres WHERE region_id=?", regionId)
	return theatres, err
}

func NewTheatreRepository() TheatreRepository {
	return &theatreRepository{db: appcontext.MySqlConnection()}
}
