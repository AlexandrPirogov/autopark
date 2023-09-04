package kernel

import (
	"autopark-service/internal/autopark/car"
	"autopark-service/internal/db"
	"autopark-service/internal/std"
)

// CreateBrand stores brand in storage
//
// Pre-cond: given brand entity and entity implementing BrandStorer interafce
//
// Post-cond: executes command to store given brand in given BrandStorer
// returns nil if command executed successfully otherwise error
func CreateBrand(b car.Brand, s db.BrandStorer) error {
	return s.StoreBrand(b)
}

// ReadBrands return list from storage
//
// Pre-cond: given brand entity and entity implementing BrandStorer interafce
//
// Post-cond: executes command to read brands by given brand-pattern from given BrandStorer
// returns nil if command executed successfully otherwise error
func ReadBrands(pattern car.Brand, s db.BrandStorer) (std.Linked[car.Brand], error) {
	return s.ReadBrands(pattern)
}
