package handler

import (
	"cinema_diary"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) addToUserList(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		newErrorResponce(c, http.StatusInternalServerError, err.Error())
	}

	var moviesList cinema_diary.MoviesList

	if err := c.BindJSON(&moviesList); err != nil {
		return
	}

	moviesList.UserId = userId

	err = h.service.AddToUserMoviesList(isWatched(moviesList))
	if err != nil {
		newErrorResponce(c, http.StatusBadRequest, err.Error())
	}

	c.JSON(http.StatusOK, "Film added to your list!")
}

func (h *Handler) getUserMoviesList(c *gin.Context) {
	var watched bool
	if c.FullPath() == "/api/movies-list/" {
		watched = true
	}
	fmt.Println(watched)
	userId, err := getUserId(c)
	if err != nil {
		newErrorResponce(c, http.StatusInternalServerError, err.Error())
	}

	moviesList, err := h.service.GetUserMoviesList(userId, watched)
	if err != nil {
		newErrorResponce(c, http.StatusBadRequest, err.Error())
	}
	c.JSON(http.StatusOK, moviesList)
}
