package services

import (
	"buridansAss/chat/src/models"
	"buridansAss/chat/src/repositories"
	"github.com/dgrijalva/jwt-go"
	"github.com/sirupsen/logrus"
	"go.uber.org/fx"
	"time"
)

type (
	Account interface {
		Register(name, email, password string) (*models.User, error)
		Authorize(name, email, password string) (*models.User, error)
	}

	Claims struct {
		Username string `json:"username"`
		jwt.StandardClaims
	}

	accountService struct {
		accountRepo repositories.User
		logger      *logrus.Logger
	}

	/*
	 *любая структура встраивающая fx.In обрабатывается как структура параметров
	 *конструктор ниже будет более читаемым
	 */
	AccountOptions struct {
		fx.In

		AccountRepo repositories.User
		Logger      *logrus.Logger
	}
)

func NewAccount(params AccountOptions) Account {
	return &accountService{
		accountRepo: params.AccountRepo,
		logger:      params.Logger,
	}
}

func GenerateToken(name string) *jwt.Token {
	claims := &Claims{
		Username: name,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(5 * time.Minute).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token
}

func (as *accountService) Register(name, email, password string) (*models.User, error) {
	return as.accountRepo.Create(name, email, password)
}

func (as *accountService) Authorize(name, email, password string) (*models.User, error) {
	user, _ := as.accountRepo.FindByEmail(email)

	return user, nil
}
