package handler

import (
	"cinema_diary"
	_ "errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) getMovies(c *gin.Context) {

	movies, err := h.service.GetMovies()
	if err != nil {
		newErrorResponce(c, http.StatusBadRequest, err.Error())
	}

	c.JSON(http.StatusOK, movies)
}

func (h *Handler) getMovie(c *gin.Context) {

	movieId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponce(c, http.StatusBadRequest, "Invalid ID")
		return
	}

	movie, err := h.service.GetMovie(movieId)
	if err != nil {
		newErrorResponce(c, http.StatusNotFound, err.Error())
	}

	c.JSON(http.StatusOK, movie)

}

func (h *Handler) addMovies(c *gin.Context) {

	userId, err := getUserId(c)
	if err != nil {
		newErrorResponce(c, http.StatusInternalServerError, err.Error())
	}

	if userId != 1 {
		newErrorResponce(c, http.StatusForbidden, "You need permition to add movies")
	}

	var movies []*cinema_diary.Movie

	if err = c.BindJSON(&movies); err != nil {
		newErrorResponce(c, http.StatusBadRequest, "Data is incorrect")
	}

	fmt.Println(movies)
	moviesId, err := h.service.AddMovies(movies)

	if err != nil {
		newErrorResponce(c, http.StatusBadRequest, err.Error())
	}

	c.JSON(http.StatusOK, moviesId)

}

func (h *Handler) deleteMovie(c *gin.Context) {

	userId, err := getUserId(c)
	if err != nil {
		newErrorResponce(c, http.StatusInternalServerError, err.Error())
	}

	if userId != 1 {
		newErrorResponce(c, http.StatusForbidden, "You need permition to add movies")
	}

	movieId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponce(c, http.StatusBadRequest, "Invalid ID")
		return
	}

	err = h.service.DeleteMovie(movieId)
	if err != nil {
		newErrorResponce(c, http.StatusBadRequest, err.Error())
	}

	c.JSON(http.StatusOK, "Movie was deleted from list")
}
