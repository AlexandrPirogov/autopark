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

type Enterprises struct {
	Enterprises []client.Enterprise
}

func EnterprisesPage(w http.ResponseWriter, r *http.Request) {

	c := rest.New("")
	enterprises, _ := c.ListEnterprises()

	data := Enterprises{
		Enterprises: enterprises,
	}
	t, err := template.New("index").ParseFiles("public/templates/enterprises/index.html")
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	err = t.Execute(w, data)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func EnterpriseByTitlePage(w http.ResponseWriter, r *http.Request) {
	c := rest.New("")
	title := chi.URLParam(r, "title")
	res, _ := c.EnterpriseByTitle(title)

	t, err := template.New("index").ParseFiles("public/templates/enterprises/enterprise.html")
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	err = t.Execute(w, res)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func EnterprisesRegisterPage(w http.ResponseWriter, r *http.Request) {
	t, err := template.New("index").ParseFiles("public/templates/enterprises/register/index.html")
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

func EnterpriseRegisterAPI(w http.ResponseWriter, r *http.Request) {

	c := rest.New("")

	var e client.Enterprise
	body, readErr := io.ReadAll(r.Body)
	if readErr != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	unmarshalErr := json.Unmarshal(body, &e)
	if unmarshalErr != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	registerErr := c.RegisterEnterprise(e)
	if registerErr != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)

}
