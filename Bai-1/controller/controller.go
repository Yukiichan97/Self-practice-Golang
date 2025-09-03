package controller

import (
	"awesomeProject6/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	service *service.Service
}

func NewController(Service *service.Service) *Controller {
	return &Controller{
		service: Service,
	}
}

func (c *Controller) GetTime(ctx *gin.Context) {
	result := c.service.GetTime()
	ctx.JSON(http.StatusOK, result)
}
