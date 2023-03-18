package handler

import (
	"cinema_diary"
	"errors"
	_ "errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

const (
	authorizationHeader = "Authorization"
	userCtx             = "userId"
)

func (h *Handler) userIdentity(c *gin.Context) {
	header := c.GetHeader(authorizationHeader)
	if header == "" {
		newErrorResponce(c, http.StatusUnauthorized, "empty auth header")
		return
	}

	fmt.Println(header)

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		newErrorResponce(c, http.StatusUnauthorized, "invalid auth header")
	}

	userId, err := h.service.Authorization.ParseToken(headerParts[1])
	if err != nil {
		newErrorResponce(c, http.StatusUnauthorized, err.Error())
		return
	}

	c.Set(userCtx, userId)
}

func getUserId(c *gin.Context) (int, error) {
	id, ok := c.Get(userCtx)
	if !ok {
		return 0, errors.New("User not found")
	}

	idInt, ok := id.(int)
	if !ok {
		return 0, errors.New("User id is of invaled type")
	}
	return idInt, nil
}

func isWatched(ml cinema_diary.MoviesList) cinema_diary.MoviesList {
	if ml.Score != 0 || ml.IsLiked != false {
		ml.IsWatched = true
	}
	return ml
}
