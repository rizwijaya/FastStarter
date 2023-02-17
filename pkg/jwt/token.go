package token

import (
	"FastStarter/app/config"
	"errors"

	"github.com/dgrijalva/jwt-go"
)

func ValidateToken(encodedToken string) (*jwt.Token, error) {
	conf, _ := config.New()
	token, err := jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("invalid token")
		}

		return []byte(conf.App.Secret_key), nil
	})

	if err != nil {
		return token, err
	}

	return token, nil
}

func GenerateToken(userID int) (string, error) {
	conf, _ := config.New()
	claim := jwt.MapClaims{}
	claim["user_id"] = userID

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	signedToken, err := token.SignedString(conf.App.Secret_key)
	if err != nil {
		return signedToken, err
	}

	return signedToken, nil
}
