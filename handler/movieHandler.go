package handler

import (
	"fmt"
	"h8-movies/dto"
	"h8-movies/entity"
	"h8-movies/pkg/errs"
	"h8-movies/pkg/helpers"
	"h8-movies/service"
	"net/http"

	_ "h8-movies/entity"

	"github.com/gin-gonic/gin"
)

type movieHandler struct {
	movieService service.MovieService
}

func NewMovieHandler(movieService service.MovieService) movieHandler {
	return movieHandler{
		movieService: movieService,
	}
}

// CreateNewMovie godoc
// @Tags movies
// @Description Create New Movie Data
// @ID create-new-movie
// @Accept json
// @Produce json
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Param RequestBody body dto.NewMovieRequest true "request body json"
// @Success 201 {object} dto.NewMovieRequest
// @Router /movies [post]
func (m movieHandler) CreateMovie(ctx *gin.Context) {
	var movieRequest dto.NewMovieRequest

	if err := ctx.ShouldBindJSON(&movieRequest); err != nil {
		errBindJson := errs.NewUnprocessibleEntityError("invalid request body")

		ctx.JSON(errBindJson.Status(), errBindJson)
		return
	}

	user := ctx.MustGet("userData").(entity.User)

	fmt.Printf("user must get: %+v\n", user)

	newMovie, err := m.movieService.CreateMovie(user.Id, movieRequest)

	if err != nil {
		ctx.JSON(err.Status(), err)
		return
	}

	ctx.JSON(http.StatusCreated, newMovie)
}

func (m movieHandler) UpdateMovieById(c *gin.Context) {
	var movieRequest dto.NewMovieRequest

	if err := c.ShouldBindJSON(&movieRequest); err != nil {
		errBindJson := errs.NewUnprocessibleEntityError("invalid request body")

		c.JSON(errBindJson.Status(), errBindJson)
		return
	}

	movieId, err := helpers.GetParamId(c, "movieId")

	if err != nil {
		c.AbortWithStatusJSON(err.Status(), err)
		return
	}

	response, err := m.movieService.UpdateMovieById(movieId, movieRequest)

	if err != nil {
		c.AbortWithStatusJSON(err.Status(), err)
		return
	}

	c.JSON(response.StatusCode, response)
}
