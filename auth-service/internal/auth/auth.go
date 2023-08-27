package auth

import (
	"auth-service/internal/auth/refresh"
	"auth-service/internal/storage/db"
	"fmt"
)

func VerifyCredentionals(login, pwd string, s db.CredentionalsStorer) (string, error) {
	res := s.LookForAdmin(login, pwd)
	if !res {
		return "", fmt.Errorf("not found")
	}

	refresh, err := refresh.GenerateRefreshToken()
	if err != nil {
		return "", err
	}

	return refresh, nil
}

func StoreRefreshToken(token string, s db.JWTTokenStorer) error {
	return s.SetRefreshToken(token)
}
