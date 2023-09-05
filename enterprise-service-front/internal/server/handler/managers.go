package handler

import (
	"encoding/json"
	"enterprise-front/internal/client"
	"enterprise-front/internal/client/rest"
	"html/template"
	"io"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type data struct {
	Title string
}

func Managers(w http.ResponseWriter, r *http.Request) {
	t, err := template.New("index").ParseFiles("public/templates/enterprises/managers/index.html")
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

func ManagersRegister(w http.ResponseWriter, r *http.Request) {
	title := chi.URLParam(r, "title")

	t, err := template.New("index").ParseFiles("public/templates/enterprises/managers/register/index.html")
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	err = t.Execute(w, data{Title: title})
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func RegisterManagerAPI(w http.ResponseWriter, r *http.Request) {
	c := rest.New("")
	var m client.Manager
	body, readErr := io.ReadAll(r.Body)
	if readErr != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	unmarshalErr := json.Unmarshal(body, &m)
	if unmarshalErr != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	_, registerErr := c.RegisterManager(m)
	if registerErr != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)

}
