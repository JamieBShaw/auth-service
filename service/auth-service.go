package service

import (
	"github.com/JamieBShaw/auth-service/domain/model"
	"github.com/JamieBShaw/auth-service/repository"
	"github.com/sirupsen/logrus"
)

type AuthService interface {
	Create(userId int64) error
}

type authService struct {
	repo repository.Repository
	log *logrus.Logger
}

func NewAuthService(repo repository.Repository, log *logrus.Logger) *authService {
	return &authService{repo: repo, log: log}
}

func (a *authService) Create(userId int64) error {
	a.log.Info("[AUTH SERVICE]: Executing Create")

	token, err := model.GetNewAccessToken(userId)
	if err != nil {
		a.log.Errorf("error generating token: %v, for user id: %v", err, userId)
	}

	err = a.repo.CreateAuth(userId, token)
	if err != nil {
		a.log.Errorf("error creating and setting token: %v, for user id: %v, error: %v", token, userId, err)
	}

	return nil
}



