package utils

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

func GenerateToken(ttl time.Duration, payload interface{}, secretKeyJWT string) (string, error) {

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	now := time.Now().UTC()

	claims["sub"] = payload
	claims["iat"] = now.Unix()
	claims["exp"] = now.Add(ttl).Unix()
	claims["nbf"] = now.Unix()

	tokenString, err := token.SignedString([]byte(secretKeyJWT))

	if err != nil {
		return "", err
	}
	return tokenString, nil

}

func ValidateToken(token string, signedJWTkey string) (interface{}, error) {

	tok, err := jwt.Parse(token, func(jwtToken *jwt.Token) (interface{}, error) {
		if _, ok := jwtToken.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("UnExpected method %s", jwtToken.Header["alg"])
		}
		return []byte(signedJWTkey), nil

	})
	if err != nil {
		return nil, fmt.Errorf("invalid token %w", err)
	}
	claims, ok := tok.Claims.(jwt.MapClaims)
	if !ok || !tok.Valid {
		return nil, errors.New("invalid token claims")
	}
	return claims["sub"],nil
}
