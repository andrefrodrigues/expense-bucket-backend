package auth

import "errors"

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

	token, err := CreateToken(payload)
	if err != nil {
		return nil, err
	}

	return NewSignupOutput(token), nil
}
