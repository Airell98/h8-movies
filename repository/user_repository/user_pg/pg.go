package user_pg

import (
	"database/sql"
	"errors"
	"h8-movies/entity"
	"h8-movies/pkg/errs"
	"h8-movies/repository/user_repository"
)

const (
	retrieveUserByEmail = `
		SELECT id, email, password from users
		WHERE email = $1;
	`

	retrieveUserById = `
		SELECT id, email, password from users
		WHERE id = $1;
	`

	createNewUser = `
		INSERT INTO "users"
		(
			email,
			password
		)
		VALUES ($1, $2)
	`
)

type userPG struct {
	db *sql.DB
}

func NewUserPG(db *sql.DB) user_repository.Repository {
	return &userPG{
		db: db,
	}
}

func (u *userPG) CreateNewUser(user entity.User) errs.MessageErr {
	_, err := u.db.Exec(createNewUser, user.Email, user.Password)

	if err != nil {
		return errs.NewInternalServerError("something went wrong")
	}

	return nil
}

func (u *userPG) GetUserById(userId int) (*entity.User, errs.MessageErr) {
	var user entity.User

	row := u.db.QueryRow(retrieveUserById, userId)

	err := row.Scan(&user.Id, &user.Email, &user.Password)

	if err != nil {
		if errors.Is(sql.ErrNoRows, err) {
			return nil, errs.NewNotFoundError("user not found")
		}
		return nil, errs.NewInternalServerError("something went wrong")
	}

	return &user, nil
}

func (u *userPG) GetUserByEmail(userEmail string) (*entity.User, errs.MessageErr) {
	var user entity.User

	row := u.db.QueryRow(retrieveUserByEmail, userEmail)

	err := row.Scan(&user.Id, &user.Email, &user.Password)

	if err != nil {
		if errors.Is(sql.ErrNoRows, err) {
			return nil, errs.NewNotFoundError("user not found")
		}
		return nil, errs.NewInternalServerError("something went wrong")
	}

	return &user, nil
}
