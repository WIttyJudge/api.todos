package usecase

import (
	"errors"

	"github.com/wittyjudge/todo-api/internal/app/todo-api/domain/entities"
	"github.com/wittyjudge/todo-api/internal/app/todo-api/domain/repository"
	"github.com/wittyjudge/todo-api/internal/app/todo-api/domain/service"
)

var (
	ErrUserNotFound      = errors.New("User does not exist")
	ErrIncorrectPassword = errors.New("Incorrect password")
)

type UserUsecase interface {
	Login(user *entities.User) (string, error)

	Store(user *entities.User) error
}

type userUsecase struct {
	repo    repository.UserRepository
	service service.UserService
	jwt     JWTUsecase
}

func NewUserUsecase(repo repository.UserRepository, service service.UserService) UserUsecase {
	return &userUsecase{
		repo:    repo,
		service: service,
		jwt:     NewJWTUsecase(),
	}
}

func (u *userUsecase) Store(user *entities.User) error {
	if err := user.Validate(); err != nil {
		return err
	}

	password, err := u.service.EncryptString(user.Password)
	if err != nil {
		return err
	}

	user.EncryptedPassword = password

	return u.repo.Store(user)
}

func (u *userUsecase) Login(user *entities.User) (string, error) {
	findedUser, err := u.repo.FindByNickname(user.Nickname)
	if err != nil {
		return "", ErrUserNotFound
	}

	err = u.service.CompatePassword(findedUser.EncryptedPassword, user.Password)
	if err != nil {
		return "", ErrIncorrectPassword
	}

	accessToken, err := u.jwt.GenerateJWT(findedUser)
	if err != nil {
		return "", err
	}

	return accessToken, nil
}
