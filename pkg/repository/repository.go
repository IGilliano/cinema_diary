package repository

import (
	"cinema_diary"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user cinema_diary.User) (int, error)
	GetUsers() []*cinema_diary.User
	GetUser(login, password string) (cinema_diary.User, error)
}

type Movies interface {
	GetMovies() ([]*cinema_diary.Movie, error)
	GetMovie(id int) (*cinema_diary.Movie, error)
	AddMovies(movies []*cinema_diary.Movie) ([]*int, error)
	DeleteMovie(id int) error
}

type MoviesList interface {
	GetUserMoviesList(userId int, watched bool) ([]*cinema_diary.MoviesList, error)
	AddToUserList(moviesList cinema_diary.MoviesList) error
}

type Repository struct {
	Authorization
	Movies
	MoviesList
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{NewAuthPostgres(db), NewMovPostgres(db), NewMovListPostgres(db)}
}
