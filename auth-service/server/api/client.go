package api

import (
	"auth-service/internal/auth"
	"auth-service/internal/storage/db"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

// LoginClient authenticate Client. If creds are correct then generate
// new Refresh Token and put it in the request header
func LoginClient(w http.ResponseWriter, r *http.Request) {
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

	client, err := auth.VerifyClientCredentionals(c.Login, c.Pwd, db.GetCurrentCredsStorerInstance())
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = auth.StoreRefreshToken(client.RefreshToken(), db.GetCurrentJWTStorerInstance())
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	cookie := setRefreshCookieToken(client.RefreshToken())
	http.SetCookie(w, cookie)

	responseBody, marshalErr := json.Marshal(client)
	if marshalErr != nil {
		log.Println(marshalErr)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(responseBody)
}

// Register Client registers new Client with given creds
func RegisterClient(w http.ResponseWriter, r *http.Request) {
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

	id, registerErr := auth.RegisterClient(c.Login, c.Pwd, db.GetCurrentCredsStorerInstance())
	if registerErr != nil {
		log.Printf("auth register err %v", registerErr)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf("{\"id\":%d}", id)))
}

func LogoutClient(w http.ResponseWriter, r *http.Request) {
	cookie := unsetRefreshCookieToken()
	http.SetCookie(w, cookie)
	w.WriteHeader(http.StatusOK)
}
