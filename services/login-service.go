package services

import (
	"fmt"
	"log"
	"os"
)

type LoginService interface {
	Login(username, password string) bool
}

type loginService struct {
	authorizedUsername, authorizedPassword string
}

func NewLoginService() LoginService {
	credentials := getAuthorized()
	log.Println("credentials: " + fmt.Sprint(credentials))

	return &loginService{
		authorizedUsername: credentials["username"],
		authorizedPassword: credentials["password"],
	}
}

func getAuthorized() map[string]string {
	username := os.Getenv("AUTH_USERNAME")
	if username == "" {
		log.Fatal("AUTH_USERNAME variable not setup in .env")
	}

	password := os.Getenv("AUTH_PASSWORD")
	if password == "" {
		log.Fatal("AUTH_PASSWORD variable not setup in .env")
	}

	return map[string]string{
		"username": username,
		"password": password,
	}
}

func (s *loginService) Login(username, password string) bool {
	return s.authorizedUsername == username && s.authorizedPassword == password
}
