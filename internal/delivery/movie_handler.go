package delivery

import (
	"github.com/gin-gonic/gin"
	"moviedb/internal/models"
	"moviedb/internal/services"
	"net/http"
	"strconv"
)

func NewMovieHandler(service *services.MovieService) *MovieHandler {
	return &MovieHandler{service: service}
}

type MovieHandler struct {
	service *services.MovieService
}

func (h *MovieHandler) GetAllMovies(c *gin.Context) {
	movies, _ := h.service.GetAllMovies()
	c.JSON(http.StatusOK, movies)
}

func (h *MovieHandler) GetMovieByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid movie ID"})
		return
	}

	movie, err := h.service.GetMovieByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Movie not found"})
		return
	}

	c.JSON(http.StatusOK, movie)
}

func (h *MovieHandler) CreateMovie(c *gin.Context) {
	var movieCreate models.Movie

	if err := c.ShouldBindJSON(&movieCreate); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	newMovie, err := h.service.Create(
		movieCreate.Title,
		movieCreate.Director.Firstname,
		movieCreate.Director.Lastname,
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create movie"})
		return
	}

	c.JSON(http.StatusCreated, newMovie)
}

func (h *MovieHandler) UpdateMovie(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid movie ID"})
		return
	}

	var movieUpdate models.Movie
	if err := c.ShouldBindJSON(&movieUpdate); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	updatedMovie, err := h.service.Update(
		id,
		movieUpdate.Title,
		movieUpdate.Director.Firstname,
		movieUpdate.Director.Lastname,
	)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Movie not found"})
		return
	}

	if movieUpdate.Director == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Director is required"})
		return
	}
	c.JSON(http.StatusOK, updatedMovie)
}

func (h *MovieHandler) DeleteMovie(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid movie ID"})
		return
	}

	if err := h.service.DeleteMovie(id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Movie not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Movie deleted successfully"})
}
