package utils

import (
	"errors"
	"os"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

var jwtSecretKey = []byte(os.Getenv("JWT_SECRET_KEY"))

type JWTClaim struct {
	Uuid string `json:"uuid"`
	jwt.StandardClaims
}

type ValidateBody struct {
	Uuid string `json:"uuid"`
}

func NewToken(uuid string, expireTime time.Time) (string, error) {
	claims := &JWTClaim{
		Uuid: uuid,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecretKey)
}

func ValidateToken(signedToken string, uuid string) (err error) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&JWTClaim{},
		func(token *jwt.Token) (interface{}, error) {
			return jwtSecretKey, nil
		},
	)

	if err != nil {
		return err
	}

	claims, ok := token.Claims.(*JWTClaim)
	if !ok {
		err = errors.New("Couldn't parse claims")
		return err
	}

	if claims.ExpiresAt < time.Now().Local().Unix() {
		err = errors.New("Token expired")
		return err
	}

	if claims.Uuid != uuid {
		err = errors.New("Claims do not match body")
		return err
	}

	return nil
}

func RefreshToken(token string) (string, error) {
	claims := &JWTClaim{}

	jwt, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtSecretKey, nil
	})

	if err != nil || !jwt.Valid {
		return "", errors.New("Token not valid")
	}

	newAccessToken, err := NewToken(claims.Uuid, time.Now().Add(time.Minute*30))

	if err != nil {
		return "", errors.New("Error while creating new token")
	}

	return newAccessToken, nil
}
