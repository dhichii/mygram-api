package helper

import (
	"errors"
	"mygram-api/src/config/env"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func GenerateJWT(id int) string {
	claims := jwt.MapClaims{
		"id":         id,
		"expires_at": time.Now().Add(1 * time.Hour).Unix(),
	}

	secret := env.GetSecretJWTEnv()
	rawToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, _ := rawToken.SignedString([]byte(secret))
	return token
}

func GetAuthorization(c *gin.Context) string {
	return c.Request.Header.Get("Authorization")
}

func GetJWT(header string) string {
	return strings.Split(header, " ")[1]
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

func ValidateJWT(c *gin.Context) (interface{}, error) {
	header := GetAuthorization(c)
	bearer := strings.HasPrefix(strings.ToLower(header), "bearer")
	if !bearer {
		return nil, errors.New("sign in to process")
	}

	token := GetJWT(header)
	return ExtractJWT(token)
}
