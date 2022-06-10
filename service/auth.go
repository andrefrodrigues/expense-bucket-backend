package service

import (
	"errors"
	"expense-bucket-api/auth"
)

type SignupDto struct {
	Email                string `json:"email"`
	Name                 string `json:"name"`
	Password             string `json:"password"`
	PasswordConfirmation string `json:"password_confirmation"`
}

func (dto SignupDto) GetTokenData() auth.TokenData {
	return auth.TokenData{
		Email: dto.Email,
		Name:  dto.Name,
	}
}

type SignupOutput struct {
	Token string `json:"token"`
}

func NewSignupOutput(token string) *SignupOutput {
	return &SignupOutput{token}
}

type AuthService struct {
}

func NewAuthService() *AuthService {
	return &AuthService{}
}

var UnmatchedPasswordsErr = errors.New("Passwords don't match")

func (a *AuthService) Signup(payload SignupDto) (*SignupOutput, error) {
	if payload.Password != payload.PasswordConfirmation {
		return nil, UnmatchedPasswordsErr
	}

	token, err := auth.CreateToken(payload.GetTokenData())
	if err != nil {
		return nil, err
	}

	return NewSignupOutput(token), nil
}
