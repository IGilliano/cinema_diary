package repository

import (
	"cinema_diary"
	"github.com/jmoiron/sqlx"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (ap *AuthPostgres) CreateUser(user cinema_diary.User) (int, error) {
	var id int
	row := ap.db.QueryRow("INSERT INTO users (name, login, password) VALUES ($1, $2,$3) RETURNING id", user.Name, user.Login, user.Password)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil

}
