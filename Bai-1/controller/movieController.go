package controller

import (
	"awesomeProject6/model"
	"awesomeProject6/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type MovieController struct {
	movieService *service.MovieService
}

func NewMovieController(service *service.MovieService) *MovieController {
	return &MovieController{movieService: service}
}

func (c MovieController) CreateMovie(ctx *gin.Context) {
	var movie model.Movie
	if err := ctx.ShouldBindBodyWithJSON(&movie); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	if err := c.movieService.CreateMovie(&movie); err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
	}

	ctx.JSON(http.StatusCreated, movie)

}

func (c MovieController) GetMovieByID(ctx *gin.Context) {
	idstr := ctx.Param("id")
	id, err := strconv.ParseUint(idstr, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	movie, err := c.movieService.GetMovieByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, movie)
}

func (c MovieController) SearchMovie(ctx *gin.Context) {
	query := ctx.Query("q")
	yearStr := ctx.Query("year")

	year := 0

	if yearStr != "" {
		if y, err := strconv.Atoi(yearStr); err == nil {
			year = y
		}
	}

	movies, err := c.movieService.SearchMovies(query, year)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, movies)
}

func (c MovieController) GetMovieOffsetPaging(ctx *gin.Context) {
	page, _ := strconv.Atoi(ctx.Query("page"))
	size, _ := strconv.Atoi(ctx.Query("size"))
	if page <= 0 {
		page = 1
	}
	if size <= 0 {
		size = 1
	}

	movies, err := c.movieService.GetMovieOffsetPaging(page, size)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, movies)
}

func (c MovieController) GetMovieCursorPaging(ctx *gin.Context) {
	cursor := ctx.Query("sort")
	sizeStr := ctx.Query("size")
	yearStr := ctx.Query("year")

	size := 10
	if sizeStr != "" {
		if l, err := strconv.Atoi(sizeStr); err == nil && l > 0 {
			size = l
		}
	}

	year, err := strconv.Atoi(yearStr)
	if err != nil {
		year = 0
	}

	result, err := c.movieService.GetMovieCursorPaging(cursor, size, year)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": "success",
		"pagination": gin.H{
			"year":        result.Year,
			"next_cursor": result.NextCursor,
			"prev_cursor": result.PrevCursor,
			"has_next":    result.HasNext,
			"has_prev":    result.HasPrev,
			"count":       len(result.Data),
		},
		"data": result.Data,
	})
}
