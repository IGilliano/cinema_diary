package service

import (
	"cinema_diary"
	"cinema_diary/pkg/repository"
	"crypto/sha1"
	"fmt"
)

const salt = "1gdfg734tybs"

type AuthService struct {
	rep repository.Authorization
}

func NewAuthService(rep repository.Authorization) *AuthService {
	return &AuthService{rep: rep}
}

func (s *AuthService) CreateUser(user cinema_diary.User) (int, error) {
	user.Password = generatePasswordHash(user.Password)

	return s.rep.CreateUser(user)
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
