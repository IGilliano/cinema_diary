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

func (s MovService) GetMovies() ([]*cinema_diary.Movie, error) {
	return s.rep.GetMovies()
}
