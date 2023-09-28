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

func TestUserHandler_CreateUser_Success(t *testing.T) {
	payload := dto.NewUserRequest{
		Email:    "john@mail.com",
		Password: "123456",
	}

	requestBody, err := json.Marshal(payload)

	userService := service_mocks.NewUserServiceMock()

	userHandler := NewUserHandler(userService)

	service_mocks.CreateNewUser = func(payload dto.NewUserRequest) (*dto.NewUserResponse, errs.MessageErr) {
		return &dto.NewUserResponse{
			Result:     "success",
			StatusCode: http.StatusCreated,
			Message:    "new user successfully registered",
		}, nil
	}

	require.Nil(t, err)

	req, err := http.NewRequest(http.MethodPost, "/users/register", bytes.NewBuffer(requestBody))

	require.Nil(t, err)

	rr := httptest.NewRecorder()

	gin.SetMode(gin.TestMode)
	route := gin.Default()

	route.POST("/users/register", userHandler.Register)

	route.ServeHTTP(rr, req)

	result := rr.Result()

	responseByte, err := ioutil.ReadAll(result.Body)

	require.Nil(t, err)

	defer result.Body.Close()

	var response dto.NewUserResponse

	err = json.Unmarshal(responseByte, &response)
	require.Nil(t, err)

	assert.Equal(t, http.StatusCreated, response.StatusCode)

}
