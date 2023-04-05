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

type userMockRepository struct {
}

func NewUserMockRespository() UserRepository {
	return &userMockRepository{}
}

func (u *userMockRepository) CreateNewUser(user entity.User) errs.MessageErr {

	return CreateNewUser(user)
}

func (u *userMockRepository) GetUserById(userId int) (*entity.User, errs.MessageErr) {
	return GetUserById(userId)
}
func (u *userMockRepository) GetUserByEmail(userEmail string) (*entity.User, errs.MessageErr) {
	return GetUserByEmail(userEmail)
}
