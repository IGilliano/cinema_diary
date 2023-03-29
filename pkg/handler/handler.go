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
		auth.GET("/", h.getUsers)
	}

	api := router.Group("/api", h.userIdentity)
	{
		movies := api.Group("/movies")
		{
			movies.GET("/", h.getMovies)
			movies.GET("/:id", h.getMovie)
			movies.POST("/", h.addMovies)
			movies.PUT("/", h.updateMovie)
			movies.DELETE("/:id", h.deleteMovie)
		}
		moviesList := api.Group("/movies-list")
		{
			moviesList.GET("/", h.getUserMoviesList)
			moviesList.GET("/watchlist", h.getUserMoviesList)
			moviesList.GET("/:id", h.getFromUserMoviesList)
			moviesList.POST("/", h.addToMoviesList)
			moviesList.PUT("/", h.updateMoviesList)
			moviesList.DELETE("/:id", h.deleteFromMoviesList)
		}
	}

	return router
}
