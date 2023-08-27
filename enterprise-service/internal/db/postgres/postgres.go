package postgres

import (
	"context"
	"enterprise-service/internal/enterprise"
	"enterprise-service/std"
	"enterprise-service/std/list"
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
	URL := "postgresql://postgres:postgres@enterprise-db:5432/postgres"
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
	rows, err := pg.conn.Query(context.Background(), QueryReadEnterprises)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	res := list.New[enterprise.Enterprise]()
	for rows.Next() {
		var e enterprise.Enterprise
		scanErr := rows.Scan(&e.Title)
		if scanErr != nil {
			log.Println(err)
			continue
		}
		res.PushBack(e)
	}

	return res, nil
}

func (pg *pgconn) ReadByID(id int) (enterprise.Enterprise, error) {
	var e enterprise.Enterprise
	err := pg.conn.QueryRow(context.Background(), QueryReadByIDEnterprises, id).Scan(&e.Title)
	return e, err
}

func (pg *pgconn) Store(e enterprise.Enterprise) error {
	_, err := pg.conn.Exec(context.Background(), QueryStoreEnterprise, e.Title)
	if err != nil {
		return err
	}
	return nil
}

func (pg *pgconn) Update(s enterprise.Enterprise) (std.Linked[enterprise.Enterprise], error) {
	return nil, nil
}
