package structures

import (
	"github.com/dgrijalva/jwt-go/v4"
)

type Claims struct {
	jwt.StandardClaims
	Username string   `json:"username"`
	UserId   int      `json:"id"`
	Roles    []string `json:"roles"`
}
