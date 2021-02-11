package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/vibhugarg123/book-my-show/appcontext"
	"github.com/vibhugarg123/book-my-show/entities"
	"time"
)

type HallRepository interface {
	InsertHall(entities.Hall) error
	FetchHallByName(string) ([]entities.Hall, error)
	FetchHallByTheatreId(int) ([]entities.Hall, error)
	FetchHallByNameAndTheatreId(entities.Hall) ([]entities.Hall, error)
}

type hallRepository struct {
	db *sqlx.DB
}

func (h hallRepository) FetchHallByNameAndTheatreId(hall entities.Hall) ([]entities.Hall, error) {
	var halls []entities.Hall
	err := h.db.Select(&halls, "SELECT * FROM halls WHERE name=? AND theatre_id=?", hall.Name, hall.TheatreId)
	return halls, err
}

func (h hallRepository) InsertHall(hall entities.Hall) error {
	tx := h.db.MustBegin()
	_, err := tx.Exec("INSERT INTO halls (name,seats,theatre_id,created_at,updated_at) VALUES (?,?,?,?,?)", hall.Name, hall.Seats, hall.TheatreId, time.Now(), time.Now())
	if err != nil {
		return err
	}
	return tx.Commit()
}

func (h hallRepository) FetchHallByName(hallName string) ([]entities.Hall, error) {
	var halls []entities.Hall
	err := h.db.Select(&halls, "SELECT * FROM halls WHERE name=?", hallName)
	return halls, err
}

func (h hallRepository) FetchHallByTheatreId(theatreId int) ([]entities.Hall, error) {
	var halls []entities.Hall
	err := h.db.Select(&halls, "SELECT * FROM halls WHERE theatre_id=?", theatreId)
	return halls, err
}

func NewHallRepository() HallRepository {
	return &hallRepository{db: appcontext.MySqlConnection()}
}
