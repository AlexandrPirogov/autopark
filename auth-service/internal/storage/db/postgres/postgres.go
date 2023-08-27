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
	URL := "postgresql://postgres:postgres@auth-postgres:5432/postgres"
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

func (pg *pgconn) LookForAdmin(login, pwd string) error {
	var id int
	err := pg.conn.QueryRow(context.Background(), QueryLookForAdmin, login, pwd).Scan(&id)
	log.Println(err)
	return err
}

// 2 distinct pg queries...
func (pg *pgconn) RegisterManager(login, pwd string) (int, error) {
	_, err := pg.conn.Exec(context.Background(), QueryInsertNewManager, login, pwd)
	log.Println(err)
	return pg.LookForManager(login, pwd)
}

func (pg *pgconn) LookForManager(login, pwd string) (int, error) {
	var id int
	err := pg.conn.QueryRow(context.Background(), QueryLookForManager, login, pwd).Scan(&id)
	log.Println(err)
	return id, err
}
