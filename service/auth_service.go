package service

import (
	"fmt"
	"h8-movies/entity"
	"h8-movies/pkg/errs"
	"h8-movies/pkg/helpers"
	"h8-movies/repository/movie_repository"
	"h8-movies/repository/user_repository"

	"github.com/gin-gonic/gin"
)

type AuthService interface {
	Authentication() gin.HandlerFunc
	Authorization() gin.HandlerFunc
}

type authService struct {
	userRepo  user_repository.Repository
	movieRepo movie_repository.Repository
}

func NewAuthService(userRepo user_repository.Repository, movieRepo movie_repository.Repository) AuthService {
	return &authService{
		userRepo:  userRepo,
		movieRepo: movieRepo,
	}
}

func (a *authService) Authorization() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		user := ctx.MustGet("userData").(entity.User)

		movieId, err := helpers.GetParamId(ctx, "movieId") // 7

		if err != nil {
			fmt.Printf("[Authorization]: %s\n", err.Error())
			ctx.AbortWithStatusJSON(err.Status(), err)
			return
		}

		movie, err := a.movieRepo.GetMovieById(movieId)

		if err != nil {
			fmt.Printf("[Authorization]: %s\n", err.Error())
			ctx.AbortWithStatusJSON(err.Status(), err)
			return
		}

		if movie.UserId != user.Id {
			unauthorizedErr := errs.NewUnauthorizedError("you are not authorized to modify the movie data")
			ctx.AbortWithStatusJSON(unauthorizedErr.Status(), unauthorizedErr)
			return
		}

		ctx.Next()

	}
}

func (a *authService) Authentication() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var invalidTokenErr = errs.NewUnauthenticatedError("invalid token")
		bearerToken := ctx.GetHeader("Authorization")

		var user entity.User

		err := user.ValidateToken(bearerToken)

		if err != nil {
			fmt.Printf("[Authentication]: %s\n", err.Error())
			ctx.AbortWithStatusJSON(err.Status(), err)
			return
		}

		result, err := a.userRepo.GetUserByEmail(user.Email)

		if err != nil {
			ctx.AbortWithStatusJSON(invalidTokenErr.Status(), invalidTokenErr)
			return
		}

		_ = result

		ctx.Set("userData", user)

		ctx.Next()
	}
}
