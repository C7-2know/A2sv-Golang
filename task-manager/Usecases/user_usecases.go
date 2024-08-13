package usecases

import (
	domain "task_manager/Domain"
)

type userUseCase struct {
	user_repo        domain.UserRepository
	password_service domain.PasswordService
	jwt_service      domain.JWTService
}

func NewUserUsecase(user_repo domain.UserRepository, pass domain.PasswordService,
	jwt domain.JWTService) domain.UserUsecase {
	return &userUseCase{user_repo: user_repo, password_service: pass, jwt_service: jwt}
}

func (uu *userUseCase) CreateUser(user domain.User) error {
	hashed_pass,err:=uu.password_service.HashPassword(user.Password)
	if err!=nil{
		return err
	}
	user.Password=hashed_pass
	return uu.user_repo.CreateUser(user)
}

func (uu *userUseCase) GetUserByEmail(email string) (domain.User, error) {
	return uu.user_repo.GetUserByEmail(email)
}

func (uu *userUseCase) Login(email, password string) (string, error) {
	user, err := uu.user_repo.GetUserByEmail(email)
	if err != nil {
		return "", err
	}
	err = uu.password_service.ComparePassword(user.Password, password)
	if err != nil {
		return "", err
	}
	token, err := uu.jwt_service.GenerateToken(user)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (uu *userUseCase) Promote(email string) error {
	user, err := uu.user_repo.GetUserByEmail(email)
	if err != nil {
		return err
	}
	user.Role = "admin"
	err = uu.user_repo.UpdateUser(email, user)
	if err != nil {
		return err
	}
	return nil
}
