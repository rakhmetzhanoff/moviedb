package repository

import (
	"gorm.io/gorm"
	"moviedb/internal/models"
)

type MovieRepositoryImpl struct {
	db *gorm.DB
}

func NewMovieRepository(db *gorm.DB) *MovieRepositoryImpl {
	return &MovieRepositoryImpl{db: db}
}

func (r *MovieRepositoryImpl) GetAll() ([]models.Movie, error) {
	var movies []models.Movie
	err := r.db.Preload("Director").Find(&movies).Error
	return movies, err
}

func (r *MovieRepositoryImpl) GetById(id int) (*models.Movie, error) {
	var movie models.Movie
	err := r.db.Preload("Director").First(&movie, id).Error
	return &movie, err
}

func (r *MovieRepositoryImpl) Create(movie *models.Movie) error {
	err := r.db.Create(movie).Error
	if err != nil {
		return err
	}

	err = r.db.Preload("Director").First(movie, movie.ID).Error
	return err
}

func (r *MovieRepositoryImpl) CreateDirector(director *models.Director) error {
	return r.db.Create(director).Error
}

func (r *MovieRepositoryImpl) Update(id int, movie *models.Movie) error {
	err := r.db.Model(&models.Movie{}).Where("id = ?", id).
		Omit("id", "CreatedAt").
		Updates(movie).Error
	if err != nil {
		return err
	}

	if movie.Director != nil {
		err = r.db.Model(&models.Director{}).Where("id = ?", movie.DirectorID).
			Updates(movie.Director).Error
		if err != nil {
			return err
		}
	}

	return nil
}

func (r *MovieRepositoryImpl) Delete(movieID int) error {
	return r.db.Delete(&models.Movie{}, movieID).Error
}
