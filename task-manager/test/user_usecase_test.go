package test

import (
	"errors"
	"fmt"
	domain "task_manager/Domain"
	usecases "task_manager/Usecases"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) CreateUser(user domain.User) error {
	args := m.Called(user)
	return args.Error(0)
}

func (m *MockUserRepository) GetUserByEmail(email string) (domain.User, error) {
	args := m.Called(email)
	return args.Get(0).(domain.User), args.Error(1)
}

func (m *MockUserRepository) UpdateUser(email string, user domain.User) error {
	args := m.Called(email, user)

	return args.Error(0)
}

// password service
type MockPasswordService struct {
	mock.Mock
}

func (m *MockPasswordService) HashPassword(pass string) (string, error) {
	args := m.Called(pass)
	return args.String(0), args.Error(1)
}

func (m *MockPasswordService) ComparePassword(hashedPass string, plainPass string) error {
	args := m.Called(hashedPass, plainPass)
	return args.Error(0)
}

// jwt service
type MockJWTService struct {
	mock.Mock
}

func (m *MockJWTService) GenerateToken(user domain.User) (string, error) {
	args := m.Called(user)
	return args.String(0), args.Error(1)
}

func (m *MockJWTService) ValidateToken(role string, headers string) error {
	args := m.Called(role, headers)
	return args.Error(0)
}

// test user usecase

type UserUseCaseTestSuite struct {
	suite.Suite
	mockUserRepo        *MockUserRepository
	mockPasswordService *MockPasswordService
	mockJWTService      *MockJWTService
	userUseCase         domain.UserUsecase
}

func (suite *UserUseCaseTestSuite) SetupTest() {
	suite.mockUserRepo = new(MockUserRepository)
	suite.mockPasswordService = new(MockPasswordService)
	suite.mockJWTService = new(MockJWTService)
	suite.userUseCase = usecases.NewUserUsecase(suite.mockUserRepo, suite.mockPasswordService, suite.mockJWTService)
}

// create user
func (suite *UserUseCaseTestSuite) TestCreateUser_success() {
	user := domain.User{
		Email:    "test@gmail.com",
		Password: "password",
		Name:     "test",
		Role:     "user",
	}
	user2 := domain.User{
		Email:    "test@gmail.com",
		Password: "password",
		Name:     "test",
		Role:     "user",
	}

	hashed_password := "hashedpass123"
	suite.mockPasswordService.On("HashPassword", "password").Return(hashed_password, nil)
	user2.Password = hashed_password
	suite.mockUserRepo.On("CreateUser", user2).Return(nil)

	err := suite.userUseCase.CreateUser(user)
	assert.NoError(suite.T(), err)
	suite.mockPasswordService.AssertExpectations(suite.T())
	suite.mockUserRepo.AssertExpectations(suite.T())

}

func (suite *UserUseCaseTestSuite) TestCreateUser_fail_Hash_fail() {
	user := domain.User{
		Email:    "test@gmail.com",
		Password: "password",
		Name:     "test",
		Role:     "user",
	}
	suite.mockPasswordService.On("HashPassword", user.Password).Return("", assert.AnError)
	err := suite.userUseCase.CreateUser(user)
	fmt.Println("error", err)
	assert.Error(suite.T(), err)
	suite.mockPasswordService.AssertExpectations(suite.T())
}

// login
func (suite *UserUseCaseTestSuite) TestLogin_success() {
	user := domain.User{
		Email:    "test@gmail.com",
		Password: "pass",
		Name:     "test",
		Role:     "user",
	}
	token := "sometokengjkasbdekuad"
	suite.mockUserRepo.On("GetUserByEmail", user.Email).Return(user, nil)
	suite.mockPasswordService.On("ComparePassword", user.Password, "pass").Return(nil)
	suite.mockJWTService.On("GenerateToken", user).Return(token, nil)

	result, err := suite.userUseCase.Login(user.Email, user.Password)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), token, result)
	suite.mockUserRepo.AssertExpectations(suite.T())
	suite.mockPasswordService.AssertExpectations(suite.T())
	suite.mockJWTService.AssertExpectations(suite.T())

}


func (suite *UserUseCaseTestSuite) TestLogin_tokenError_fail(){
	user := domain.User{
		Email:    "test@gmail.com",
		Password: "pass",
		Name:     "test",
		Role:     "user",
	}
	suite.mockUserRepo.On("GetUserByEmail", user.Email).Return(user, nil)
	suite.mockPasswordService.On("ComparePassword", user.Password, "pass").Return(nil)
	suite.mockJWTService.On("GenerateToken",user).Return("",errors.New("token error"))
	_,err:=suite.userUseCase.Login(user.Email,user.Password)
	assert.Error(suite.T(),err)
	suite.mockJWTService.AssertExpectations(suite.T())
}




func (suite *UserUseCaseTestSuite) TestPromote_Success() {
    user := domain.User{
        Email:    "test@example.com",
        Password: "hashedpassword123",
        Name:     "Test User",
        Role:     "user",
    }

    suite.mockUserRepo.On("GetUserByEmail", user.Email).Return(user, nil)
    suite.mockUserRepo.On("UpdateUser", user.Email, mock.MatchedBy(func(u domain.User) bool {
        return u.Role == "admin"
    })).Return(nil)

    err := suite.userUseCase.Promote(user.Email)
    assert.NoError(suite.T(), err)
    suite.mockUserRepo.AssertExpectations(suite.T())
}

func (suite *UserUseCaseTestSuite) TestPromote_fail() {
    suite.mockUserRepo.On("GetUserByEmail", "unknown@example.com").Return(domain.User{}, errors.New("user not found"))

    err := suite.userUseCase.Promote("unknown@example.com")
    assert.Error(suite.T(), err)
    suite.mockUserRepo.AssertExpectations(suite.T())
}


func TestUserUseCaseTestSuite(t *testing.T) {
    suite.Run(t, new(UserUseCaseTestSuite))
}