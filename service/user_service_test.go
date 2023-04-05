package service

import (
	"h8-movies/dto"
	"h8-movies/entity"
	"h8-movies/pkg/errs"
	"h8-movies/repository/user_repository"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserService_CreateNewUser_Success(t *testing.T) {
	payload := dto.NewUserRequest{
		Email:    "test@gmail.com",
		Password: "123456",
	}

	userRepo := user_repository.NewUserMockRespository()

	userService := NewUserService(userRepo)

	user_repository.CreateNewUser = func(user entity.User) errs.MessageErr {
		return nil
	}

	result, err := userService.CreateNewUser(payload)

	assert.Nil(t, err)

	assert.NotNil(t, result)

	assert.Equal(t, "success", result.Result)

	assert.Equal(t, http.StatusCreated, result.StatusCode)
}

func TestUserService_CreateNewUser_BadRequestError(t *testing.T) {
	userService := NewUserService(nil)

	tests := []struct {
		name        string
		expectation errs.MessageErr
		payload     dto.NewUserRequest
	}{
		{
			name:        "should return email cannot be empty error",
			expectation: errs.NewBadRequest("email cannot be empty"),
			payload: dto.NewUserRequest{
				Email:    "",
				Password: "123456",
			},
		},
		{
			name:        "should return password cannot be empty error",
			expectation: errs.NewBadRequest("password cannot be empty"),
			payload: dto.NewUserRequest{
				Email:    "test@gmail.com",
				Password: "",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := userService.CreateNewUser(tt.payload)

			assert.Nil(t, result)

			assert.NotNil(t, err)

			assert.Equal(t, tt.expectation.Status(), err.Status())

			assert.Equal(t, tt.expectation.Message(), err.Message())

			assert.Equal(t, tt.expectation.Error(), err.Error())
		})
	}

}

func TestUserService_CreateNewUser_HashPasswordError(t *testing.T) {
	userService := NewUserService(nil)

	payload := dto.NewUserRequest{
		Email:    "test@gmail.com",
		Password: "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
	}

	result, err := userService.CreateNewUser(payload)

	assert.Nil(t, result)

	assert.NotNil(t, err)

	assert.Equal(t, "something went wrong", err.Message())

	assert.Equal(t, http.StatusInternalServerError, err.Status())

	assert.Equal(t, "INTERNAL_SERVER_ERROR", err.Error())
}

func TestUserService_CreateNewUser_DbError(t *testing.T) {
	userRepo := user_repository.NewUserMockRespository()

	userService := NewUserService(userRepo)

	user_repository.CreateNewUser = func(user entity.User) errs.MessageErr {

		return errs.NewInternalServerError("something went wrong")
	}

	payload := dto.NewUserRequest{
		Email:    "test@gmail.com",
		Password: "123456",
	}

	result, err := userService.CreateNewUser(payload)

	assert.NotNil(t, err)

	assert.Nil(t, result)

	assert.Equal(t, "something went wrong", err.Message())

	assert.Equal(t, http.StatusInternalServerError, err.Status())

	assert.Equal(t, "INTERNAL_SERVER_ERROR", err.Error())

}
