package api

import (
	"auth-service/internal/auth/jwt"
	"net/http"
)

func setRefreshCookieToken(val string) *http.Cookie {
	return &http.Cookie{
		Name:  jwt.RerfeshTokenCookieField,
		Value: val,
		Path:  "/",
	}
}

func unsetRefreshCookieToken() *http.Cookie {
	return &http.Cookie{
		Name:   jwt.RerfeshTokenCookieField,
		Value:  "",
		MaxAge: -1,
	}
}

func setAccessCookieToken(val string) *http.Cookie {
	return &http.Cookie{
		Name:  jwt.AccessTokenCookieField,
		Value: val,
		Path:  "/",
	}
}

func unsetAccessCookieToken() *http.Cookie {
	return &http.Cookie{
		Name:   jwt.AccessTokenCookieField,
		Value:  "",
		MaxAge: -1,
	}
}
