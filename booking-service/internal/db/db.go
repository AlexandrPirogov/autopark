package db

import (
	"booking-service/internal/db/postgres"
	"booking-service/internal/entity"
)

type BookingStorer interface {
	// Approved executes QueryBookingApproved
	//
	// Pre-cond: given booking with set ID to approve
	//
	// Post-cond: query was executed. If successfull, returns nil. Otherwise returns error

	Approved(entity.Booking) error
	// Canceled executes QueryBookingCanceled
	//
	// Pre-cond: given booking with set ID to cancel
	//
	// Post-cond: query was executed. If successfull, returns nil. Otherwise returns error
	Canceled(entity.Booking) error

	// CarChoosed executes QueryBookingCarChoosed
	//
	// Pre-cond: given booking with set ID to ChooseCAr and CarID
	//
	// Post-cond: query was executed. If successfull, returns nil. Otherwise returns error

	CarChoosed(entity.Booking) error
	// Created executes QueryBookingCreate
	//
	// Pre-cond: given booking with set userID
	//
	// Post-cond: query was executed. If successfull, returns id of newly created record.
	// Otherwise returns -1 and error

	Created(entity.Booking) (int, error)
	// Finish executes QueryBookingFinish
	//
	// Pre-cond: given booking with set userID and set ID
	//
	// Post-cond: query was executed. If successfull returns nil.
	// Otherwise returns error
	Finish(entity.Booking) error
}

// GetConnInstance returns
func GetConnInstance() BookingStorer {
	return postgres.GetInstance()
}
