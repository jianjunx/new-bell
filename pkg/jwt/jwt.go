package jwt

import (
	"new-bell/models"
	"time"

	gojwt "github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte("new.bell")

type Claims struct {
	UserName string
	Uid      int
	gojwt.StandardClaims
}

func Award(user *models.UserModal, issuer string) (string, error) {
	expireTime := time.Now().Add(7 * 24 * time.Hour)
	claims := &Claims{
		UserName: user.UserName,
		Uid:      user.Uid,
		StandardClaims: gojwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    issuer,
		},
	}
	token := gojwt.NewWithClaims(gojwt.SigningMethodHS256, claims)
	tokenStr, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}
	return tokenStr, nil
}

func ParseToken(tokenStr string) (gojwt.Token, *Claims, error) {
	claims := &Claims{}
	token, err := gojwt.ParseWithClaims(tokenStr, claims, func(t *gojwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	return *token, claims, err
}
