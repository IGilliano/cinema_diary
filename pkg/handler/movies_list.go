package handler

import (
	"cinema_diary"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) addToUserList(c *gin.Context) {
	id, _ := c.Get(userCtx)
	fmt.Println(id)

	var moviesList cinema_diary.MoviesList

	if err := c.BindJSON(&moviesList); err != nil {
		return
	}

	err := h.service.AddToUserList(moviesList)
	if err != nil {
		newErrorResponce(c, http.StatusBadRequest, err.Error())
	}

	c.JSON(http.StatusOK, "Film added to your list!")
}
