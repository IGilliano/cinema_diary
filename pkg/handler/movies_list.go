package handler

import (
	"cinema_diary"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) addToMoviesList(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		newErrorResponce(c, http.StatusInternalServerError, err.Error())
		return
	}

	var moviesList cinema_diary.MoviesList

	if err = c.BindJSON(&moviesList); err != nil {
		return
	}

	if _, err = h.service.GetMovie(moviesList.MovieId); err != nil {
		newErrorResponce(c, http.StatusBadRequest, "Movie does not exist")
		return
	}

	moviesList.UserId = userId

	err = h.service.AddToUserMoviesList(isWatched(moviesList))
	if err != nil {
		newErrorResponce(c, http.StatusBadRequest, err.Error())
		return
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
		return
	}

	moviesList, err := h.service.GetUserMoviesList(userId, watched)
	if err != nil {
		newErrorResponce(c, http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, moviesList)
}

func (h *Handler) updateMoviesList(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		newErrorResponce(c, http.StatusUnauthorized, err.Error())
		return
	}

	var moviesList cinema_diary.MoviesList

	if err = c.BindJSON(&moviesList); err != nil {
		newErrorResponce(c, http.StatusBadRequest, err.Error())
		return
	}

	moviesList.UserId = userId

	if err = h.service.UpdateMoviesList(&moviesList); err != nil {
		newErrorResponce(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, "Your movies list was updated")
}

func (h *Handler) getFromUserMoviesList(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		newErrorResponce(c, http.StatusUnauthorized, err.Error())
		return
	}

	movieId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponce(c, http.StatusBadRequest, "ID is invalid")
		return
	}

	movieFromList, err := h.service.GetFromUserMoviesList(userId, movieId)
	if err != nil {
		newErrorResponce(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, movieFromList)
}

func (h *Handler) deleteFromMoviesList(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		newErrorResponce(c, http.StatusUnauthorized, err.Error())
		return
	}

	movieId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponce(c, http.StatusBadRequest, "ID is invalid")
		return
	}

	err = h.service.DeleteFromMovieList(userId, movieId)
	if err != nil {
		newErrorResponce(c, http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, "Movie was deleted from your movies list")

}
