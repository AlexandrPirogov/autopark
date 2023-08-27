package api

import (
	"autopark-service/internal/autopark/car"
	"autopark-service/internal/db"
	kernel "autopark-service/internal/kernel/brand"
	"autopark-service/internal/std"
	"encoding/json"
	"io"
	"net/http"
)

// RegisterBrand receives json body, unmarshal it and trying to store in DB
func RegisterBrand(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var b car.Brand
	marshalErr := json.Unmarshal(body, &b)
	if marshalErr != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	createErr := kernel.CreateBrand(b, db.GetConnInstance())
	if createErr != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

// RegisterBrand receives json body, unmarshal it and reads
// all brands by given pattern
func ReadBrands(w http.ResponseWriter, r *http.Request) {
	var b car.Brand
	l, readErr := kernel.ReadBrands(b, db.GetConnInstance())
	if readErr != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	slice := std.AsSlice(l.Iterator())

	responseBody, marshalErr := json.Marshal(slice)
	if marshalErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(responseBody)
}
