package service

import (
	"github.com/JamieBShaw/auth-service/domain/model"
	"github.com/JamieBShaw/auth-service/repository"
	"github.com/sirupsen/logrus"
)

type AuthService interface {
	Create(userId int64) (*model.AuthTokens, error)
	Delete(accessUuid string) error
}

type authService struct {
	repo repository.Repository
	log  *logrus.Logger
}

func NewAuthService(repo repository.Repository, log *logrus.Logger) *authService {
	return &authService{repo: repo, log: log}
}

func (a *authService) Create(userId int64) (*model.AuthTokens, error) {
	a.log.Info("[AUTH SERVICE]: Executing Create")

	token, err := model.GetNewAccessToken(userId)
	if err != nil {
		a.log.Errorf("error generating token: %v, for user id: %v", err, userId)
		return nil, err
	}

	err = a.repo.CreateAuth(userId, token)
	if err != nil {
		a.log.Errorf("error creating and setting token: %v, for user id: %v, error: %v", token, userId, err)
		return nil, err
	}

	return &model.AuthTokens{
		AccessToken:  token.AccessToken,
		RefreshToken: token.RefreshToken,
	}, nil
}

func (a *authService) Delete(accessUuid string) error {
	a.log.Info("[AUTH SERVICE]: Executing Delete")

	err := a.repo.DeleteAuth(accessUuid)
	if err != nil {
		a.log.Errorf("error deleting user token with accessUuid: %v, error: %v", accessUuid, err)
		return err
	}

	return nil
}
