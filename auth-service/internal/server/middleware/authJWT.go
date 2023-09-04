package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/rs/zerolog/log"

	auth "auth-service/internal/auth/jwt"

	"github.com/golang-jwt/jwt/v5"
)

// AuthJWT verify that request contains in cookie "refresh-token" refresh token
func AuthJWT(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if err := verifyRefreshToken(r); err != nil {
			log.Warn().Msgf("verify err %v", err)
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func verifyRefreshToken(r *http.Request) error {
	var cookie *http.Cookie
	cookie, err := r.Cookie(auth.RerfeshTokenCookieField)
	if cookie == nil || err != nil {
		log.Warn().Msgf("%v", r.Header[auth.RerfeshTokenCookieField])
		return fmt.Errorf("error while reading token %v", err)
	}
	tokenVal := cookie.Value[strings.Index(cookie.Value, "=")+1:]
	token, err := jwt.Parse(tokenVal, func(t *jwt.Token) (interface{}, error) {
		_, ok := t.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			log.Warn().Msg("not ok")
		}
		return []byte(auth.JWTSecret), nil

	})

	if err != nil {
		log.Warn().Msgf("%v", err)
		return err
	}

	if token.Valid {
		return nil
	}

	return fmt.Errorf("refresh token is invalid")
}
