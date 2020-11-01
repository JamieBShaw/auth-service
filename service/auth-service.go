package service

import "github.com/JamieBShaw/auth-service/repository"

type AuthService interface {
	Login(userId int64) error
}

type authService struct {
	repo repository.Repository
}

func NewAuthService(repo repository.Repository) *authService {
	return &authService{repo: repo}
}

func (a authService) Create(userId int64) {
	
}



