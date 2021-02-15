package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/vibhugarg123/book-my-show/appcontext"
	"github.com/vibhugarg123/book-my-show/entities"
	"time"
)

type BookingRepository interface {
	InsertBooking(*entities.Booking) error
	FetchBookingByUserId(int) ([]entities.Booking, error)
}

type bookingRepository struct {
	db *sqlx.DB
}

func (b bookingRepository) FetchBookingByUserId(userId int) ([]entities.Booking, error) {
	var bookingList []entities.Booking
	err := b.db.Select(&bookingList, "SELECT * FROM bookings WHERE user_id=?", userId)
	return bookingList, err
}

func (b bookingRepository) InsertBooking(booking *entities.Booking) error {
	tx := b.db.MustBegin()
	result, err := tx.Exec("INSERT INTO bookings (user_id,show_id,seats,created_at,updated_at) VALUES (?,?,?,?,?)", booking.UserId, booking.ShowId, booking.Seats, time.Now(), time.Now())
	if err != nil {
		return err
	}
	bookingId, err := result.LastInsertId()
	if err != nil {
		return err
	}
	booking.Id = int(bookingId)
	return tx.Commit()
}

func NewBookingRepository() BookingRepository {
	return &bookingRepository{db: appcontext.MySqlConnection()}
}
