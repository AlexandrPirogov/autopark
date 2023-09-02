// postgres package holds the functionality to work with postgres
package postgres

import (
	"auth-service/internal/config"
	"context"
	"sync"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog/log"
)

// Instance of singleton
var singletonConn *pgconn = nil
var initOnce sync.Once

func init() {
	GetInstance()
}

func GetInstance() *pgconn {
	initOnce.Do(func() {
		singletonConn = new()
	})
	return singletonConn
}

type pgconn struct {
	conn *pgxpool.Pool
}

// LookForAdmin searchs admin by given login and pwd.
//
// Pre-cond: given login and password
//
// Post-cond: executes query tos elect admin with given login and pwd.
// If admin was found, returns min otherwise returnes error
func (pg *pgconn) LookForAdmin(login, pwd string) error {
	var id int
	err := pg.conn.QueryRow(context.Background(), QueryLookForAdmin, login, pwd).Scan(&id)
	log.Warn().Msgf("%v", err)
	return err
}

// RegisterManager register new manager with given login and pwd
//
// Pre-cond: given login and pwd to register new manager. Login must be unique
//
// Post-cond: query was executed. If success, returns id of registere manager and nil.
// Otherwise returns error
func (pg *pgconn) RegisterManager(login, pwd string) (int, error) {
	_, err := pg.conn.Exec(context.Background(), QueryInsertNewManager, login, pwd)
	log.Warn().Msgf("%v", err)
	return pg.LookForManager(login, pwd)
}

// LookForManager looking for manager with given login and pwd
//
// Pre-cond: given login and pwd
//
// Post-cond: query was executed. If success, returns id of registere manager and nil.
// Otherwise returns error
func (pg *pgconn) LookForManager(login, pwd string) (int, error) {
	var id int
	err := pg.conn.QueryRow(context.Background(), QueryLookForManager, login, pwd).Scan(&id)
	log.Warn().Msgf("%v", err)
	return id, err
}

// RegisterClient register new Client with given login and pwd
//
// Pre-cond: given login and pwd to register new Client. Login must be unique
//
// Post-cond: query was executed. If success, returns id of registere Client and nil.
// Otherwise returns error
func (pg *pgconn) RegisterClient(login, pwd string) (int, error) {
	_, err := pg.conn.Exec(context.Background(), QueryInsertNewClient, login, pwd)
	log.Warn().Msgf("%v", err)
	return pg.LookForClient(login, pwd)
}

// LookForClient looking for Client with given login and pwd
//
// Pre-cond: given login and pwd
//
// Post-cond: query was executed. If success, returns id of registere Client and nil.
// Otherwise returns error
func (pg *pgconn) LookForClient(login, pwd string) (int, error) {
	var id int
	err := pg.conn.QueryRow(context.Background(), QueryLookForClient, login, pwd).Scan(&id)
	log.Warn().Msgf("%v", err)
	return id, err
}

// new returns pointer to the new instance of postrgres connection
// or existing one if it was created earlier.
func new() *pgconn {
	return &pgconn{
		conn: tryConnect(),
	}
}

// tryConnect checks connectnions
func tryConnect() *pgxpool.Pool {
	URL := config.PostgresURL()
	conn, err := pgxpool.New(context.Background(), URL)
	if err != nil {
		log.Fatal().Msgf("%v", err)
	}

	log.Warn().Msgf("pinging db %s", URL)
	ping := conn.Ping(context.Background())
	if ping != nil {
		log.Fatal().Msgf("error while pinging %v", ping)
	}
	log.Warn().Msg("connected to postgres db successfully...")

	return conn
}
