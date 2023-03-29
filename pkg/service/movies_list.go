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

func (s MovListService) AddToUserMoviesList(moviesList *cinema_diary.MoviesList) error {
	return s.rep.AddToMoviesList(moviesList)
}

func (s MovListService) GetUserMoviesList(userId int, watched bool) ([]*cinema_diary.MoviesList, error) {
	return s.rep.GetUserMoviesList(userId, watched)
}

func (s MovListService) GetFromUserMoviesList(userId int, movieId int) (*cinema_diary.MoviesList, error) {
	return s.rep.GetFromUserMoviesList(userId, movieId)

}
func (s MovListService) UpdateMoviesList(moviesList *cinema_diary.MoviesList) error {
	if _, err := s.rep.GetFromUserMoviesList(moviesList.UserId, moviesList.MovieId); err != nil {
		return err
	}
	return s.rep.UpdateMoviesList(moviesList)
}

func (s MovListService) DeleteFromMovieList(userId int, movieId int) error {
	if _, err := s.rep.GetFromUserMoviesList(userId, movieId); err != nil {
		return err
	}
	return s.rep.DeleteFromMovieList(userId, movieId)
}
