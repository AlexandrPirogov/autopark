package postgres

import (
	"booking-service/internal/entity"
	"booking-service/internal/kernel/bookingfsm"
	"context"
	"time"

	"github.com/rs/zerolog/log"
)

// QueryBookingApproved set status of booking to approved
const QueryBookingApproved = `update bookings set status = 'approved', finished_at = now() where id = $1`

// QueryBookingCanceled set status of booking to canceled
const QueryBookingCanceled = `update bookings set status = 'canceled', finished_at = now() where id = $1`

// QueryBookingCarChoosed set status of booking to car_choosed
const QueryBookingCarChoosed = `update bookings set status = 'car_choosed', car_id = $1 where id = $2`

// QueryBookingCreate inserts new record for booking
const QueryBookingCreate = `insert into bookings values(default, $1, -1, $2, $3, $4) returning id`

// QueryBookingFinished set status of booking to finished
const QueryBookingFinished = `update bookings set status = 'finished', finished_at = now() where id = $1`

// Approved executes QueryBookingApproved
//
// Pre-cond: given booking with set ID to approve
//
// Post-cond: query was executed. If successfull, returns nil. Otherwise returns error
func (pg *pgconn) Approved(b entity.Booking) error {
	_, err := pg.conn.Exec(context.Background(), QueryBookingApproved, b.ID)
	return err
}

// Canceled executes QueryBookingCanceled
//
// Pre-cond: given booking with set ID to cancel
//
// Post-cond: query was executed. If successfull, returns nil. Otherwise returns error
func (pg *pgconn) Canceled(b entity.Booking) error {
	_, err := pg.conn.Exec(context.Background(), QueryBookingCanceled, b.ID)
	return err
}

// CarChoosed executes QueryBookingCarChoosed
//
// Pre-cond: given booking with set ID to ChooseCAr and CarID
//
// Post-cond: query was executed. If successfull, returns nil. Otherwise returns error
func (pg *pgconn) CarChoosed(b entity.Booking) error {
	_, err := pg.conn.Exec(context.Background(), QueryBookingCarChoosed, b.CarID, b.ID)
	return err
}

// Created executes QueryBookingCreate
//
// Pre-cond: given booking with set userID
//
// Post-cond: query was executed. If successfull, returns id of newly created record.
// Otherwise returns -1 and error
func (pg *pgconn) Created(b entity.Booking) (int, error) {
	var id int
	err := pg.conn.QueryRow(context.Background(), QueryBookingCreate, b.UserID, bookingfsm.StateCreated, time.Now(), time.Now()).Scan(&id)
	if err != nil {
		log.Warn().Msgf("%v", err)
	}
	return id, err
}

// Finish executes QueryBookingFinish
//
// Pre-cond: given booking with set userID and set ID
//
// Post-cond: query was executed. If successfull returns nil.
// Otherwise returns error
func (pg *pgconn) Finish(b entity.Booking) error {
	_, err := pg.conn.Exec(context.Background(), QueryBookingFinished, b.ID)
	return err
}
