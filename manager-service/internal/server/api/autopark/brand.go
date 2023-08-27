package autopark

import "net/http"

func BrandList(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func BrandRegister(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
}
