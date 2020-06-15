package usecase

import (
	"errors"

	"github.com/wittyjudge/todo-api/internal/app/todo-api/domain/entities"
	"github.com/wittyjudge/todo-api/internal/app/todo-api/domain/repository"
	"github.com/wittyjudge/todo-api/internal/app/todo-api/domain/service"
)

var (
	ErrUserNotFound = errors.New("User does not exist")
)

type UserUsecase interface {
	Login(user *entities.User) (*entities.User, error)

	Store(user *entities.User) error
}

type userUsecase struct {
	repo    repository.UserRepository
	service service.UserService
}

func NewUserUsecase(repo repository.UserRepository, service service.UserService) UserUsecase {
	return &userUsecase{repo, service}
}

func (u *userUsecase) Store(user *entities.User) error {
	if err := u.service.PasswordValidation(user.Password); err != nil {
		return err
	}

	password, err := u.service.EncryptString(user.Password)
	if err != nil {
		return err
	}

	user.EncryptedPassword = password

	return u.repo.Store(user)
}

func (u *userUsecase) Login(user *entities.User) (*entities.User, error) {
	findedUser, err := u.repo.FindByNickname(user.Nickname)
	if err != nil {
		return nil, ErrUserNotFound
	}

	err = u.service.CompatePassword(findedUser.EncryptedPassword, user.Password)
	if err != nil {
		return nil, err
	}
	// if !exists {
	// 	return nil, errors.New("Password incorrect")
	// }

	return findedUser, nil
}
