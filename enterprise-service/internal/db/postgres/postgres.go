package postgres

import (
	"context"
	"enterprise-service/internal/enterprise"
	"enterprise-service/internal/std"
	"log"
	"sync"

	"github.com/jackc/pgx/v5/pgxpool"
)

// Instance of singleton
var singletonConn *pgconn = nil
var initOnce sync.Once

func GetInstance() *pgconn {
	initOnce.Do(func() {
		singletonConn = new()
	})
	return singletonConn
}

type pgconn struct {
	conn *pgxpool.Pool
}

func new() *pgconn {
	return &pgconn{
		conn: tryConnect(),
	}
}

func tryConnect() *pgxpool.Pool {
	URL := "postgresql://postgres:postgres@localhost:10000/postgres"
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

func (pg *pgconn) Delete(e enterprise.Enterprise) error {
	return nil
}

func (pg *pgconn) Read(e enterprise.Enterprise) (std.Linked[enterprise.Enterprise], error) {
	return nil, nil
}

func (pg *pgconn) Store(e enterprise.Enterprise) error {
	return nil
}

func (pg *pgconn) Update(s enterprise.Enterprise) (std.Linked[enterprise.Enterprise], error) {
	return nil, nil
}
