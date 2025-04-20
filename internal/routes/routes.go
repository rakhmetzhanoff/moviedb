package routes

import (
	"github.com/gin-gonic/gin"
	"moviedb/internal/auth"
	"moviedb/internal/db"
	"moviedb/internal/delivery"
	"moviedb/internal/middleware"
	"moviedb/internal/repository"
	"moviedb/internal/services"
)

func SetupRoutes(r *gin.Engine) {
	api := r.Group("/api/v1")
	api.POST("/register", auth.Register)
	api.POST("/login", auth.Login)
	api.GET("/me", middleware.AuthRequired(), auth.Me)

	dbInstance := db.DB
	repo := repository.NewMovieRepository(dbInstance)
	service := services.NewMovieService(repo)
	handler := delivery.NewMovieHandler(service)

	movies := api.Group("/movies")
	movies.Use(middleware.AuthRequired())
	{
		movies.GET("/", handler.GetAllMovies)
		movies.GET("/:id", handler.GetMovieByID)
	}

	adminMovies := api.Group("/movies")
	adminMovies.Use(middleware.AuthRequired(), middleware.AdminOnly())
	{
		adminMovies.POST("/", handler.CreateMovie)
		adminMovies.PUT("/:id", handler.UpdateMovie)
		adminMovies.DELETE("/:id", handler.DeleteMovie)
	}
}
