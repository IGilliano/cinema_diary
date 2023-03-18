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

func (m MovPostgres) GetMovie(id int) (*cinema_diary.Movie, error) {
	var movie cinema_diary.Movie

	rows, err := m.db.Query("SELECT * FROM movies WHERE m_id = $1", id)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		err = rows.Scan(&movie.Id, &movie.Name, &movie.Director, &movie.Year)
		if err != nil {
			return nil, err
		}
	}

	return &movie, err
}

func (m MovPostgres) AddMovies(movies []*cinema_diary.Movie) ([]*int, error) {
	var moviesId []*int
	for _, movie := range movies {
		var id int
		row := m.db.QueryRow("INSERT INTO movies (m_name, director, year) VALUES ($1, $2, $3)", movie.Name, movie.Director, movie.Year)
		if err := row.Scan(&id); err != nil {
			continue
		}
		moviesId = append(moviesId, &id)
	}
	return moviesId, nil
}

func (m MovPostgres) DeleteMovie(id int) error {
	_, err := m.db.Exec("DELETE FROM movies WHERE m_id = $1", id)
	return err
}
