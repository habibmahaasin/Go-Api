package jwttoken

import (
	"errors"
	"gop-api/app/config"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type JwtToken interface {
	GenerateJwt(user_id string, name string) (string, error)
	ValidateJwt(encode string) (*jwt.Token, error)
}

type jwtToken struct{}

func NewJwtToken() *jwtToken {
	return &jwtToken{}
}

func (j *jwtToken) GenerateJwt(user_uuid string, name string) (string, error) {
	conf, _ := config.Init()
	secret := []byte(conf.App.Secret)
	claim := jwt.MapClaims{}
	claim["user_id"] = user_uuid
	claim["name"] = name
	claim["exp"] = time.Now().Add(time.Minute * 2880).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	signedToken, err := token.SignedString(secret)
	if err != nil {
		return signedToken, err
	}

	return signedToken, nil
}

func (j *jwtToken) ValidateJwt(encode string) (*jwt.Token, error) {
	conf, _ := config.Init()
	secret := []byte(conf.App.Secret)

	token, err := jwt.Parse(encode, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("Invalid Token")
		}

		return []byte(secret), nil
	})

	if err != nil {
		return token, err
	}

	return token, nil
}
