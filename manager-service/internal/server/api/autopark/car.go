package autopark

import "net/http"

func CarList(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func CarRegister(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
}

func CarDelete(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNoContent)
}
