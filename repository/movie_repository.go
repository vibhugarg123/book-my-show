package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/vibhugarg123/book-my-show/appcontext"
	"github.com/vibhugarg123/book-my-show/entities"
	"time"
)

type MovieRepository interface {
	InsertMovie(entities.Movie) error
	FetchMovieByName(string) ([]entities.Movie, error)
	FetchActiveMovies() ([]entities.Movie, error)
}

type movieRepository struct {
	db *sqlx.DB
}

func (m movieRepository) FetchMovieByName(movieName string) ([]entities.Movie, error) {
	var movies []entities.Movie
	err := m.db.Select(&movies, "SELECT * FROM movies WHERE name=?", movieName)
	return movies, err
}

func (m movieRepository) InsertMovie(movie entities.Movie) error {
	tx := m.db.MustBegin()
	_, err := tx.Exec("INSERT INTO movies (name,director_name,release_date,is_active,created_at,updated_at) VALUES (?,?,?,?,?,?)", movie.Name, movie.DirectorName, movie.ReleaseDate, movie.IsActive, time.Now(), time.Now())
	if err != nil {
		return err
	}
	return tx.Commit()
}

func (m movieRepository) FetchActiveMovies() ([]entities.Movie, error) {
	var movies []entities.Movie
	err := m.db.Select(&movies, "SELECT * FROM movies WHERE is_active=TRUE")
	return movies, err
}

func NewMovieRepository() MovieRepository {
	return &movieRepository{db: appcontext.MySqlConnection()}
}
