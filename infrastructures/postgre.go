package infrastructures

import (
	"fmt"

	"github.com/gocraft/dbr/v2"
	_ "github.com/lib/pq"
)

var (
	host     = "localhost"
	port     = 5432
	user     = "amalia"
	password = "123"
	dbname   = "erp"
)

var psqlInfo = fmt.Sprintf("host=%s port=%d user=%s "+
	"password=%s dbname=%s sslmode=disable",
	host, port, user, password, dbname)

func AttendanceConn() *dbr.Session {

	conn, err := dbr.Open("postgres", psqlInfo, nil)
	if err != nil {
		panic(err)
	}

	conn.SetMaxOpenConns(10)
	sess := conn.NewSession(nil)
	return sess
}
