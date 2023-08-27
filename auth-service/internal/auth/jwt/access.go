package jwt

import (
	"log"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const AccessTokenCookieField = "access-token"
const JWTAccesssSecret = "super-secret-access-auth-key"

func GenerateAccessToken() (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	claims["exp"] = time.Now().Add(time.Minute * 5).Unix()

	tokenStr, err := token.SignedString([]byte(JWTSecret))
	log.Printf("generated access token. err: %v", err)
	return tokenStr, err
}
