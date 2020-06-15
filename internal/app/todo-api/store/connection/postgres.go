package connection

import (
	"log"

	"github.com/jackc/pgx"
	"github.com/wittyjudge/todo-api/internal/app/todo-api/utils"
)

// Connect inits connection to PostgreSQL
func Connect() *pgx.Conn {
	// db := open("localhost", "5432", "postgres", "admin", "todo")
	conn := utils.GetEnv("POSTGRES_URL")
	db := open(conn)
	return db
}

func open(connStr string) *pgx.Conn {
	// connSrt := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s", host, port, user, password, dbname)
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
