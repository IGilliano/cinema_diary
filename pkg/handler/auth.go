package handler

import (
	"cinema_diary"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func (h *Handler) signUp(c *gin.Context) {
	var input cinema_diary.User

	if err := c.BindJSON(&input); err != nil {
		return
	}

	id, err := h.service.Authorization.CreateUser(input)
	if err != nil {
		log.Fatal(err)
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		`id`: id,
	})

}

func (h *Handler) signIn(c *gin.Context) {
	fmt.Println("Oh! PRIVET!")

}
