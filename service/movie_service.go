package service

import (
	"h8-movies/dto"
	"h8-movies/entity"
	"h8-movies/pkg/errs"
	"h8-movies/pkg/helpers"
	"h8-movies/repository/movie_repository"
	"net/http"
)

type MovieService interface {
	CreateMovie(userId int, payload dto.NewMovieRequest) (*dto.NewMovieResponse, errs.MessageErr)
	UpdateMovieById(movieId int, movieRequest dto.NewMovieRequest) (*dto.NewMovieResponse, errs.MessageErr)
}

type movieService struct {
	movieRepo movie_repository.Repository
}

func NewMovieService(movieRepo movie_repository.Repository) MovieService {
	return &movieService{
		movieRepo: movieRepo,
	}
}

func (m *movieService) UpdateMovieById(movieId int, movieRequest dto.NewMovieRequest) (*dto.NewMovieResponse, errs.MessageErr) {

	err := helpers.ValidateStruct(movieRequest)

	if err != nil {
		return nil, err
	}

	payload := entity.Movie{
		Id:       movieId,
		Title:    movieRequest.Title,
		ImageUrl: movieRequest.ImageUrl,
		Price:    movieRequest.Price,
	}

	err = m.movieRepo.UpdateMovieById(payload)

	if err != nil {
		return nil, err
	}

	response := dto.NewMovieResponse{
		StatusCode: http.StatusOK,
		Result:     "success",
		Message:    "movie data successfully updated",
	}

	return &response, nil
}

func (m *movieService) CreateMovie(userId int, payload dto.NewMovieRequest) (*dto.NewMovieResponse, errs.MessageErr) {
	err := helpers.ValidateStruct(payload)

	if err != nil {
		return nil, err
	}

	movieRequest := &entity.Movie{
		Title:    payload.Title,
		Price:    payload.Price,
		ImageUrl: payload.ImageUrl,
		UserId:   userId,
	}

	_, err = m.movieRepo.CreateMovie(movieRequest)

	if err != nil {
		return nil, err
	}

	response := dto.NewMovieResponse{
		StatusCode: http.StatusCreated,
		Result:     "success",
		Message:    "new movie data successfully created",
	}

	return &response, err
}
