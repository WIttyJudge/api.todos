package service

import (
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
		return err
	}

	return nil
}
