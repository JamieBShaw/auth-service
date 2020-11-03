package sqlite

import (
	"database/sql"
	"github.com/JamieBShaw/auth-service/domain/model"
)

type repository struct {
	Db *sql.DB
}

func (r repository) CreateAuth(userId int64, tkn *model.AccessToken) error {
	return nil
}

func NewRepo(db *sql.DB) *repository {
	return &repository{Db: db}
}
