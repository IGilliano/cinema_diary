package repository

import (
	"cinema_diary"
	"fmt"
	"github.com/jmoiron/sqlx"
	"log"
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

func (ap *AuthPostgres) GetUsers() []*cinema_diary.User {
	var users []*cinema_diary.User
	rows, err := ap.db.Query("SELECT * FROM users")
	if err != nil {
		log.Fatal(err)
		return nil
	}

	for rows.Next() {
		var user cinema_diary.User
		err = rows.Scan(&user.Id, &user.Name, &user.Login, &user.Password)
		if err != nil {
			return nil
		}
		users = append(users, &user)
	}
	return users
}

func (ap *AuthPostgres) GetUser(login, password string) (cinema_diary.User, error) {
	var user cinema_diary.User
	query := fmt.Sprintf("SELECT id FROM users WHERE login=$1 AND password=$2")
	err := ap.db.Get(&user, query, login, password)

	return user, err
}
