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

type userServiceMocks struct{}

func NewUserServiceMocks() service.UserService {
	return &userServiceMocks{}
}

func (u *userServiceMocks) CreateNewUser(payload dto.NewUserRequest) (*dto.NewUserResponse, errs.MessageErr) {
	return CreateNewUser(payload)
}
func (u *userServiceMocks) Login(newUserRequest dto.NewUserRequest) (*dto.LoginResponse, errs.MessageErr) {
	return Login(newUserRequest)
}
