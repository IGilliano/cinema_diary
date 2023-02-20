package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) getMovies(c *gin.Context) {
	id, _ := c.Get(userCtx)
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}