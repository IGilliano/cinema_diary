package service

import (
	"cinema_diary"
	"cinema_diary/pkg/repository"
)

type Authorization interface {
	CreateUser(user cinema_diary.User) (int, error)
	GetUsers() []*cinema_diary.User
	GenerateToken(login, password string) (string, error)
	ParseToken(token string) (int, error)
}

type Movies interface {
	GetMovies() ([]*cinema_diary.Movie, error)
	AddToUserList(moviesList cinema_diary.MoviesList) error
}

type Service struct {
	Authorization
	Movies
}

func NewService(rep *repository.Repository) *Service {
	return &Service{NewAuthService(rep.Authorization), NewMovService(rep.Movies)}
}
