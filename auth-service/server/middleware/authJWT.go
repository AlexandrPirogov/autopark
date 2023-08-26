package middleware

import (
	"auth-service/internal/auth/refresh"
	"log"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

func AuthJWT(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		var cookie *http.Cookie
		cookie, err := r.Cookie("refresh-token")

		if cookie == nil || err != nil {
			log.Println(r.Header[refresh.RerfeshTokenCookieField])
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		tokenVal := cookie.Value[strings.Index(cookie.Value, "=")+1:]
		log.Println(tokenVal)
		token, err := jwt.Parse(tokenVal, func(t *jwt.Token) (interface{}, error) {
			_, ok := t.Method.(*jwt.SigningMethodHMAC)
			if !ok {
				log.Println("not ok")
				w.WriteHeader(http.StatusUnauthorized)
			}
			return []byte(refresh.JWTSecret), nil

		})

		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		if token.Valid {
			next.ServeHTTP(w, r)
		}
	})
}
