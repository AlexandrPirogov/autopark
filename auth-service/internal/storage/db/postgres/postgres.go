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
	URL := "postgresql://postgres:postgres@localhost:10002/postgres"
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

func (pg *pgconn) LookForAdmin(login, pwd string) bool {
	var uid []byte
	err := pg.conn.QueryRow(context.Background(), QueryLookForAdmin, login, pwd).Scan(&uid)
	log.Println(err)
	return err == nil
}
