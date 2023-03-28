package handler

import (
	"cinema_diary"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type signInInput struct {
	Login    string `json:"login" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (h *Handler) signUp(c *gin.Context) {
	var input cinema_diary.User

	if err := c.BindJSON(&input); err != nil {
		newErrorResponce(c, http.StatusBadRequest, "Invalid input body")
		return
	}

	id, err := h.service.Authorization.CreateUser(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		`id`: id,
	})

}

func (h *Handler) signIn(c *gin.Context) {
	var input signInInput

	if err := c.BindJSON(&input); err != nil {
		return
	}

	token, err := h.service.Authorization.GenerateToken(input.Login, input.Password)
	if err != nil {
		fmt.Println(err)
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})

}

func (h *Handler) getUsers(c *gin.Context) {
	user, err := h.service.Authorization.GetUsers()
	if err != nil {
		newErrorResponce(c, http.StatusBadRequest, err.Error())
	}
	c.JSON(http.StatusOK, user)
}
