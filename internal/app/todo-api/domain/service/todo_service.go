package service

import (
	"github.com/wittyjudge/todo-api/internal/app/todo-api/domain/entities"
	"github.com/wittyjudge/todo-api/internal/app/todo-api/domain/repository"
)

type TodoService struct {
	repo repository.TodoRepository
}

func (s *TodoService) isEmpty(todo *entities.Todo) string {
	if len(todo.Title) == 0 {
		return "title cannot be empty"
	}

	return ""
}
