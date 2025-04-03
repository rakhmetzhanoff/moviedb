package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"moviedb/internal/delivery"
	"moviedb/internal/repository"
	"moviedb/internal/services"
)

func SetupRoutes(r *gin.Engine, db *gorm.DB) {
	movieRepo := repository.NewMovieRepository(db)

	movieService := services.NewMovieService(movieRepo)

	movieHandler := delivery.NewMovieHandler(movieService)

	movies := r.Group("api/v1/movies")
	{
		movies.GET("/", movieHandler.GetAllMovies)
		movies.POST("/", movieHandler.CreateMovie)
		movies.GET("/:id", movieHandler.GetMovieByID)
		movies.PUT("/:id", movieHandler.UpdateMovie)
		movies.DELETE("/:id", movieHandler.DeleteMovie)
	}
}
