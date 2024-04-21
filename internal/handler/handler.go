package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"recommendation/internal/services"
)

type Handler struct {
	services *services.Service
}

func NewHandler(service *services.Service) *Handler {
	return &Handler{services: service}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "working",
		})
	})

	return router
}
