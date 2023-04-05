package handler

import (
	"bytes"
	"encoding/json"
	"h8-movies/dto"
	"h8-movies/pkg/errs"
	"h8-movies/service/service_mocks"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUserHandler_Register_Success(t *testing.T) {
	userService := service_mocks.NewUserServiceMocks()

	userHandler := NewUserHandler(userService)

	payload := dto.NewUserRequest{
		Email:    "test@gmail.com",
		Password: "123456",
	}

	service_mocks.CreateNewUser = func(payload dto.NewUserRequest) (*dto.NewUserResponse, errs.MessageErr) {
		result := dto.NewUserResponse{
			Result:     "success",
			Message:    "user registered successfully",
			StatusCode: http.StatusCreated,
		}

		return &result, nil
	}

	jsonByte, err := json.Marshal(payload)

	require.Nil(t, err)

	req, err := http.NewRequest(http.MethodPost, "/users/register", bytes.NewBuffer(jsonByte))

	rr := httptest.NewRecorder()

	gin.SetMode(gin.TestMode)

	route := gin.Default()

	route.POST("/users/register", userHandler.Register)

	route.ServeHTTP(rr, req)

	result := rr.Result()

	responseBody, err := ioutil.ReadAll(result.Body)

	require.Nil(t, err)

	defer result.Body.Close()

	var userResponse dto.NewUserResponse

	err = json.Unmarshal(responseBody, &userResponse)

	require.Nil(t, err)

	assert.Equal(t, http.StatusCreated, userResponse.StatusCode)

	assert.Equal(t, "user registered successfully", userResponse.Message)

	assert.Equal(t, "success", userResponse.Result)
}
