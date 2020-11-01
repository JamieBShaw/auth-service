package redis

import (
	"github.com/JamieBShaw/auth-service/domain/model"
	"github.com/go-redis/redis/v7"
	"strconv"
	"time"
)

type repo struct {
	client *redis.Client
}

func NewRepo(client *redis.Client) *repo {
	return &repo{client: client}
}

func (r repo) CreateAuth(userId int64, tkn *model.AccessToken) error {
	at := time.Unix(tkn.AtExpires, 0) //converting Unix to UTC(to Time object)
	rt := time.Unix(tkn.RtExpires, 0)
	now := time.Now()

	errAccess := r.client.Set(tkn.AccessUuid, strconv.Itoa(int(userId)), at.Sub(now)).Err()
	if errAccess != nil {
		return errAccess
	}
	errRefresh := r.client.Set(tkn.RefreshUuid, strconv.Itoa(int(userId)), rt.Sub(now)).Err()
	if errRefresh != nil {
		return errRefresh
	}
	return nil
}

