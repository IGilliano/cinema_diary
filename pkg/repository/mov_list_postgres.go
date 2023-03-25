package repository

import (
	"cinema_diary"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type MovListPostgres struct {
	db *sqlx.DB
}

func NewMovListPostgres(db *sqlx.DB) *MovListPostgres {
	return &MovListPostgres{db: db}
}

func (m MovListPostgres) AddToMoviesList(moviesList *cinema_diary.MoviesList) error {
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

func (m MovListPostgres) UpdateMoviesList(moviesList *cinema_diary.MoviesList) error {
	_, err := m.db.Exec("UPDATE movies_list SET is_watched = $1, is_liked = $2, score = $3 WHERE u_id = $4 and m_id = $5", moviesList.IsWatched, moviesList.IsLiked, moviesList.Score, moviesList.UserId, moviesList.MovieId)
	return err
}

func (m MovListPostgres) GetFromUserMoviesList(userId int, movieId int) (*cinema_diary.MoviesList, error) {
	var movieFromList cinema_diary.MoviesList
	query := fmt.Sprintf(`SELECT u_id, m_id, is_liked, is_watched, score FROM movies_list WHERE u_id = $1 and m_id = $2`)
	if err := m.db.Get(&movieFromList, query, userId, movieId); err != nil {
		return nil, err
	}
	return &movieFromList, nil
}

func (m MovListPostgres) DeleteFromMovieList(userId int, movieId int) error {
	_, err := m.db.Exec("DELETE FROM movies_list WHERE u_id = $1 and m_id = $2", userId, movieId)
	return err
}
