package service

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type UserService struct{}

func (s *UserService) EncryptString(str string) (string, error) {
	b, err := bcrypt.GenerateFromPassword([]byte(str), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(b), nil
}

func (s *UserService) CompatePassword(hashedStr, str string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(hashedStr), []byte(str)); err != nil {
		return errors.New("Incorrect password")
	}

	return nil
}

func (s *UserService) PasswordValidation(password string) error {
	if len(password) == 0 {
		return errors.New("Required password")
	}

	if len(password) < 5 {
		return errors.New("Password must have more characters then 5")
	}

	return nil
}
