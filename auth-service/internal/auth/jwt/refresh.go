package jwt

import (
	"log"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const RerfeshTokenCookieField = "refresh-token"
const JWTSecret = "super-secret-auth-key"

func GenerateRefreshToken() (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	tokenStr, err := token.SignedString([]byte(JWTSecret))
	log.Printf("generated token. err: %v", err)
	return tokenStr, err
}
