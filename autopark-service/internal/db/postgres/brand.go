// brand.go holds the functionality to CRUD brand entity
package postgres

import (
	"autopark-service/internal/autopark/car"
	"autopark-service/internal/std"
	"autopark-service/internal/std/list"
	"context"

	"github.com/rs/zerolog/log"
)

// QueryReadBrands reads all brands from DB
const QueryReadBrands = `select brand from brands`

// QueryStoreBrand stores given brand in DB
const QueryStoreBrand = `insert into Brands values(default, $1)`

// QueryDeleteBrand delete brands from DB that match given pattern
const QueryDeleteBrand = `delete from Brands where brand = $1`

// ReadBrands executes QueryReadBrands and returns list of'em
//
// Pre-cond: given brand-pattern instance
//
// Post-cond: query was executed.
// If success then returns nil and brands that match given pattern was removed
// Otherwise returns error
func (pg *pgconn) ReadBrands(pattern car.Brand) (std.Linked[car.Brand], error) {
	rows, err := pg.conn.Query(context.Background(), QueryReadBrands)
	if err != nil {
		log.Warn().Msgf("%v", err)
		return nil, err
	}

	res := list.New[car.Brand]()
	for rows.Next() {
		var b car.Brand
		scanErr := rows.Scan(&b.Brand)
		if scanErr != nil {
			log.Warn().Msgf("%v", err)
			continue
		}
		res.PushBack(b)
	}

	return res, nil
}

// StoreBrand executing QueryStoreBrand query
//
// Pre-cond: given brand instance
//
// Post-cond: query was executed.
// If success then returns nil and given brand was stored in DB
// Otherwise returns error
func (pg *pgconn) StoreBrand(b car.Brand) error {
	_, err := pg.conn.Exec(context.Background(), QueryStoreBrand, b.Brand)
	if err != nil {
		return err
	}
	return nil
}
