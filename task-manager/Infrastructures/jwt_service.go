package service

import (
	"errors"
	"fmt"
	"strings"
	domain "task_manager/Domain"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var JwtSecret = "4533413ksandukjhuweas4const6"

type jwtService struct{

}

func NewJwtService() domain.JWTService{
	return &jwtService{}
}

func (js *jwtService) GenerateToken(user domain.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"email": user.Email,
			"role":  user.Role,
			"exp":   time.Now().Add(time.Hour).Unix(),
		})
	jwtToken, err := token.SignedString(JwtSecret)
	if err != nil {
		return "", err
	}

	return jwtToken, nil
}

func (js *jwtService) ValidateToken(role string,headers string) error {
	if headers == "" {
		return errors.New("Authorization header is required")
	}

	authParts := strings.Split(headers, " ")
	if len(authParts) != 2 || strings.ToLower(authParts[0]) != "bearer" {
		return errors.New("Authorization header format must be Bearer {token}")
	}
	token, err := jwt.Parse(authParts[1], func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(JwtSecret), nil
	})
	if err != nil {
		return err
	}
	if role!=""{
		if claims,ok:= token.Claims.(jwt.MapClaims); ok && token.Valid{
			// if token is expired
			if float64(time.Now().Unix()) > claims["exp"].(float64){
				return errors.New("Token expired")
			}
			// check if the role matches
			if claims["role"]!=role{
				return errors.New("Unauthorized user")
			}
		}
	}
	return nil
}
