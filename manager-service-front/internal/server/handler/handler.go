package handler

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"log"
	"manager-service-front/internal/client"
	"manager-service-front/internal/client/rest"
	"net/http"
	"strings"
)

type data struct {
	Cars []client.Car
}

func Index(w http.ResponseWriter, r *http.Request) {
	token, _ := retrieveRefreshToken(r)

	c := rest.New(token)
	res, _ := c.ListCars()

	data := data{res}
	t, err := template.New("index").ParseFiles("public/templates/index.html")
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

func Login(w http.ResponseWriter, r *http.Request) {
	t, err := template.New("login").ParseFiles("public/templates/login.html")
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

func LoginAPI(w http.ResponseWriter, r *http.Request) {
	log.Printf("before %v", r.Cookies())
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("err while reading body %v", err)
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	var m client.Manager
	unmarshalErr := json.Unmarshal(body, &m)
	if unmarshalErr != nil {
		log.Printf("err while unmarshal body %v", err)
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	c := rest.New("")
	cookie, err := c.Authenticate(m)
	log.Printf("cookie %v", *cookie)
	if err != nil {
		log.Printf("err while setting cookie %v", err)
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	http.SetCookie(w, cookie)
	w.WriteHeader(http.StatusOK)
}

func retrieveRefreshToken(r *http.Request) (string, error) {
	var cookie *http.Cookie
	cookie, err := r.Cookie(client.RerfeshTokenCookieField)

	if cookie == nil || err != nil {
		log.Printf("%v", r.Header[client.RerfeshTokenCookieField])
		return "", fmt.Errorf("error while reading token %v", err)
	}

	tokenVal := cookie.Value[strings.Index(cookie.Value, "=")+1:]

	return tokenVal, err
}
