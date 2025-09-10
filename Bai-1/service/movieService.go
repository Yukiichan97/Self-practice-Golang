package service

import (
	"awesomeProject6/model"
	"awesomeProject6/service/db"
	"encoding/base64"
	"encoding/json"
	"fmt"

	"gorm.io/gorm"
)

type MovieService struct {
}

type CursorData struct {
	ID   uint `json:"id"`
	Year int  `json:"year,omitempty"`
}

type Cursor struct {
	Data       []model.Movie `json:"data"`
	Year       int           `json:"year,omitempty"`
	NextCursor string        `json:"next_cursor,omitempty"`
	PrevCursor string        `json:"prev_cursor,omitempty"`
	HasNext    bool          `json:"has_next"`
	HasPrev    bool          `json:"has_prev"`
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

func (s *MovieService) GetMovieOffsetPaging(page int, size int) ([]model.Movie, error) {
	data := db.DB
	offset := (page - 1) * size
	var movies []model.Movie
	data = data.Offset(offset).Limit(size)
	err := data.Find(&movies).Error
	return movies, err
}

func (s *MovieService) GetMovieCursorPaging(cursorStr string, size int, year int) (*Cursor, error) {
	var cursorData *CursorData
	var err error
	data := db.DB
	if cursorStr != "" {
		fmt.Println(cursorStr)
		cursorData, err = s.decodeCursor(cursorStr)
		if err != nil {
			return nil, fmt.Errorf("invalid cursor: %v", err)
		}
	}

	var movies []model.Movie
	query := data.Where("year = ?", year).Find(&model.Movie{})

	movies, err = s.getCursorPageByIdYear(query, cursorData, size)

	if err != nil {
		return nil, err
	}

	// Tạo response với cursors
	response := &Cursor{
		Data:    movies,
		Year:    year,
		HasNext: len(movies) == size, // Có thể có more data
	}

	// Tạo next cursor
	if len(movies) > 0 && response.HasNext {
		lastMovie := movies[len(movies)-1]
		nextCursor := &CursorData{
			ID:   lastMovie.ID,
			Year: year,
		}
		response.NextCursor, _ = s.encodeCursor(nextCursor)
		nextCursor.Year = lastMovie.Year
	}

	// Tạo prev cursor
	if len(movies) > 0 && cursorData != nil {
		firstMovie := movies[0]
		prevCursor := &CursorData{
			ID:   firstMovie.ID,
			Year: year,
		}

		prevCursor.Year = firstMovie.Year
		response.PrevCursor, _ = s.encodeCursor(prevCursor)
		response.HasPrev = true
	}

	return response, nil
}

func (s *MovieService) encodeCursor(cursor *CursorData) (string, error) {
	jsonBytes, err := json.Marshal(cursor)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(jsonBytes), nil
}

func (s *MovieService) decodeCursor(cursorStr string) (*CursorData, error) {
	jsonBytes, err := base64.StdEncoding.DecodeString(cursorStr)
	if err != nil {
		return nil, err
	}

	var cursor CursorData
	err = json.Unmarshal(jsonBytes, &cursor)
	if err != nil {
		return nil, err
	}
	fmt.Println(&cursor)
	return &cursor, nil
}

func (s *MovieService) getCursorPageByIdYear(query *gorm.DB, cursor *CursorData, size int) ([]model.Movie, error) {
	var movies []model.Movie

	if cursor != nil {
		query = query.Where("year = ? AND id > ?",
			cursor.Year, cursor.ID)
	}

	err := query.Order("year ASC, id ASC").Limit(size + 1).Find(&movies).Error

	if len(movies) > size {
		movies = movies[:size]
	}

	return movies, err
}
