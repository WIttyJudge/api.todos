package usecase

import (
	"github.com/wittyjudge/todo-api/internal/app/todo-api/domain/entities"
	"github.com/wittyjudge/todo-api/internal/app/todo-api/domain/repository"
)

type TodoUsecase interface {
	FetchAll() ([]entities.Todo, error)

	Store(todo *entities.Todo) error
	Delete(id int) (bool, error)
}

type todoUsecase struct {
	repo repository.TodoRepository
}

func NewTodoUsecase(repo repository.TodoRepository) TodoUsecase {
	return &todoUsecase{repo}
}

func (u *todoUsecase) FetchAll() ([]entities.Todo, error) {
	return u.repo.FetchAll()
}

func (u *todoUsecase) Store(todo *entities.Todo) error {
	return u.repo.Store(todo)
}

func (u *todoUsecase) Delete(id int) (bool, error) {
	return u.repo.Delete(id)
}
