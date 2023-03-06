package repository

import (
	"cinema_diary"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type MovPostgres struct {
	db *sqlx.DB
}

func NewMovPostgres(db *sqlx.DB) *MovPostgres {
	return &MovPostgres{db: db}

}

func (m MovPostgres) GetMovies() ([]*cinema_diary.Movie, error) {
	var movies []*cinema_diary.Movie
	row := fmt.Sprintf("SELECT * FROM movies")
	rows, err := m.db.Query(row)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var movie cinema_diary.Movie
		err = rows.Scan(&movie.Id, &movie.Name, &movie.Director, &movie.Year)
		if err != nil {
			return nil, err
		}
		movies = append(movies, &movie)
	}

	return movies, nil
}

func (m MovPostgres) AddToUserList(moviesList cinema_diary.MoviesList) error {
	_, err := m.db.Exec("INSERT INTO movies_list(u_id, m_id, is_watched, is_liked, score) VALUES ($1, $2, $3, $4, $5)", moviesList.UserId, moviesList.MovieId, moviesList.IsWatched, moviesList.IsLiked, moviesList.Score)
	return err
}
