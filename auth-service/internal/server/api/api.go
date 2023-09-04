package api

import (
	"auth-service/internal/auth"
	"auth-service/internal/storage/db"
	"encoding/json"
	"io"
	"net/http"
)

type Creds struct {
	Login string `json:"login"`
	Pwd   string `json:"pwd"`
}

// LoginAdmin authenticate admin. If creds are correct then generate
// new Refresh Token and put it in the request header
func LoginAdmin(w http.ResponseWriter, r *http.Request) {
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

	RT, err := auth.VerifyAdminCredentionals(c.Login, c.Pwd, db.GetCurrentCredsStorerInstance())
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

// LogoutAdmin logout admin and unsets it RefreshToken
func LogoutAdmin(w http.ResponseWriter, r *http.Request) {
	cookie := unsetRefreshCookieToken()
	http.SetCookie(w, cookie)
	w.WriteHeader(http.StatusOK)
}

// VerifyJWT just a stub for api-gateway auth
func VerifyJWT(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
