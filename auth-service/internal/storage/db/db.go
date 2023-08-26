package db

import (
	"auth-service/internal/storage/db/postgres"
	"auth-service/internal/storage/db/redis"
)

type CredentionalsStorer interface {
	LookForAdmin(login, pwd string) bool
}

type JWTTokenStorer interface {
	SetRefreshToken(val string) error
}

func GetCurrentCredsStorerInstance() CredentionalsStorer {
	return postgres.GetInstance()
}

func GetCurrentJWTStorerInstance() JWTTokenStorer {
	return redis.GetInstance()
}

func SetRefreshToken(val string) error {
	return
}
