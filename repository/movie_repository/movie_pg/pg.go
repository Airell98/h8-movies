package movie_pg

import (
	"database/sql"
	"errors"
	"fmt"
	"h8-movies/entity"
	"h8-movies/pkg/errs"
	"h8-movies/repository/movie_repository"
)

const (
	getMovieByIdQuery = `
		SELECT id, title, userId, imageUrl, price, createdAt, updatedAt from "movies"
		WHERE id = $1;
	`

	updateMovieByIdQuery = `
		UPDATE "movies"
		SET title = $2,
		imageUrl = $3,
		price = $4
		WHERE id = $1;
	`
)

type moviePG struct {
	db *sql.DB
}

func NewMoviePG(db *sql.DB) movie_repository.Repository {
	return &moviePG{
		db: db,
	}
}

func (m *moviePG) UpdateMovieById(payload entity.Movie) errs.MessageErr {
	_, err := m.db.Exec(updateMovieByIdQuery, payload.Id, payload.Title, payload.ImageUrl, payload.Price)

	if err != nil {

		return errs.NewInternalServerError("something went wrong")
	}

	return nil
}

func (m *moviePG) GetMovieById(movieId int) (*entity.Movie, errs.MessageErr) {
	row := m.db.QueryRow(getMovieByIdQuery, movieId)

	var movie entity.Movie

	err := row.Scan(&movie.Id, &movie.Title, &movie.UserId, &movie.ImageUrl, &movie.Price, &movie.CreatedAt, &movie.UpdatedAt)

	if err != nil {
		if errors.Is(sql.ErrNoRows, err) {
			return nil, errs.NewNotFoundError("movie not found")
		}

		return nil, errs.NewInternalServerError("something went wrong")
	}

	return &movie, nil
}

func (m *moviePG) CreateMovie(moviePayload *entity.Movie) (*entity.Movie, errs.MessageErr) {
	createMovieQuery := `
		INSERT INTO "movies"
		(
			title,
			imageUrl,
			price,
			userId
		)
		VALUES($1, $2, $3, $4)
		RETURNING id,title, imageUrl, price, userId;
	`
	row := m.db.QueryRow(createMovieQuery, moviePayload.Title, moviePayload.ImageUrl, moviePayload.Price, moviePayload.UserId)

	var movie entity.Movie

	err := row.Scan(&movie.Id, &movie.Title, &movie.ImageUrl, &movie.Price, &movie.UserId)

	if err != nil {
		fmt.Printf("err: %v\n", err)
		return nil, errs.NewInternalServerError("something went wrong")
	}

	return &movie, nil

}
