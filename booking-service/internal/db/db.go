package db

import (
	"booking-service/internal/db/postgres"
	"booking-service/internal/entity"
)

type BookingStorer interface {
	Approved(b entity.Booking) error
	Canceled(b entity.Booking) error
	CarChoosed(b entity.Booking) error
	Created(b entity.Booking) (int, error)
}

func GetConnInstance() BookingStorer {
	return postgres.GetInstance()
}
