package service

import (
	"cinema_diary"
	"cinema_diary/pkg/repository"
)

type Authorization interface {
	CreateUser(user cinema_diary.User) (int, error)
	GetUsers() ([]*cinema_diary.User, error)
	GenerateToken(login, password string) (string, error)
	ParseToken(token string) (int, error)
}

type Movies interface {
	GetMovies() ([]*cinema_diary.Movie, error)
	GetMovie(id int) (*cinema_diary.Movie, error)
	AddMovies(movies []*cinema_diary.Movie) ([]*int, error)
	UpdateMovie(movie *cinema_diary.Movie) error
	DeleteMovie(id int) error
}

type MoviesList interface {
	AddToUserMoviesList(moviesList *cinema_diary.MoviesList) error
	GetUserMoviesList(userId int, watched bool) ([]*cinema_diary.MoviesList, error)
	GetFromUserMoviesList(userId int, movieId int) (*cinema_diary.MoviesList, error)
	UpdateMoviesList(moviesList *cinema_diary.MoviesList) error
	DeleteFromMovieList(userId int, movieId int) error
}

type Service struct {
	Authorization
	Movies
	MoviesList
}

func NewService(rep *repository.Repository) *Service {
	return &Service{NewAuthService(rep.Authorization), NewMovService(rep.Movies), NewMovListService(rep.MoviesList)}
}
