package main

import (
	"awesomeProject6/config"
	"awesomeProject6/controller"
	"awesomeProject6/model"
	"awesomeProject6/service"
	"awesomeProject6/service/db"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	APIConfig, err := config.NewConfig()
	if err != nil {
		fmt.Println("Failed to load config: ", err)
	}

	db.ConnectDB(APIConfig)
	db.DB.AutoMigrate(&model.Movie{})

	timeService := service.NewTimeService()
	timeController := controller.NewController(timeService)

	movieService := service.NewMovieService()
	movieController := controller.NewMovieController(movieService)

	router := gin.Default()
	router.GET("/healthz", timeController.GetTime)
	router.POST("/movies", movieController.CreateMovie)
	//router.GET("/movies", movieController.GetMovieOffsetPaging)
	router.GET("/movies", movieController.GetMovieCursorPaging)
	router.GET("/movies/:id", movieController.GetMovieByID)
	router.GET("/movies/search", movieController.SearchMovie)
	//db.SeedMoviesFromCSV("service/db/movies.csv")
	router.Run(":8080")

}
