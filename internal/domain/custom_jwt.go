package domain

import "github.com/golang-jwt/jwt/v4"

type JwtCustomClaims struct {
	Name string `json:"name"`
	Role string `json:"role"`
	ID   int64  `json:"id"`
	jwt.RegisteredClaims
}

type JwtCustomRefreshClaims struct {
	Role string `json:"role"`
	ID   int64  `json:"id"`
	jwt.RegisteredClaims
}
