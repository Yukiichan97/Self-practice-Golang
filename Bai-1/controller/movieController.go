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
