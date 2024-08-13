package service

import (
	"errors"
	domain "task_manager/Domain"

	"golang.org/x/crypto/bcrypt"
)
type passwordService struct{
}

func NewPasswordService() domain.PasswordService{
	return &passwordService{}
}

func (*passwordService) HashPassword(password string) (string, error) {
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPass), nil
}

func (*passwordService)ComparePassword(pass1, pass2 string) error{
	if bcrypt.CompareHashAndPassword([]byte(pass1), []byte(pass2)) != nil {
		return errors.New("passwords do not match")
	}
	return nil

}
