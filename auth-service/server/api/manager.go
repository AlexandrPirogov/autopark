package api

import (
	"auth-service/internal/auth"
	"auth-service/internal/storage/db"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func LoginManager(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var c Creds
	marshalErr := json.Unmarshal(body, &c)
	if marshalErr != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	RT, err := auth.VerifyManagerCredentionals(c.Login, c.Pwd, db.GetCurrentCredsStorerInstance())
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = auth.StoreRefreshToken(RT, db.GetCurrentJWTStorerInstance())
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	cookie := setRefreshCookieToken(RT)
	http.SetCookie(w, cookie)
	w.WriteHeader(http.StatusOK)
}

func RegisterManager(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var c Creds
	marshalErr := json.Unmarshal(body, &c)
	log.Println(c)
	if marshalErr != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	registerErr := auth.RegisterManager(c.Login, c.Pwd, db.GetCurrentCredsStorerInstance())
	if registerErr != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func LogoutManager(w http.ResponseWriter, r *http.Request) {
	cookie := unsetRefreshCookieToken()
	http.SetCookie(w, cookie)
	w.WriteHeader(http.StatusOK)
}
