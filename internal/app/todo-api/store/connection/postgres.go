package connection

import (
	"fmt"
	"log"

	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/wittyjudge/todo-api/internal/app/todo-api/utils"
)

// Connect inits connection to PostgreSQL.
func Connect() *sqlx.DB {
	user := utils.GetEnv("POSTGRES_USER")
	password := utils.GetEnv("POSTGRES_PASSWORD")
	host := utils.GetEnv("POSTGRES_HOST")
	port := utils.GetEnv("POSTGRES_PORT")
	dbName := utils.GetEnv("POSTGRES_DB")

	// postgres://postgres:admin@localhost:5432/todo
	connSrt := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", user, password, host, port, dbName)
	db := open(connSrt)
	return db
}

func open(connStr string) *sqlx.DB {
	db, err := sqlx.Connect("pgx", connStr)
	if err != nil {
		log.Fatal(err)
	}

	// conn, err := pgx.ParseConnectionString(connStr)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// db, err := pgx.Connect(conn)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	return db
}
