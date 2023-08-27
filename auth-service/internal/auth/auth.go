package auth

import (
	"auth-service/internal/auth/jwt"
	"auth-service/internal/storage/db"
	"fmt"
)

func VerifyAdminCredentionals(login, pwd string, s db.CredentionalsStorer) (string, error) {
	err := s.LookForAdmin(login, pwd)
	if err != nil {
		return "", fmt.Errorf("not found")
	}

	refresh, err := jwt.GenerateRefreshToken()
	if err != nil {
		return "", err
	}

	return refresh, nil
}

func RegisterManager(login, pwd string, s db.CredentionalsStorer) (int, error) {
	return s.RegisterManager(login, pwd)
}

func VerifyManagerCredentionals(login, pwd string, s db.CredentionalsStorer) (string, error) {
	_, err := s.LookForManager(login, pwd)
	if err != nil {
		return "", fmt.Errorf("not found")
	}

	refresh, err := jwt.GenerateRefreshToken()
	if err != nil {
		return "", err
	}

	return refresh, nil
}

func StoreRefreshToken(token string, s db.JWTTokenStorer) error {
	return s.SetRefreshToken(token)
}
