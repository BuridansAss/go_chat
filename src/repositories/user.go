package repositories

import "buridansAss/chat/src/models"

type User interface {
	Create(name, email, password string) *models.User
}

func Create(name, email, pass string) *models.User {
	return &models.User{
		Name:     name,
		Email:    email,
		Password: pass,
		Token:    "token",
	}
}
