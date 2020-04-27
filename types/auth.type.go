package types

import (
	"github.com/dgrijalva/jwt-go"
)

type AuthInfo struct {
	ID     int    `json:"id"`
	Pseudo string `json:"pseudo"`
	Pays   string `json:"pays"`
	Ville  string `json:"ville"`
	jwt.StandardClaims
}
