package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/vibhugarg123/book-my-show/appcontext"
	"github.com/vibhugarg123/book-my-show/entities"
	"time"
)

type UserRepository interface {
	InsertUser(entities.User) error
	FetchUserByEmailId(string) ([]entities.User, error)
	FetchUserByEmailIdAndPassword(string, string) ([]entities.User, error)
}

type userRepository struct {
	db *sqlx.DB
}

func NewUserRepository() UserRepository {
	return &userRepository{db: appcontext.MySqlConnection()}
}

func (u userRepository) InsertUser(user entities.User) error {
	tx := u.db.MustBegin()
	_, err := tx.Exec("INSERT INTO users  (first_name,last_name,email_id,password,created_at,updated_at) VALUES (?,?,?,?,?,?)", user.FirstName, user.LastName, user.EmailId, user.Password, time.Now(), time.Now())
	if err != nil {
		return err
	}
	return tx.Commit()
}

func (u userRepository) FetchUserByEmailId(email_id string) (user []entities.User, err error) {
	var existingUser []entities.User
	err = u.db.Select(&existingUser, "SELECT * FROM users WHERE email_id=?", email_id)
	return existingUser, err
}

func (u userRepository) FetchUserByEmailIdAndPassword(email_id string, password string) (user []entities.User, err error) {
	var existingUser []entities.User
	err = u.db.Select(&existingUser, "SELECT * FROM users WHERE email_id=? AND password=?", email_id, password)
	return existingUser, err
}
