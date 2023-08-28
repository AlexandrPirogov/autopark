// postgres package holds the functionality to work with postgres
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

// LookForAdmin searchs admin by given login and pwd.
//
// Pre-cond: given login and password
//
// Post-cond: executes query tos elect admin with given login and pwd.
// If admin was found, returns min otherwise returnes error
func (pg *pgconn) LookForAdmin(login, pwd string) error {
	var id int
	err := pg.conn.QueryRow(context.Background(), QueryLookForAdmin, login, pwd).Scan(&id)
	log.Println(err)
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
	log.Println(err)
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
	log.Println(err)
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
	log.Println(err)
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
	log.Println(err)
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
