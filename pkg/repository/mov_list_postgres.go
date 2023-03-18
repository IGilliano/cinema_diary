package repository

import (
	"cinema_diary"
	"github.com/jmoiron/sqlx"
)

type MovListPostgres struct {
	db *sqlx.DB
}

func NewMovListPostgres(db *sqlx.DB) *MovListPostgres {
	return &MovListPostgres{db: db}
}

func (m MovListPostgres) AddToUserList(moviesList cinema_diary.MoviesList) error {
	_, err := m.db.Exec("INSERT INTO movies_list(u_id, m_id, is_watched, is_liked, score) VALUES ($1, $2, $3, $4, $5)", moviesList.UserId, moviesList.MovieId, moviesList.IsWatched, moviesList.IsLiked, moviesList.Score)
	return err
}

func (m MovListPostgres) GetUserMoviesList(userId int, watched bool) ([]*cinema_diary.MoviesList, error) {
	var movieslist []*cinema_diary.MoviesList

	rows, err := m.db.Query("SELECT * FROM movies_list WHERE u_id = $1 and is_watched = $2", userId, watched)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var movieFromList cinema_diary.MoviesList
		err = rows.Scan(&movieFromList.UserId, &movieFromList.MovieId, &movieFromList.IsWatched, &movieFromList.IsLiked, &movieFromList.Score)
		if err != nil {
			return nil, err
		}
		movieslist = append(movieslist, &movieFromList)
	}
	return movieslist, nil
}
