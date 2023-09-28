package user_repository

import (
	"h8-movies/entity"
	"h8-movies/pkg/errs"
)

var (
	CreateNewUser  func(user entity.User) errs.MessageErr
	GetUserById    func(userId int) (*entity.User, errs.MessageErr)
	GetUserByEmail func(userEmail string) (*entity.User, errs.MessageErr)
)

type repositoryMock struct {
}

func NewRepositoryMock() Repository {
	return &repositoryMock{}
}

func (rm *repositoryMock) CreateNewUser(user entity.User) errs.MessageErr {
	return CreateNewUser(user)
}

func (rm *repositoryMock) GetUserById(userId int) (*entity.User, errs.MessageErr) {
	return GetUserById(userId)
}

func (rm *repositoryMock) GetUserByEmail(userEmail string) (*entity.User, errs.MessageErr) {
	return GetUserByEmail(userEmail)
}
