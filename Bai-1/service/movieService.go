package service

import (
	"awesomeProject6/model"
	"awesomeProject6/service/db"
)

type MovieService struct {
}

func NewMovieService() *MovieService {
	return &MovieService{}
}

func (s *MovieService) CreateMovie(movie *model.Movie) error {
	return db.DB.Create(movie).Error
}

func (s *MovieService) GetMovieByID(id uint) (*model.Movie, error) {
	var movie model.Movie
	err := db.DB.First(&movie, id).Error
	if err != nil {
		return nil, err
	}
	return &movie, nil
}

func (s *MovieService) SearchMovies(query string, year int) ([]model.Movie, error) {
	var movies []model.Movie
	data := db.DB

	if query != "" {
		data = data.Where("title ILIKE ? OR genre ILIKE ?",
			"%"+query+"%", "%"+query+"%")
	}

	if year > 0 {
		data = data.Where("year = ?", year)
	}

	err := data.Find(&movies).Error
	return movies, err
}
