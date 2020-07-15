package postgres

import (
	"github.com/jmoiron/sqlx"
	"github.com/wittyjudge/todo-api/internal/app/todo-api/domain/entities"
	"github.com/wittyjudge/todo-api/internal/app/todo-api/domain/repository"
)

type postgresUser struct {
	db *sqlx.DB
}

func NewPostgresUser(db *sqlx.DB) repository.UserRepository {
	return &postgresUser{db}
}

func (psql *postgresUser) FindByNickname(nickname string) (*entities.User, error) {
	user := &entities.User{}

	// err := psql.db.QueryRow("SELECT nickname, encrypted_password FROM users WHERE nickname = $1", nickname).Scan(&user.Nickname)
	// if err != nil {
	// 	return nil, err
	// }

	sql := `SELECT id, nickname, encrypted_password FROM users WHERE nickname = $1`
	err := psql.db.QueryRow(sql, nickname).Scan(&user.ID, &user.Nickname, &user.EncryptedPassword)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (psql *postgresUser) Store(user *entities.User) error {
	sql := `INSERT INTO users(nickname, encrypted_password) VALUES ($1, $2)`
	_, err := psql.db.Exec(sql, &user.Nickname, &user.EncryptedPassword)
	if err != nil {
		return nil
	}
	return nil
}
