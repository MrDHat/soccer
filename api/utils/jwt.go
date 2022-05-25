package utils

import (
	"soccer-manager/config"

	"github.com/dgrijalva/jwt-go"
)

// jwt vars
var (
	JWTSigningKey    = config.JWTSigningKey()
	JWTSigningMethod = "HS256"
)

// JwtKey stores the information necessary to generate a jwt token
type JwtKey struct {
	Claims      jwt.MapClaims
	TokenString string
}

// GenerateJWT method will create and set jwt key in TokenString field
func (j *JwtKey) GenerateJWT() error {
	token := jwt.New(jwt.GetSigningMethod(JWTSigningMethod))
	token.Claims = j.Claims
	tokenString, err := token.SignedString([]byte(JWTSigningKey))
	if err != nil {
		return err
	}
	j.TokenString = tokenString
	return nil
}

// ParseJWT method will parse the token string and extract claims fields
func (j *JwtKey) ParseJWT() error {
	claims := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(j.TokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(JWTSigningKey), nil
	})
	if !token.Valid || err != nil {
		return err
	}
	j.Claims = claims
	return nil
}
