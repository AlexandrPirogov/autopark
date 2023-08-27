package api

import (
	"encoding/json"
	"enterprise-service/internal/db"
	"enterprise-service/internal/enterprise"
	"enterprise-service/internal/kernel"
	"io"
	"log"
	"net/http"
)

// RegisterEnterprise creates new enterprise entity in system
//
// Pre-cond: given unique enterprise entity
//
// Post-cond: new enterprise entity for created for system
func RegisterEnterprise(w http.ResponseWriter, r *http.Request) {
	var e enterprise.Enterprise
	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	marhsalErr := json.Unmarshal(body, &e)
	if marhsalErr != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	storeErr := kernel.Store(e, db.GetConnInstance())
	if storeErr != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

// Delete removes given enterprise entity from system
//
// Pre-cond: given existing enterprise to remove
//
// Post-cond: enterprise was removed from system
func DeleteEnterprise(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNoContent)
}

// Read read all existing enterprises in system by given pattern
//
// Pre-cond: given pattern for enterprises
//
// Post-cond: returns list of matched enterprises
func ReadEnerprises(w http.ResponseWriter, r *http.Request) {
	var e enterprise.Enterprise
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	marhsalErr := json.Unmarshal(body, &e)
	if marhsalErr != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	list, storeErr := kernel.Read(e, db.GetConnInstance())
	if storeErr != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	responseBody := []byte{}
	for {
		item, ok := list.PopFront()
		if !ok {
			break
		}

		marshaled, err := json.Marshal(item)
		if err != nil {
			continue
		}

		responseBody = append(responseBody, marshaled...)
	}

	w.WriteHeader(http.StatusOK)
	w.Write(responseBody)
}

// Update updates existing enterprise entity in system
//
// Pre-cond: given enterprise pattern to update and new params for them
//
// Post-cond: all enterprises that matches the given pattern was update
// with given new params
func UpdateEnterprises(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
