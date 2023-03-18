package service

import (
	"cinema_diary"
	"cinema_diary/pkg/repository"
)

type MovListService struct {
	rep repository.MoviesList
}

func NewMovListService(rep repository.MoviesList) *MovListService {
	return &MovListService{rep: rep}
}

func (s MovListService) AddToUserMoviesList(moviesList cinema_diary.MoviesList) error {
	return s.rep.AddToUserList(moviesList)
}

func (s MovListService) GetUserMoviesList(userId int, watched bool) ([]*cinema_diary.MoviesList, error) {
	return s.rep.GetUserMoviesList(userId, watched)
}
