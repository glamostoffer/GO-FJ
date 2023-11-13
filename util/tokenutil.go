package util

import (
	"GO-FJ/internal/domain"
	jwt "github.com/golang-jwt/jwt/v4"
	"github.com/sirupsen/logrus"
	"time"
)

func CreateAccessToken(user *domain.User, secret string, expiry int) (accessToken string, err error) {
	exp := &jwt.NumericDate{time.Now().Add(time.Hour * time.Duration(expiry))}
	claims := &domain.JwtCustomClaims{
		Name: user.Name,
		ID:   user.ID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: exp,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(secret))
	if err != nil {
		logrus.Errorf("error during creating access token string: %s", err.Error())
		return "", err
	}
	return t, err
}
