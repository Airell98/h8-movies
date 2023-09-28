package service

import (
	"h8-movies/dto"
	"h8-movies/entity"
	"h8-movies/pkg/errs"
	"h8-movies/repository/user_repository"
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
}
