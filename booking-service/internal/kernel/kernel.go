package kernel

import (
	"booking-service/internal/db"
	"booking-service/internal/entity"
)

func ChooseCarBooking(b entity.Booking, s db.BookingStorer) error {
	return s.CarChoosed(b)
}

func CreateBooking(b entity.Booking, s db.BookingStorer) (int, error) {
	return s.Created(b)
}

func ApproveBooking(b entity.Booking, s db.BookingStorer) error {
	return s.Approved(b)
}

func CancelBooking(b entity.Booking, s db.BookingStorer) error {
	return s.Canceled(b)
}
