// auth hold functionallity for authentication and token generation
package auth

import (
	"auth-service/internal/auth/jwt"
	"auth-service/internal/storage/db"
	"fmt"
)

type ClientCreds struct {
	UserID int    `json:"u_id"`
	jwt    string `json: "-"`
}

func (c ClientCreds) RefreshToken() string {
	return c.jwt
}

// VerifyAdminCredentionals checks if admin's creds are correct
//
// Pre-cond: given admin's login and pwd and CredentionalsStorer to search in
//
// Post-cond: if creds are correct - generate refresh token and returns it with nil
// Otherwise returns empty string and error
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

// RegisterManager register new manager in DB
//
// Pre-cond: given unique login and pwd for new Manager and CredentionalsStorer to store in
//
// Post-cond: if manager was written to CredentaionalsStorer successfully then returns id of manager and nil.
// Otherwise returns error
func RegisterManager(login, pwd string, s db.CredentionalsStorer) (int, error) {
	return s.RegisterManager(login, pwd)
}

// VerifyManagerCredentionals checks if managers's creds are correct
//
// Pre-cond: given managers's login and pwd and CredentionalsStorer to search in
//
// Post-cond: if creds are correct - generate refresh token and returns it with nil
// Otherwise returns empty string and error
func VerifyManagerCredentionals(login, pwd string, s db.CredentionalsStorer) (string, error) {
	id, err := s.LookForManager(login, pwd)
	if err != nil {
		return "", fmt.Errorf("not found")
	}

	claims := map[string]any{
		"id": id,
	}
	refresh, err := jwt.GenerateRefreshWithClaimsToken(claims)
	if err != nil {
		return "", err
	}

	return refresh, nil
}

// RegisterClient register new Client in DB
//
// Pre-cond: given unique login and pwd for new Client and CredentionalsStorer to store in
//
// Post-cond: if Client was written to CredentaionalsStorer successfully then returns id of Client and nil.
// Otherwise returns error
func RegisterClient(login, pwd string, s db.CredentionalsStorer) (int, error) {
	return s.RegisterClient(login, pwd)
}

// VerifyClientCredentionals checks if Clients's creds are correct
//
// Pre-cond: given Clients's login and pwd and CredentionalsStorer to search in
//
// Post-cond: if creds are correct - generate refresh token and returns it with nil
// Otherwise returns empty string and error
func VerifyClientCredentionals(login, pwd string, s db.CredentionalsStorer) (ClientCreds, error) {
	var c ClientCreds
	id, err := s.LookForClient(login, pwd)
	if err != nil {
		return c, fmt.Errorf("not found")
	}

	refresh, err := jwt.GenerateRefreshToken()
	if err != nil {
		return c, err
	}
	c.UserID = id
	c.jwt = refresh

	return c, nil
}

func StoreRefreshToken(token string, s db.JWTTokenStorer) error {
	return s.SetRefreshToken(token)
}
