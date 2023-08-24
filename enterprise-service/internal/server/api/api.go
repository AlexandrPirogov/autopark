package api

import "net/http"

// RegisterEnterprise creates new enterprise entity in system
//
// Pre-cond: given unique enterprise entity
//
// Post-cond: new enterprise entity for created for system
func RegisterEnterprise(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

// Delete removes given enterprise entity from system
//
// Pre-cond: given existing enterprise to remove
//
// Post-cond: enterprise was removed from system
func Delete(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

// Read read all existing enterprises in system by given pattern
//
// Pre-cond: given pattern for enterprises
//
// Post-cond: returns list of matched enterprises
func Read(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func Update(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
