package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/vibhugarg123/book-my-show/appcontext"
	"github.com/vibhugarg123/book-my-show/entities"
	"time"
)

type ShowRepository interface {
	InsertShow(*entities.Show) error
	FetchShowByMovieIdHallIdShowDate(entities.Show) ([]entities.Show, error)
	FetchShowByShowId(int) ([]entities.Show, error)
	UpdateSeatsByShowId(int, int) error
}

type showRepository struct {
	db *sqlx.DB
}

func (s showRepository) InsertShow(show *entities.Show) error {
	tx := s.db.MustBegin()
	timing, err := tx.Exec("INSERT INTO timings (name, start_time,end_time,created_at,updated_at) VALUES (?,?,?,?,?)", show.TimingId.Name, show.TimingId.StartTime, show.TimingId.EndTime, time.Now(), time.Now())
	if err != nil {
		return err
	}
	timingId, err := timing.LastInsertId()
	if err != nil {
		return err
	}
	showInserted, err := tx.Exec("INSERT INTO shows (movie_id,hall_id,show_date,timing_id,seat_price,available_seats,created_at,updated_at) VALUES (?,?,?,?,?,?,?,?)", show.MovieId.Int64, show.HallId.Int64, show.ShowDate, timingId, show.SeatPrice, show.AvailableSeats, time.Now(), time.Now())
	if err != nil {
		return err
	}
	showInsertedId, err := showInserted.LastInsertId()
	if err != nil {
		return err
	}
	show.TimingId.Id = int(timingId)
	show.Id = int(showInsertedId)
	return tx.Commit()
}

//Assuming show_date will give start date time of show
func (s showRepository) FetchShowByMovieIdHallIdShowDate(show entities.Show) ([]entities.Show, error) {
	var showReturned []entities.Show
	query := `SELECT
		s.id, 
		s.movie_id, 
		s.hall_id, 
		s.show_date,
			t.id "timing_id.id",
			t.name "timing_id.name",
			t.start_time "timing_id.start_time",
			t.end_time "timing_id.end_time",
            t.created_at "timing_id.created_at",
            t.updated_at "timing_id.updated_at",
		s.seat_price,
		s.available_seats,
		s.created_at,
		s.updated_at
	FROM shows s
	JOIN timings t ON s.timing_id=t.id WHERE s.movie_id=? AND s.hall_id=? AND s.show_date=?`
	err := s.db.Select(&showReturned, query, show.MovieId.Int64, show.HallId.Int64, show.ShowDate)
	if err != nil {
		return []entities.Show{}, err
	}
	return showReturned, err
}

func (s showRepository) FetchShowByShowId(showId int) ([]entities.Show, error) {
	var showReturned []entities.Show
	query := `SELECT
		s.id, 
		s.movie_id, 
		s.hall_id, 
		s.show_date,
			t.id "timing_id.id",
			t.name "timing_id.name",
			t.start_time "timing_id.start_time",
			t.end_time "timing_id.end_time",
            t.created_at "timing_id.created_at",
            t.updated_at "timing_id.updated_at",
		s.seat_price,
		s.available_seats,
		s.created_at,
		s.updated_at
	FROM shows s
	JOIN timings t ON s.timing_id=t.id WHERE s.id=?`
	err := s.db.Select(&showReturned, query, showId)
	if err != nil {
		return []entities.Show{}, err
	}
	return showReturned, err
}

func (s showRepository) UpdateSeatsByShowId(seats int, showId int) error {
	query := `UPDATE shows SET available_seats=? WHERE id=?`
	tx := s.db.MustBegin()
	_, err := tx.Exec(query, seats, showId)
	if err != nil {
		return err
	}
	return tx.Commit()
}

func NewShowRepository() ShowRepository {
	return &showRepository{db: appcontext.MySqlConnection()}
}