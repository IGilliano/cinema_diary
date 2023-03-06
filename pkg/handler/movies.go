package handler

import (
	_ "errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type movie struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Director string `json:"director"`
	Year     string `json:"year"`
}

func (h *Handler) getMovies(c *gin.Context) {
	id, _ := c.Get(userCtx)
	fmt.Println(id)

	movies, err := h.service.GetMovies()
	if err != nil {
		newErrorResponce(c, http.StatusBadRequest, err.Error())
	}

	c.JSON(http.StatusOK, movies)
}
