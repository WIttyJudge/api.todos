package repository

import (
	"github.com/wittyjudge/todo-api/internal/app/todo-api/domain/entities"
)

// TodoRepository ...
type TodoRepository interface {
	FetchAll() ([]entities.Todo, error)

	Store(todo *entities.Todo) error
	Delete(id int) (bool, error)
}
