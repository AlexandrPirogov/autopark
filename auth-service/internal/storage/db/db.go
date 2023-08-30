// package db collectects interfaces for working with
// users' authentication
package db

import (
	"auth-service/internal/storage/db/postgres"
	"auth-service/internal/storage/db/redis"
)

// CredentionalsStorer is interface for working with db and users of system
type CredentionalsStorer interface {
	// LookForAdmin searchs admin by given login and pwd.
	//
	// Pre-cond: given login and password
	//
	// Post-cond: executes query tos elect admin with given login and pwd.
	// If admin was found, returns min otherwise returnes error
	LookForAdmin(login, pwd string) error

	// LookForManager looking for manager with given login and pwd
	//
	// Pre-cond: given login and pwd
	//
	// Post-cond: query was executed. If success, returns id of registere manager and nil.
	// Otherwise returns error
	LookForManager(login, pwd string) (int, error)

	// LookForClient looking for Client with given login and pwd
	//
	// Pre-cond: given login and pwd
	//
	// Post-cond: query was executed. If success, returns id of Client and nil.
	// Otherwise returns error
	LookForClient(login, pwd string) (int, error)

	// RegisterClient register new Client with given login and pwd
	//
	// Pre-cond: given login and pwd to register new Client. Login must be unique
	//
	// Post-cond: query was executed. If success, returns id of registered Client and nil.
	// Otherwise returns error
	RegisterClient(login, pwd string) (int, error)

	// RegisterClient register new Client with given login and pwd
	//
	// Pre-cond: given login and pwd to register new Client. Login must be unique
	//
	// Post-cond: query was executed. If success, returns id of registere Client and nil.
	// Otherwise returns error
	RegisterManager(login, pwd string) (int, error)
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
