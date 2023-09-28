package service

import (
	"h8-movies/dto"
	"h8-movies/entity"
	"h8-movies/pkg/errs"
	"h8-movies/repository/user_repository"
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserService_CreateNewUser_Success(t *testing.T) {
	payload := dto.NewUserRequest{
		Email:    "john@mail.com",
		Password: "123456",
	}

	userRepo := user_repository.NewRepositoryMock()

	userService := NewUserService(userRepo)

	user_repository.CreateNewUser = func(user entity.User) errs.MessageErr {
		return nil
	}

	result, err := userService.CreateNewUser(payload)

	assert.Nil(t, err)

	assert.NotNil(t, result)

	assert.Equal(t, http.StatusCreated, result.StatusCode)
}

func TestUserService_CreateNewUser_InvalidPayload(t *testing.T) {
	payload := dto.NewUserRequest{
		Email:    "",
		Password: "123456",
	}

	userRepo := user_repository.NewRepositoryMock()

	userService := NewUserService(userRepo)

	result, err := userService.CreateNewUser(payload)

	assert.NotNil(t, err)

	assert.Nil(t, result)

	assert.Equal(t, "has to be a valid email", err.Message())
}

func TestUserService_CreateNewUser_HashPasswordError(t *testing.T) {
	payload := dto.NewUserRequest{
		Email:    "john@mail.com",
		Password: strings.Repeat("w", 73),
	}

	userRepo := user_repository.NewRepositoryMock()

	userService := NewUserService(userRepo)

	result, err := userService.CreateNewUser(payload)

	assert.NotNil(t, err)

	assert.Nil(t, result)

	assert.Equal(t, http.StatusInternalServerError, err.Status())

}

func TestUserService_CreateNewUser_DBError(t *testing.T) {
	payload := dto.NewUserRequest{
		Email:    "john@mail.com",
		Password: "123456",
	}

	userRepo := user_repository.NewRepositoryMock()

	userService := NewUserService(userRepo)

	user_repository.CreateNewUser = func(user entity.User) errs.MessageErr {
		return errs.NewInternalServerError("something went wrong")
	}

	result, err := userService.CreateNewUser(payload)

	assert.NotNil(t, err)

	assert.Nil(t, result)

	assert.Equal(t, http.StatusInternalServerError, err.Status())

	assert.Equal(t, "something went wrong", err.Message())

}
