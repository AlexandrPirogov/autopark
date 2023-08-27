// postgres holds CRUD functionality to work with Postgres
package postgres

import (
	"context"
	"log"
	"sync"

	"github.com/jackc/pgx/v5/pgxpool"
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
	URL := "postgresql://postgres:postgres@autoparkdb:5432/postgres"
	conn, err := pgxpool.New(context.Background(), URL)
	if err != nil {
		log.Fatal(err)
	}

	ping := conn.Ping(context.Background())
	if ping != nil {
		log.Fatalf("error while pinging %v", ping)
	}
	return conn
}
