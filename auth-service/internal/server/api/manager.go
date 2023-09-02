package api

import (
	"auth-service/internal/auth"
	"auth-service/internal/storage/db"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/rs/zerolog/log"
)

// LoginManager authenticate manager. If creds are correct then generate
// new Refresh Token and put it in the request header
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

// Register manager registers new manager with given creds
func RegisterManager(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("auth-service: io.ReadAll err %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var c Creds
	marshalErr := json.Unmarshal(body, &c)
	if marshalErr != nil {
		log.Printf("auth-service: unmarshal err %v", marshalErr)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	log.Debug().Msgf("registering manager %v", c)
	id, registerErr := auth.RegisterManager(c.Login, c.Pwd, db.GetCurrentCredsStorerInstance())
	if registerErr != nil {
		log.Printf("auth register err %v", registerErr)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	log.Debug().Msgf("registeried manager %v succcessfully", c)
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf("{\"id\":%d}", id)))
}

func LogoutManager(w http.ResponseWriter, r *http.Request) {
	cookie := unsetRefreshCookieToken()
	http.SetCookie(w, cookie)
	w.WriteHeader(http.StatusOK)
}
