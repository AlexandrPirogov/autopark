package handler

import (
	"enterprise-front/internal/client"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
)

func Index(w http.ResponseWriter, r *http.Request) {
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

func APILogin(w http.ResponseWriter, r *http.Request) {

}

func retrieveRefreshToken(r *http.Request) (string, error) {
	var cookie *http.Cookie
	cookie, err := r.Cookie(client.RerfeshTokenCookieField)

	if cookie == nil || err != nil {
		//log.Warn().Msgf("%v", r.Header[client.RerfeshTokenCookieField])
		return "", fmt.Errorf("error while reading token %v", err)
	}

	tokenVal := cookie.Value[strings.Index(cookie.Value, "=")+1:]

	return tokenVal, err
}
