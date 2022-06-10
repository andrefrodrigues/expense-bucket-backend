package auth

import (
	"crypto/ed25519"
	"encoding/base64"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/gommon/log"
)

type TokenData struct {
	Email string
	Name  string
}

var NoPrivateKeyErr = errors.New("Private key missing")
var NoPublicKeyErr = errors.New("Public key missing")
var InvalidTokenFieldErr = errors.New("Invalid token field")

//CreateToken creates a new token
func CreateToken(dto TokenData) (string, error) {

	claims := jwt.MapClaims{
		"email": dto.Email,
		"name":  dto.Name,
		"exp":   time.Now().Add(time.Hour * 72).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodEdDSA, claims)
	keyString := os.Getenv("PRIVATE_KEY")
	if keyString == "" {
		return "", NoPrivateKeyErr
	}
	privBytes, err := base64.StdEncoding.DecodeString(keyString)
	if err != nil {
		return "", err
	}
	priv := ed25519.PrivateKey(privBytes)
	t, err := token.SignedString(priv)
	if err != nil {
		return "", err
	}
	return t, nil
}

func parseToken(token string) (*jwt.Token, error) {
	return jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodEd25519); !ok {

			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		keyString := os.Getenv("PUBLIC_KEY")
		if keyString == "" {
			return nil, NoPublicKeyErr
		}
		pubBytes, err := base64.StdEncoding.DecodeString(keyString)
		if err != nil {
			return nil, err
		}
		pub := ed25519.PublicKey(pubBytes)
		return pub, nil
	})
}

//IsTokenValid verifies if the token is valid
func IsTokenValid(token string) bool {
	tokenValue, err := parseToken(token)
	if err != nil {
		log.Errorf("Error while validating token", err)
		return false
	}

	return tokenValue.Valid
}

//GetTokenData fetches the data of the jwt
func GetTokenData(token string) (*TokenData, error) {
	tokenValue, err := parseToken(token)
	if err != nil {
		log.Errorf("Error while getting token data", err)
		return nil, err
	}

	if claims, ok := tokenValue.Claims.(jwt.MapClaims); ok {

		userName, ok := claims["name"].(string)
		if !ok {
			return nil, InvalidTokenFieldErr
		}
		email, ok := claims["email"].(string)
		if !ok {
			return nil, InvalidTokenFieldErr
		}
		return &TokenData{
			Name:  userName,
			Email: email,
		}, nil
	} else {
		return nil, err
	}
}
