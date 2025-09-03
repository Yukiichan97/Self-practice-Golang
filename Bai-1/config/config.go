package config

import (
	"github.com/gin-gonic/gin"
)

func NewConfig() *gin.Engine {
	router := gin.Default()
	return router
}
