package service

import (
	"h8-movies/dto"
	"h8-movies/entity"
	"h8-movies/pkg/errs"
	"h8-movies/pkg/helpers"
	"h8-movies/repository/user_repository"
	"net/http"
)

type UserService interface {
	CreateNewUser(payload dto.NewUserRequest) (*dto.NewUserResponse, errs.MessageErr)
	Login(newUserRequest dto.NewUserRequest) (*dto.LoginResponse, errs.MessageErr)
}

type userService struct {
	userRepo user_repository.Repository
}

func NewUserService(userRepo user_repository.Repository) UserService {
	return &userService{
		userRepo: userRepo,
	}
}

func (u *userService) Login(newUserRequest dto.NewUserRequest) (*dto.LoginResponse, errs.MessageErr) {
	err := helpers.ValidateStruct(newUserRequest)

	if err != nil {
		return nil, err
	}

	user, err := u.userRepo.GetUserByEmail(newUserRequest.Email)

	if err != nil {
		if err.Status() == http.StatusNotFound {
			return nil, errs.NewBadRequest("invalid email/password")
		}
		return nil, err
	}

	isValidPassword := user.ComparePassword(newUserRequest.Password)

	if !isValidPassword {
		return nil, errs.NewBadRequest("invalid email/password")
	}

	token := user.GenerateToken()

	response := dto.LoginResponse{
		Result:     "success",
		StatusCode: http.StatusOK,
		Message:    "successfully logged in",
		Data: dto.TokenResponse{
			Token: token,
		},
	}

	return &response, nil
}

func (u *userService) CreateNewUser(payload dto.NewUserRequest) (*dto.NewUserResponse, errs.MessageErr) {
	err := helpers.ValidateStruct(payload)

	if err != nil {
		return nil, err
	}

	user := entity.User{
		Email:    payload.Email,
		Password: payload.Password,
	}

	err = user.HashPassword()

	if err != nil {
		return nil, err
	}

	err = u.userRepo.CreateNewUser(user)

	if err != nil {
		return nil, err
	}

	response := dto.NewUserResponse{
		Result:     "success",
		StatusCode: http.StatusCreated,
		Message:    "user registered successfully",
	}

	return &response, nil
}
