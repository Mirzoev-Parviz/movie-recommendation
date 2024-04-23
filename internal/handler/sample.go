package handler

import (
	"github.com/Mirzoev-Parviz/movie-recommendation/internal/dto"
	"github.com/Mirzoev-Parviz/movie-recommendation/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) GetRCM(c *gin.Context) {
	var input dto.RCMInput

	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ides := h.services.Recommend(dto.Interactions, dto.Items, dto.Users, input.UserID)
	rcms := utils.ConvertRCM(ides)

	c.JSON(http.StatusOK, rcms)
}
