package movie_pg

import (
	"h8-movies/entity"
	"net/http"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMoviePG_GetMovieById_Success(t *testing.T) {
	db, mock, err := sqlmock.New()

	require.Nil(t, err)

	defer db.Close()

	currentTime := time.Now()

	moviePG := NewMoviePG(db)

	query := regexp.QuoteMeta(getMovieByIdQuery)

	_, _ = query, mock

	movie := &entity.Movie{
		Id:        1,
		Title:     "Movie Test",
		Price:     4000,
		UserId:    1,
		ImageUrl:  "http.com",
		CreatedAt: currentTime,
		UpdatedAt: currentTime,
	}

	_ = movie

	row := sqlmock.NewRows([]string{"id", "title", "userId", "imageUrl", "price", "createdAt", "updatedAt"})

	row.AddRow(movie.Id, movie.Title, movie.UserId, movie.ImageUrl, movie.Price, movie.CreatedAt, movie.UpdatedAt)

	mock.ExpectQuery(query).WithArgs(200).WillReturnRows(row)

	result, err := moviePG.GetMovieById(200)

	assert.Nil(t, err)

	assert.NotNil(t, result)
}

func TestMoviePG_GetMovieById_NotFoundError(t *testing.T) {
	db, mock, errMock := sqlmock.New()

	require.Nil(t, errMock)

	defer db.Close()

	query := regexp.QuoteMeta(getMovieByIdQuery)

	row := sqlmock.NewRows([]string{"id", "title", "userId", "imageUrl", "price", "createdAt", "updatedAt"})

	moviePG := NewMoviePG(db)

	mock.ExpectQuery(query).WithArgs(1).WillReturnRows(row)

	result, err := moviePG.GetMovieById(1)

	require.Nil(t, result)

	require.NotNil(t, err)

	assert.Equal(t, http.StatusNotFound, err.Status())

	assert.Equal(t, "movie not found", err.Message())

	assert.Equal(t, "NOT_FOUND", err.Error())
}

func TestMoviePG_GetMovieById_InternalServerError(t *testing.T) {
	db, mock, errMock := sqlmock.New()

	require.Nil(t, errMock)

	moviePG := NewMoviePG(db)

	query := regexp.QuoteMeta(getMovieByIdQuery)

	currentTime := time.Now()

	movie := &entity.Movie{
		Id:        1,
		Title:     "Movie Test",
		Price:     4000,
		UserId:    1,
		ImageUrl:  "http.com",
		CreatedAt: currentTime,
		UpdatedAt: currentTime,
	}

	row := sqlmock.NewRows([]string{"id", "title", "userId", "imageUrl", "price", "createdAt", "updatedAt"})

	row.AddRow(movie.UpdatedAt, movie.Id, movie.Title, movie.UserId, movie.ImageUrl, movie.Price, movie.CreatedAt)

	mock.ExpectQuery(query).WithArgs(200).WillReturnRows(row)

	result, err := moviePG.GetMovieById(200)

	assert.Nil(t, result)

	assert.NotNil(t, err)
}
