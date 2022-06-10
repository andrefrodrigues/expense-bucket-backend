package auth

import (
	"crypto/ed25519"
	"encoding/base64"
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type TokenData struct {
	Email string
	Name  string
}

var NoPrivateKeyErr = errors.New("Private key missing")

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
