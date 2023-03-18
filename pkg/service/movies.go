package service

import (
	"cinema_diary"
	"cinema_diary/pkg/repository"
)

type MovService struct {
	rep repository.Movies
}

func NewMovService(rep repository.Movies) *MovService {
	return &MovService{rep: rep}
}

func (s MovService) AddMovies(movies []*cinema_diary.Movie) ([]*int, error) {
	return s.rep.AddMovies(movies)
}

func (s MovService) GetMovies() ([]*cinema_diary.Movie, error) {
	return s.rep.GetMovies()
}

func (s MovService) GetMovie(id int) (*cinema_diary.Movie, error) {
	return s.rep.GetMovie(id)
}

func (s MovService) DeleteMovie(id int) error {
	return s.rep.DeleteMovie(id)
}
