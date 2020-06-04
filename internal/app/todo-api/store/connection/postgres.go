package connection

import (
	"fmt"
	"log"

	"github.com/jackc/pgx"
)

func Connect() *pgx.Conn {
	db := open("localhost", "5432", "postgres", "admin", "todo")
	return db
}

func open(host, port, user, password, dbname string) *pgx.Conn {
	connSrt := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s", host, port, user, password, dbname)
	conn, err := pgx.ParseConnectionString(connSrt)
	if err != nil {
		log.Fatal(err)
	}

	db, err := pgx.Connect(conn)
	if err != nil {
		log.Fatal(err)
	}

	return db
}
