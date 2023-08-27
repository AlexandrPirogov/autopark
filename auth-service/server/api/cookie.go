package api

import (
	"auth-service/internal/auth/refresh"
	"net/http"
)

func setRefreshCookieToken(val string) *http.Cookie {
	return &http.Cookie{
		Name:  refresh.RerfeshTokenCookieField,
		Value: val,
		Path:  "/",
	}
}

func unsetRefreshCookieToken() *http.Cookie {
	return &http.Cookie{
		Name:   refresh.RerfeshTokenCookieField,
		Value:  "",
		MaxAge: -1,
	}
}
