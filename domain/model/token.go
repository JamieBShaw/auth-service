package model

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/twinj/uuid"
	"os"
	"time"
)

const (
	expirationTime = 1
)

type AccessToken struct {
	AccessToken string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	
	AccessUuid string `json:"access_uuid"`
	RefreshUuid string `json:"refresh_uuid"`
	
	AtExpires   int64  `json:"at_expires"`
	RtExpires   int64  `json:"rt_expires"`

	UserId      int64  `json:"user_id"`
	ClientId    int64  `json:"client_id,omitempty"`
}

func GetNewAccessToken(userId int64) (*AccessToken, error) {

	tkn := &AccessToken{}
	tkn.AtExpires = time.Now().Add(expirationTime * time.Hour).Unix()
	tkn.AccessUuid = uuid.NewV4().String()

	tkn.RtExpires = time.Now().Add(time.Hour * 24 * 7).Unix()
	tkn.RefreshUuid = uuid.NewV4().String()

	var err error
	//Creating Access Token
	os.Setenv("ACCESS_SECRET", "jdnfksdmfksd")
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["access_uuid"] = tkn.AccessUuid
	atClaims["user_id"] = userId
	atClaims["exp"] = tkn.AtExpires

	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	tkn.AccessToken, err = at.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
	if err != nil {
		return nil, err
	}
	//Creating Refresh Token
	os.Setenv("REFRESH_SECRET", "mcmvmkmsdnfsdmfdsjf") //this should be in an env file
	rtClaims := jwt.MapClaims{}
	rtClaims["refresh_uuid"] = tkn.RefreshUuid
	rtClaims["user_id"] = userId
	rtClaims["exp"] = tkn.RtExpires
	rt := jwt.NewWithClaims(jwt.SigningMethodHS256, rtClaims)
	tkn.RefreshToken, err = rt.SignedString([]byte(os.Getenv("REFRESH_SECRET")))
	if err != nil {
		return nil, err
	}


	return tkn, nil
}
