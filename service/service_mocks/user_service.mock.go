package service_mocks

import (
	"h8-movies/dto"
	"h8-movies/pkg/errs"
	"h8-movies/service"
)

var (
	CreateNewUser func(payload dto.NewUserRequest) (*dto.NewUserResponse, errs.MessageErr)
	Login         func(newUserRequest dto.NewUserRequest) (*dto.LoginResponse, errs.MessageErr)
)

type userServiceMock struct{}

func NewUserServiceMock() service.UserService {
	return &userServiceMock{}
}

func (u *userServiceMock) CreateNewUser(payload dto.NewUserRequest) (*dto.NewUserResponse, errs.MessageErr) {
	return CreateNewUser(payload)
}
func (u *userServiceMock) Login(newUserRequest dto.NewUserRequest) (*dto.LoginResponse, errs.MessageErr) {
	return Login(newUserRequest)
}
