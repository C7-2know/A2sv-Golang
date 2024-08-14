package service

import (
	"errors"
	"fmt"
	"os"
	"strings"
	domain "task_manager/Domain"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// var JwtSecret = []byte("4533413ksandukjhuweas4const6")

var JwtSecret= []byte(os.Getenv("JwtSecret"))
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
		return "", errors.New("could not generate token")
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
	if claims,ok:= token.Claims.(jwt.MapClaims); ok && token.Valid{
		// if token is expired
		if float64(time.Now().Unix()) > claims["exp"].(float64){
			return errors.New("Token expired")
		}
		// check if the role matches
		if role!="" && claims["role"]!=role{
			return errors.New("Unauthorized user")
		}
	}
	
	return nil
}
