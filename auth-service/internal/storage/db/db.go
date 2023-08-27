package db

import (
	"auth-service/internal/storage/db/postgres"
	"auth-service/internal/storage/db/redis"
)

type CredentionalsStorer interface {
	LookForAdmin(login, pwd string) error
	LookForManager(login, pwd string) error
	RegisterManager(login, pwd string) error
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
