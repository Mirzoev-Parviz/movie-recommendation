package handler

import (
	"net/http"

	"github.com/Mirzoev-Parviz/movie-recommendation/internal/services"
	"github.com/gin-gonic/gin"
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

	router.POST("/recommendations", h.GetRCM)
	router.POST("/watched-movies", h.GetWatched)
	return router
}
