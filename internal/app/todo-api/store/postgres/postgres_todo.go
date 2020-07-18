package postgres

import (
	"github.com/jmoiron/sqlx"
	"github.com/wittyjudge/todo-api/internal/app/todo-api/domain/entities"
	"github.com/wittyjudge/todo-api/internal/app/todo-api/domain/repository"
)

type postgresTodo struct {
	db *sqlx.DB
}

func NewPostgresTodo(db *sqlx.DB) repository.TodoRepository {
	return &postgresTodo{db}
}

func (psql *postgresTodo) FetchAll() ([]entities.Todo, error) {
	todos := []entities.Todo{}
	err := psql.db.Select(&todos, "SELECT * FROM todos")
	if err != nil {
		return nil, err
	}

	return todos, nil
}

func (psql *postgresTodo) Store(todo *entities.Todo) error {
	sql := `INSERT INTO todos (title, task) VALUES ($1, $2)`

	_, err := psql.db.Exec(sql, &todo.Title, &todo.Task)
	if err != nil {
		return err
	}

	return nil
}

func (psql *postgresTodo) Delete(id int) (bool, error) {
	sql := `DELETE FROM todos WHERE id = $1`
	psql.db.MustExec(sql, id)

	// _, err := psql.db.Exec(sql, id)
	// if err != nil {
	// 	return false, err
	// }

	return true, nil
}
