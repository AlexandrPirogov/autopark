// postgres holds CRUD functionality to work with Postgres
package postgres

import (
	"booking-service/internal/config"
	"context"
	"sync"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog/log"
)

// Instance of singleton
var singletonConn *pgconn = nil
var initOnce sync.Once

// GetInstance returns *pgconn
// If pgconn was created returns the old one
func GetInstance() *pgconn {
	initOnce.Do(func() {
		singletonConn = new()
	})
	return singletonConn
}

// pgconn just a wrapper for pgxpool.Pool
type pgconn struct {
	conn *pgxpool.Pool
}

// new creates new connection to Postgres via pool
func new() *pgconn {
	return &pgconn{
		conn: tryConnect(),
	}
}

// tryConnect trying to ping to verify that database is avaible
// and given URL is correct
func tryConnect() *pgxpool.Pool {
	URL := config.PostgresURL()
	conn, err := pgxpool.New(context.Background(), URL)
	if err != nil {
		log.Fatal().Msgf("%v", err)
	}

	ping := conn.Ping(context.Background())
	if ping != nil {
		log.Fatal().Msgf("error while pinging %v", ping)
	}
	return conn
}
