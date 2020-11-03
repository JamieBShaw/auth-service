package redis

import (
	"github.com/JamieBShaw/auth-service/domain/model"
	"github.com/go-redis/redis/v7"
	"github.com/sirupsen/logrus"
	"strconv"
	"time"
)

type repo struct {
	client *redis.Client
	log    *logrus.Logger
}

func NewRepo(client *redis.Client, log *logrus.Logger) *repo {
	return &repo{client: client, log: log}
}

func (r *repo) CreateAuth(userId int64, tkn *model.AccessToken) error {
	r.log.Info("[REDIS REPO] Executing CreateAuthToken")

	at := time.Unix(tkn.AtExpires, 0) //converting Unix to UTC(to Time object)
	rt := time.Unix(tkn.RtExpires, 0)
	now := time.Now()

	err := r.client.Set(tkn.AccessUuid, strconv.Itoa(int(userId)), at.Sub(now)).Err()
	if err != nil {
		return err
	}
	err = r.client.Set(tkn.RefreshUuid, strconv.Itoa(int(userId)), rt.Sub(now)).Err()
	if err != nil {
		return err
	}
	return nil
}

func (r *repo) DeleteAuth(accessUid string) error {
	r.log.Info("[REDIS REPO] Executing CreateAuthToken")

	deleted, err := r.client.Del(accessUid).Result()
	if err != nil || deleted == 0 {
		return err
	}

	return nil
}
