package service

import "cinema_diary"

func (s MovService) AddToUserList(moviesList cinema_diary.MoviesList) error {
	return s.rep.AddToUserList(moviesList)
}
