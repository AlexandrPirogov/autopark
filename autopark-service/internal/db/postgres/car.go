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

// Delete cars entities from db
// that match given pattern
const QueryDeleteCar = `
delete from cars where brand_id =
(select id from brands where brand = $1)
and uid = $2 and type = $3`

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
	return nil
}
