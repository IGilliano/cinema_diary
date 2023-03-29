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
		return
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
		return
	}

	c.JSON(http.StatusOK, movie)

}

func (h *Handler) addMovies(c *gin.Context) {

	userId, err := getUserId(c)
	if err != nil {
		newErrorResponce(c, http.StatusInternalServerError, err.Error())
		return
	}

	if userId != 1 {
		newErrorResponce(c, http.StatusForbidden, "You need permition to add movies")
		return
	}

	var movies []*cinema_diary.Movie

	if err = c.BindJSON(&movies); err != nil {
		newErrorResponce(c, http.StatusBadRequest, "Data is incorrect")
		return
	}

	fmt.Println(movies)
	moviesId, err := h.service.AddMovies(movies)

	if err != nil {
		newErrorResponce(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, moviesId)

}

func (h *Handler) deleteMovie(c *gin.Context) {

	userId, err := getUserId(c)
	if err != nil {
		newErrorResponce(c, http.StatusUnauthorized, err.Error())
		return
	}

	if userId != 1 {
		newErrorResponce(c, http.StatusForbidden, "You dont have permition to add movies")
		return
	}

	movieId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponce(c, http.StatusBadRequest, "Invalid ID")
		return
	}

	err = h.service.DeleteMovie(movieId)
	if err != nil {
		newErrorResponce(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, "Movie was deleted from list")
}

func (h *Handler) updateMovie(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		newErrorResponce(c, http.StatusUnauthorized, err.Error())
		return
	}

	if userId != 1 {
		newErrorResponce(c, http.StatusForbidden, "You dont have permission to update movies")
		return
	}

	var movie cinema_diary.Movie

	err = c.BindJSON(&movie)
	if err != nil {
		newErrorResponce(c, http.StatusBadRequest, err.Error())
		return
	}

	err = h.service.UpdateMovie(&movie)
	if err != nil {
		newErrorResponce(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, "Film was updated!")
}
