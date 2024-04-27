package handler

import (
	"net/http"

	"github.com/Mirzoev-Parviz/movie-recommendation/internal/dto"
	"github.com/Mirzoev-Parviz/movie-recommendation/utils"
	"github.com/gin-gonic/gin"
)

func (h *Handler) GetRCM(c *gin.Context) {
	var input dto.RCMInput

	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if input.UserID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "empty request"})
		return
	}

	ides := h.services.Recommend(dto.Interactions, dto.Items, dto.Users, input.UserID)
	rcms := utils.ConvertRCM(ides)

	c.JSON(http.StatusOK, rcms)
}

func (h *Handler) GetWatched(c *gin.Context) {
	var input dto.RCMInput

	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	items := h.services.GetUserWatchedMovies(input.UserID)

	c.JSON(http.StatusOK, items)
}
