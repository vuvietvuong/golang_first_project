package jwt

import (
	"github.com/golang-jwt/jwt/v4"
	"golang-basic/config"
	"golang-basic/domain/model"
	"os"
	"strconv"
	"time"
)

func GenerateJWT(user *model.User) (string, error) {
	tokenTTL, _ := strconv.Atoi(os.Getenv("TOKEN_TTL"))
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  user.ID,
		"iat": time.Now().Unix(),
		"eat": time.Now().Add(time.Second * time.Duration(tokenTTL)).Unix(),
	})
	return token.SignedString(config.PrivateKey())
}
