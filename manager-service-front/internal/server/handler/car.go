package handler

import (
	"html/template"
	"io"
	"log"
	"manager-service-front/internal/client/rest"
	"net/http"
)

func RegisterCarAPI(w http.ResponseWriter, r *http.Request) {
	token, _ := retrieveRefreshToken(r)
	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	c := rest.New(token)
	restErr := c.RegisterCar(body)
	if restErr != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	t, err := template.New("index").ParseFiles("public/templates/index.html")
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	err = t.Execute(w, nil)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}
