package user_repository

import (
	"h8-movies/entity"
	"h8-movies/pkg/errs"
)

type UserRepository interface {
	CreateNewUser(user entity.User) errs.MessageErr
	GetUserById(userId int) (*entity.User, errs.MessageErr)
	GetUserByEmail(userEmail string) (*entity.User, errs.MessageErr)
}
