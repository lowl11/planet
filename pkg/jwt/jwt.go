package jwt

import (
	"encoding/json"
	"errors"
	"github.com/golang-jwt/jwt/v4"
)

type JWT struct {
	key []byte
}

func New(key []byte) *JWT {
	return &JWT{
		key: key,
	}
}

func (jwt JWT) Generate(object any) (string, error) {
	return Generate(jwt.key, object)
}

func (jwt JWT) Parse(token string, export any) error {
	return Parse(jwt.key, token, export)
}

func Generate(key []byte, object any) (string, error) {
	objectInBytes, err := json.Marshal(object)
	if err != nil {
		return "", err
	}

	objectMap := make(map[string]any)
	if err = json.Unmarshal(objectInBytes, &objectMap); err != nil {
		return "", err
	}

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	for claimsKey, claimsValue := range objectMap {
		claims[claimsKey] = claimsValue
	}

	stringToken, err := token.SignedString(key)
	if err != nil {
		return "", err
	}

	return stringToken, nil
}

func Parse(key []byte, token string, export any) error {
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("parse JWT error")
		}

		return key, nil
	})
	if err != nil {
		return err
	}

	claimsMapInBytes, err := json.Marshal(parsedToken.Claims)
	if err != nil {
		return err
	}

	return json.Unmarshal(claimsMapInBytes, &export)
}
