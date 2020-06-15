package repository

import "github.com/wittyjudge/todo-api/internal/app/todo-api/domain/entities"

type UserRepository interface {
	FindByNickname(nickname string) (*entities.User, error)

	Store(user *entities.User) error
}
