package postgres

import (
	"booking-service/internal/entity"
	"booking-service/internal/kernel/bookingfsm"
	"context"
	"log"
	"time"
)

const QueryBookingApproved = `update bookings set status = 'approved', finished_at = now() where id = $1`

const QueryBookingCanceled = `update bookings set status = 'canceled', finished_at = now() where id = $1`

const QueryBookingCarChoosed = `update bookings set status = 'car_choosed', car_id = $1 where id = $2`

const QueryBookingCreate = `insert into bookings values(default, $1, -1, $2, $3, $4) returning id`

func (pg *pgconn) Approved(b entity.Booking) error {
	_, err := pg.conn.Exec(context.Background(), QueryBookingApproved, b.ID)
	return err
}

func (pg *pgconn) Canceled(b entity.Booking) error {
	_, err := pg.conn.Exec(context.Background(), QueryBookingCanceled, b.ID)
	return err
}

func (pg *pgconn) CarChoosed(b entity.Booking) error {
	log.Println("choosing car booking writing into db ", b)
	_, err := pg.conn.Exec(context.Background(), QueryBookingCarChoosed, b.CarID, b.ID)
	log.Println("choosing car booking writing err ", err)
	return err
}

func (pg *pgconn) Created(b entity.Booking) (int, error) {
	var id int
	err := pg.conn.QueryRow(context.Background(), QueryBookingCreate, b.UserID, bookingfsm.StateCreated, time.Now(), time.Now()).Scan(&id)
	if err != nil {
		log.Println(err)
	}
	return id, err
}
