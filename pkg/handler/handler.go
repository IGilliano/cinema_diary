package handler

import (
	"cinema_diary/pkg/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
		auth.GET("/users", h.getUsers)
	}

	api := router.Group("/api", h.userIdentity)
	{
		api.GET("/movies", h.getMovies)
	}

	return router
}
