package service

import (
	"cinema_diary"
	"cinema_diary/pkg/repository"
	"crypto/sha1"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

const (
	salt     = "1gdfg734tybs"
	tokenTTL = 12 * time.Hour
)

type tokenClaims struct {
	jwt.RegisteredClaims
	Id int `json:"id"`
}

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

func (s *AuthService) GetUsers() []*cinema_diary.User {
	return s.rep.GetUsers()
}

func (s *AuthService) GenerateToken(login, password string) (string, error) {
	user, err := s.rep.GetUser(login, generatePasswordHash(password))
	if err != nil {
		fmt.Println("Error! Incorrect login or password")
		return "", err
	}

	key := []byte("test")
	claims := &tokenClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(5 * time.Hour)),
		},
		Id: user.Id,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	fmt.Println("TOKEN:", token)
	fmt.Println(token.SignedString(key))
	return token.SignedString(key)
}
