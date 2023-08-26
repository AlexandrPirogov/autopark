package api

import "net/http"

func EnterprisesList(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
