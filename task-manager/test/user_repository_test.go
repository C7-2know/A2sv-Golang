// +build secret
package test

import (
	"context"
	domain "task_manager/Domain"
	repository "task_manager/Repositories"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UserRepositoryTestSuite struct {
	suite.Suite
	db *mongo.Database
}

func (suite *UserRepositoryTestSuite) SetupTest() {
	options := options.Client().ApplyURI("mongodb://localhost:27017")

	client, _ := mongo.Connect(context.Background(), options)
	db := client.Database("test")
	suite.db = db

}

func (suite *UserRepositoryTestSuite) TestCreateUser() {
	user := domain.User{
		Name:     "test",
		Email:    "test@gmail.com",
		Password: "password",
		Role:     "",
	}
	userRepo := repository.NewUserRepository(suite.db)
	err := userRepo.CreateUser(user)
	assert.NoError(suite.T(), err)
}

func (suite *UserRepositoryTestSuite) TestCreateUser_fail() {
	user := domain.User{
		Name:     "test",
		Password: "password",
		Role:     "",
	}
	userRepo := repository.NewUserRepository(suite.db)
	err := userRepo.CreateUser(user)
	assert.Error(suite.T(), err)
}

func (suite *UserRepositoryTestSuite) TestGetUserByEmail() {

	userRepo := repository.NewUserRepository(suite.db)
	_, err := userRepo.GetUserByEmail("test@gmail.com")
	assert.NoError(suite.T(), err)
}

func (suite *UserRepositoryTestSuite) TestGetUserByEmail_fail() {

	userRepo := repository.NewUserRepository(suite.db)
	_, err := userRepo.GetUserByEmail("tes@gmail.com")
	assert.Error(suite.T(), err)
}
func (suite *UserRepositoryTestSuite) TestUpdateUser() {
	user := domain.User{
		Name:     "test",
		Email:    "test@gmail.com",
		Password: "password",
		Role:     "",
	}
	userRepo := repository.NewUserRepository(suite.db)
	err := userRepo.UpdateUser("test@gmail.com", user)
	assert.NoError(suite.T(), err)
}

func (suite *UserRepositoryTestSuite) TestUpdateUser_fail() {
	user := domain.User{
		Name:     "test",
		Email:    "nw@gmail.com",
		Password: "password",
	}
	userRepo := repository.NewUserRepository(suite.db)
	err := userRepo.UpdateUser("unkown@gmail.com", user)
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), "no user found", err.Error())
}

func TestUserRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(UserRepositoryTestSuite))
}

func (suite *UserRepositoryTestSuite) TearDownTest() {
	suite.db.Drop(context.Background())
}
