package movie_repository

import (
	"h8-movies/entity"
	"h8-movies/pkg/errs"
)

type Repository interface {
	CreateMovie(moviePayload *entity.Movie) (*entity.Movie, errs.MessageErr)
	GetMovieById(movieId int) (*entity.Movie, errs.MessageErr)
	UpdateMovieById(payload entity.Movie) errs.MessageErr
}
