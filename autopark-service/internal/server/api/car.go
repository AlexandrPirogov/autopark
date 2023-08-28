// api package hold http functions for service API
package api

import (
	"autopark-service/internal/autopark/car"
	"autopark-service/internal/db"
	kernel "autopark-service/internal/kernel/car"
	"autopark-service/internal/std"
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

// CreateCar receives json body, unmarshal it and trying to store in Storage
func CreateCar(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var b car.Car
	marshalErr := json.Unmarshal(body, &b)
	if marshalErr != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	createErr := kernel.CreateCar(b, db.GetConnInstance())
	if createErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

// ReadCars receives json body as pattern, unmarshal it
// and reads all cars from db that match the pattern
func ReadCars(w http.ResponseWriter, r *http.Request) {
	var c car.Car

	l, err := kernel.ReadCars(c, db.GetConnInstance())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
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

// DeleteCar receives json body as pattern, unmarshal it
// and deletes all cars from db that match the pattern
func DeleteCar(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var b car.Car
	marshalErr := json.Unmarshal(body, &b)
	if marshalErr != nil {
		log.Println(marshalErr)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	createErr := kernel.DeleteCar(b, db.GetConnInstance())
	if createErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func SetCar(w http.ResponseWriter, r *http.Request) {
	uid := chi.URLParam(r, "uid")
	c := car.Car{UID: uid}
	err := kernel.SetCar(c, db.GetConnInstance())
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func UnsetCar(w http.ResponseWriter, r *http.Request) {
	uid := chi.URLParam(r, "uid")
	c := car.Car{UID: uid}
	err := kernel.UnsetCar(c, db.GetConnInstance())
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
}
