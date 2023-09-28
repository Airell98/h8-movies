package handler

import (
	"h8-movies/docs"
	"h8-movies/infra/config"
	"h8-movies/infra/database"
	"h8-movies/repository/movie_repository/movie_pg"
	"h8-movies/repository/user_repository/user_pg"
	"h8-movies/service"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
	// swagger embed files
)

func StartApp() {
	config.LoadAppConfig()

	database.InitiliazeDatabase()

	var port = config.GetAppConfig().Port

	db := database.GetDatabaseInstance()

	movieRepo := movie_pg.NewMoviePG(db)

	movieService := service.NewMovieService(movieRepo)

	movieHandler := NewMovieHandler(movieService)

	userRepo := user_pg.NewUserPG(db)

	userService := service.NewUserService(userRepo)

	userHandler := NewUserHandler(userService)

	authService := service.NewAuthService(userRepo, movieRepo)

	docs.SwaggerInfo.Title = "Belajar DDD"
	docs.SwaggerInfo.Description = "Ini adalah API dengan pattern DDD"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "h8-movies-production.up.railway.app"
	docs.SwaggerInfo.Schemes = []string{"https"}

	route := gin.Default()

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
