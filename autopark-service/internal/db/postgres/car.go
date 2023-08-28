// car.go holds the functionality to CRUD car entity
package postgres

import (
	"autopark-service/internal/autopark/car"
	"autopark-service/internal/std"
	"autopark-service/internal/std/list"
	"context"
	"log"
)

// Reads ALL cars from db
const QueryReadCars = `
select brand, type, uid from cars c
join brands b on b.id = c.brand_id`

// Create new record and stores it in database
// brand is required
const QueryStoreCar = `
insert into cars values(
	default,
	(select id from brands where brand = $1),
	$2, 
	$3)`

const QueryStoreCarForBookings = `
insert into cars_for_booking values(
	default,
	(select id from cars where uid = $1),
	'set', 
	0)`

// Delete cars entities from db
// that match given pattern
const QueryDeleteCar = `
delete from cars where brand_id =
(select id from brands where brand = $1)
and uid = $2 and type = $3`

const QuerySetCar = `update cars_for_booking set status = 'set' where id = (select id from cars where uid = $1)`

const QueryUnsetCar = `update cars_for_booking set status = 'unset' where id = (select id from cars where uid = $1)`

const QueryReadSetCars = `
select c.uid, b.brand, type, status from cars_for_booking cfb
join cars c on c.id = cfb.car_id
join brands b on c.brand_id = b.id
where cfb.status = 'set'
`

// DeleteCars executing QueryDeleteCar query
//
// Pre-cond: given car-pattern
//
// Post-cond: query was executed.
// If success then returns nil and cars that match given pattern was removed
// Otherwise returns error
func (pg *pgconn) DeleteCars(pattern car.Car) error {
	_, err := pg.conn.Exec(context.Background(), QueryDeleteCar, pattern.Brand, pattern.UID, pattern.Type)
	if err != nil {
		return err
	}
	return nil
}

// error executing QueryReadCars query
//
// Pre-cond: given car-pattern
//
// Post-cond: query was executed.
// If success then returns list of read cars and nil
// Otherwise returns nil and error
func (pg *pgconn) ReadCars(c car.Car) (std.Linked[car.Car], error) {
	rows, err := pg.conn.Query(context.Background(), QueryReadCars)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	res := list.New[car.Car]()
	for rows.Next() {
		var c car.Car
		scanErr := rows.Scan(&c.Brand, &c.Type, &c.UID)
		if scanErr != nil {
			log.Println(err)
			continue
		}
		res.PushBack(c)
	}

	return res, nil
}

func (pg *pgconn) ReadSetCars() (std.Linked[car.Car], error) {
	rows, err := pg.conn.Query(context.Background(), QueryReadSetCars)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	res := list.New[car.Car]()
	for rows.Next() {
		var c car.Car
		scanErr := rows.Scan(&c.UID, &c.Brand, &c.Type, &c.Status)
		if scanErr != nil {
			log.Println(scanErr)
			continue
		}
		res.PushBack(c)
	}

	return res, nil
}

// StoreCar executing QueryStoreCar query
//
// Pre-cond: given car instance
//
// Post-cond: query was executed.
// If success then returns nil and given car was stored in DB
// Otherwise returns error
func (pg *pgconn) StoreCar(c car.Car) error {
	_, err := pg.conn.Exec(context.Background(), QueryStoreCar, c.Brand, c.UID, c.Type)
	if err != nil {
		return err
	}

	_, err = pg.conn.Exec(context.Background(), QueryStoreCarForBookings, c.UID)
	if err != nil {
		return err
	}
	return nil
}

func (pg *pgconn) SetCar(c car.Car) error {
	_, err := pg.conn.Exec(context.Background(), QuerySetCar, c.UID)
	if err != nil {
		return err
	}
	return nil
}

func (pg *pgconn) UnsetCar(c car.Car) error {
	_, err := pg.conn.Exec(context.Background(), QueryUnsetCar, c.UID)
	if err != nil {
		return err
	}
	return nil
}
