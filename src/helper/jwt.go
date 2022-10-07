package helper

import (
	"errors"
	"mygram-api/src/config/env"
	"time"

	"github.com/golang-jwt/jwt"
)

func GenerateJWT(id int) string {
	claims := jwt.MapClaims{
		"Id":        id,
		"ExpiresAt": time.Now().Add(1 * time.Hour).Unix(),
	}

	secret := env.GetSecretJWTEnv()
	rawToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, _ := rawToken.SignedString([]byte(secret))
	return token
}

func ExtractJWT(tokenStr string) (interface{}, error) {
	secret := env.GetSecretJWTEnv()
	token, _ := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	if _, ok := token.Claims.(jwt.MapClaims); !ok && !token.Valid {
		return nil, errors.New("sign in to process")
	}

	return token.Claims.(jwt.MapClaims), nil
}
