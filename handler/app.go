package handler

import (
	"h8-movies/database"
	"h8-movies/docs"
	"h8-movies/repository/movie_repository/movie_pg"
	"h8-movies/repository/user_repository/user_pg"
	"h8-movies/service"
	"os"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
	// swagger embed files
)

func StartApp() {
	var port = os.Getenv("PORT")
	database.InitiliazeDatabase()

	db := database.GetDatabaseInstance()

	movieRepo := movie_pg.NewMoviePG(db)

	movieService := service.NewMovieService(movieRepo)

	movieHandler := NewMovieHandler(movieService)

	userRepo := user_pg.NewUserPG(db)

	userService := service.NewUserService(userRepo)

	userHandler := NewUserHandler(userService)

	authService := service.NewAuthService(userRepo, movieRepo)

	route := gin.Default()

	docs.SwaggerInfo.Title = "Belajar DDD"
	docs.SwaggerInfo.Description = "Ini adalah API dengan pattern DDD"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8080"
	docs.SwaggerInfo.Schemes = []string{"http"}

	route.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	movieRoute := route.Group("/movies")
	{

		movieRoute.Use(authService.Authentication())

		movieRoute.POST("/", movieHandler.CreateMovie)
		movieRoute.PUT("/:movieId", authService.Authorization(), movieHandler.UpdateMovieById)
	}

	userRoute := route.Group("/users")
	{
		userRoute.POST("/register", userHandler.Register)

		userRoute.POST("/login", userHandler.Login)
	}

	route.Run(":" + port)
}

// {  "planId": 1,
//   "identificationNumber": "3175082103981004",
//   "name": "Tsana",
//   "pob": "Jakarta",
//   "dob": "21-03-1992",
//   "phoneNumber": "089637750999",
//   "email": "tsana@gmail.com",
//   "postalCode": "12710",
//   "address": "Jakarta",
//   "gender": "male",
//   "identificationImageFileId": "9999",
//   }
