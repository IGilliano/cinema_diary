package repository

import (
	"cinema_diary"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user cinema_diary.User) (int, error)
}

type Repository struct {
	Authorization
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{NewAuthPostgres(db)}
}
