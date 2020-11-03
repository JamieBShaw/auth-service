package repository

import "github.com/JamieBShaw/auth-service/domain/model"

type Repository interface {
	CreateAuth(userId int64, tkn *model.AccessToken) error
	DeleteAuth(accessUid string) error
}
