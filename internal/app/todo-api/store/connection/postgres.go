package connection

import (
	"fmt"
	"log"

	"github.com/jackc/pgx"
	"github.com/wittyjudge/todo-api/internal/app/todo-api/utils"
)

// Connect inits connection to PostgreSQL.
func Connect() *pgx.Conn {
	host := utils.GetEnv("POSTGRES_HOST")
	port := utils.GetEnv("POSTGRES_PORT")
	user := utils.GetEnv("POSTGRES_USER")
	password := utils.GetEnv("POSTGRES_PASSWORD")
	dbName := utils.GetEnv("POSTGRES_DB")

	connSrt := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s", host, port, user, password, dbName)
	db := open(connSrt)
	return db
}

func open(connStr string) *pgx.Conn {
	conn, err := pgx.ParseConnectionString(connStr)
	if err != nil {
		log.Fatal(err)
	}

	db, err := pgx.Connect(conn)
	if err != nil {
		log.Fatal(err)
	}

	return db
}
