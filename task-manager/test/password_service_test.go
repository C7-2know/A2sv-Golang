package test

import (
	domain "task_manager/Domain"
	service "task_manager/Infrastructures"
	"testing"

	"github.com/stretchr/testify/suite"
)

type PasswordServiceTestSuite struct {
	suite.Suite
	pass_service domain.PasswordService
}

func (suite *PasswordServiceTestSuite) SetupTest() {
	suite.pass_service = service.NewPasswordService()

}
func (suite *PasswordServiceTestSuite) TestHashPassword() {
	pswd := "123456"
	hashed_pswd, err := suite.pass_service.HashPassword(pswd)
	suite.NoError(err)
	suite.NotEmpty(hashed_pswd)
}
func (suite *PasswordServiceTestSuite) TestComparePassword() {
	pswd := "123456"
	hashed_pswd, _ := suite.pass_service.HashPassword(pswd)
	err := suite.pass_service.ComparePassword( hashed_pswd,pswd)
	suite.NoError(err)

}

func (suite *PasswordServiceTestSuite) TearDownTest() {
	suite.pass_service = nil
}

func TestPasswordServiceTestSuite(t *testing.T) {
	suite.Run(t, new(PasswordServiceTestSuite))
}
