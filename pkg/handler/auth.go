package handler

import (
	"cinema_diary"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) signUp(c *gin.Context) {
	var input cinema_diary.User

	if err := c.BindJSON(&input); err != nil {
		return
	}

	id, err := h.service.Authorization.CreateUser(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		`id`: id,
	})

}

type signInInput struct {
	Login    string `json:"login" binding:"required"`
	Password string `json:"password" binding:"required"`
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
	user := h.service.Authorization.GetUsers()
	fmt.Println(user[0].Name, user[0].Password)
}
