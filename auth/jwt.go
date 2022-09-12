package auth

import (
	"app/configs"
	"app/models"
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

type JWTClaim struct {
	Id    uint   `json:"id"`
	Name  string `json:"username"`
	Email string `json:"email"`
	jwt.StandardClaims
}

func GenerateJWT(user models.User) (tokenString string, err error) {
	expirationTime := time.Now().Add(1 * time.Hour)

	claims := &JWTClaim{
		Id:    user.ID,
		Email: user.Email,
		Name:  user.Name,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err = token.SignedString([]byte(configs.EnvString("JWT_STRING")))

	return tokenString, err
}

func ValidateToken(signedToken string) (err error) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&JWTClaim{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(configs.EnvString("JWT_STRING")), nil
		},
	)

	if err != nil {
		return err
	}

	claims, ok := token.Claims.(*JWTClaim)

	if !ok {
		err = errors.New("couldn't parse claims")
		return err
	}

	if claims.ExpiresAt < time.Now().Local().Unix() {
		err = errors.New("token expired")
		return err
	}
	return nil
}
