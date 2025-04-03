package services

import (
	"moviedb/internal/models"
)

type MovieRepository interface {
	GetAll() ([]models.Movie, error)
	GetById(id int) (*models.Movie, error)
	Create(movie *models.Movie) error
	Update(id int, movie *models.Movie) error
	Delete(movieID int) error
	CreateDirector(director *models.Director) error // Добавляем этот метод
}

type MovieService struct {
	repo MovieRepository
}

func NewMovieService(movieRepo MovieRepository) *MovieService {
	return &MovieService{repo: movieRepo}
}

func (s *MovieService) GetAllMovies() ([]models.Movie, error) {
	return s.repo.GetAll()
}

func (s *MovieService) GetMovieByID(id int) (*models.Movie, error) {
	return s.repo.GetById(id)
}

func (s *MovieService) Create(title, firstname, lastname string) (*models.Movie, error) {
	director := &models.Director{
		Firstname: firstname,
		Lastname:  lastname,
	}

	err := s.repo.CreateDirector(director)
	if err != nil {
		return nil, err
	}

	movie := &models.Movie{
		Title:      title,
		DirectorID: director.ID,
	}

	err = s.repo.Create(movie)
	return movie, err
}

func (s *MovieService) Update(id int, title, firstname, lastname string) (*models.Movie, error) {
	movie, err := s.repo.GetById(id)
	if err != nil {
		return nil, err
	}

	movie.Title = title

	if movie.Director != nil {
		movie.Director.Firstname = firstname
		movie.Director.Lastname = lastname
	} else {
		movie.Director = &models.Director{
			Firstname: firstname,
			Lastname:  lastname,
		}
	}

	err = s.repo.Update(id, movie)
	if err != nil {
		return nil, err
	}

	return s.GetMovieByID(id)
}

func (s *MovieService) DeleteMovie(movieID int) error {
	return s.repo.Delete(movieID)
}
