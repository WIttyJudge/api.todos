package postgres

import (
	"github.com/jackc/pgx"
	"github.com/wittyjudge/todo-api/internal/app/todo-api/domain/entities"
	"github.com/wittyjudge/todo-api/internal/app/todo-api/domain/repository"
)

type postgresTodo struct {
	db *pgx.Conn
}

func NewPostgresTodo(db *pgx.Conn) repository.TodoRepository {
	return &postgresTodo{db}
}

func (psql *postgresTodo) FetchAll() ([]entities.Todo, error) {
	todos := make([]entities.Todo, 0)

	sql := `SELECT * FROM todos ORDER BY id DESC`
	rows, err := psql.db.Query(sql)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		todo := &entities.Todo{}
		if err := rows.Scan(&todo.ID, &todo.Title, &todo.Task, &todo.Completed, &todo.CreatedAt, &todo.CompletedAt); err != nil {
			return nil, err
		}

		todos = append(todos, *todo)
	}

	if rows.Err() != nil {
		return nil, rows.Err()
	}

	return todos, nil
}

func (psql *postgresTodo) Store(todo *entities.Todo) error {
	sql := `INSERT INTO todos(title, task) VALUES ($1, $2)`
	_, err := psql.db.Exec(sql, &todo.Title, &todo.Task)
	if err != nil {
		return err
	}

	return nil
}

func (psql *postgresTodo) Delete(id int) (bool, error) {
	sql := `DELETE FROM todos WHERE id = $1`
	_, err := psql.db.Exec(sql, id)
	if err != nil {
		return false, err
	}

	return true, nil
}
