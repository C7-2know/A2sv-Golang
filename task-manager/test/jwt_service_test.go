package test

import (
	domain "task_manager/Domain"
	service "task_manager/Infrastructures"
	"testing"

	"github.com/stretchr/testify/suite"
)

type JwtServiceTestSuite struct {
	suite.Suite
	jwtService domain.JWTService
}

func (suite *JwtServiceTestSuite) SetupTest() {
	suite.jwtService = service.NewJwtService()
}

func (suite *JwtServiceTestSuite) TestGenerateToken() {
	user := domain.User{
		Name:     "Test",
		Email:    "test@gmail.com",
		Password: "123456",
		Role:     "admin",
	}
	token, err := suite.jwtService.GenerateToken(user)
	suite.NoError(err)
	suite.NotEmpty(token)

}

func (suite *JwtServiceTestSuite) TestValidateToken() {
	role := ""
	token, _ := suite.jwtService.GenerateToken(domain.User{"Test", "test@gmail.com", "1345682", ""})
	headers := "Bearer " + token
	err := suite.jwtService.ValidateToken(role, headers)
	suite.NoError(err)

}

func TestJwtServiceTestSuite(t *testing.T) {
	suite.Run(t, new(JwtServiceTestSuite))
}

func (suite *JwtServiceTestSuite) TearDownTest() {
	suite.jwtService = nil
}
