package routes

import (
	"github.com/gin-gonic/gin"
	"moviedb/internal/auth"
	"moviedb/internal/db"
	"moviedb/internal/delivery"
	"moviedb/internal/repository"
	"moviedb/internal/services"
)

func SetupRoutes(r *gin.Engine) {
	r.POST("/login", auth.Login)
	r.POST("/register", auth.Register)

	dbInstance := db.DB
	repo := repository.NewMovieRepository(dbInstance)
	service := services.NewMovieService(repo)
	handler := delivery.NewMovieHandler(service)

	movies := r.Group("api/v1/movies")
	{
		movies.GET("/", handler.GetAllMovies)
		movies.GET("/:id", handler.GetMovieByID)
		movies.POST("/", handler.CreateMovie)
		movies.PUT("/:id", handler.UpdateMovie)
		movies.DELETE("/:id", handler.DeleteMovie)
	}
}
