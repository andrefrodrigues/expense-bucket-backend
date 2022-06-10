package service

import (
	"context"
	"encoding/base64"
	"errors"
	"expense-bucket-api/auth"
	"expense-bucket-api/model"

	"gorm.io/gorm"
)

type SignupDto struct {
	Email                string `json:"email"`
	Name                 string `json:"name"`
	Password             string `json:"password"`
	PasswordConfirmation string `json:"password_confirmation"`
}

type SignupOutput struct {
	Token string `json:"token"`
}

func NewSignupOutput(token string) *SignupOutput {
	return &SignupOutput{token}
}

type AuthService struct {
	db *gorm.DB
}

type AuthServiceConfig struct {
	DB *gorm.DB
}

func NewAuthService(config AuthServiceConfig) *AuthService {
	return &AuthService{
		db: config.DB,
	}
}

var UnmatchedPasswordsErr = errors.New("Passwords don't match")
var ExistingUserErr = errors.New("User exists")
var UserNotFoundErr = errors.New("User not found")
var InvalidPasswordErr = errors.New("Invalid password")

//Signup performs the signup operation for a brand new user!
func (a *AuthService) Signup(ctx context.Context, payload SignupDto) (*SignupOutput, error) {
	if payload.Password != payload.PasswordConfirmation {
		return nil, UnmatchedPasswordsErr
	}
	exists, err := a.userExists(ctx, payload.Email)

	if err != nil {
		return nil, err
	}
	if exists {
		return nil, ExistingUserErr
	}
	user := model.User{
		Name:     payload.Name,
		Password: payload.Password,
		Email:    payload.Email,
	}
	result := a.db.WithContext(ctx).Create(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	token, err := auth.CreateToken(user.GetTokenData())
	if err != nil {
		return nil, err
	}

	return NewSignupOutput(token), nil
}

func (a *AuthService) userExists(ctx context.Context, email string) (bool, error) {
	user := model.User{Email: email}

	result := a.db.WithContext(ctx).First(&user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, result.Error
	}

	return true, nil
}

type LoginDto struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

//Login performs the login operation for an existing user
func (a *AuthService) Login(ctx context.Context, dto LoginDto) (*SignupOutput, error) {
	email := dto.Email
	user := model.User{Email: email}

	if result := a.db.WithContext(ctx).First(&user); result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, UserNotFoundErr
		}
		return nil, result.Error
	}

	hash, err := base64.StdEncoding.DecodeString(user.Password)
	if err != nil {
		return nil, err
	}
	passwordDTO := auth.PasswordDTO{
		Salt: user.Salt,
		Hash: hash,
	}

	if !auth.ValidatePassword(dto.Password, &passwordDTO) {
		return nil, InvalidPasswordErr
	}

	if err != nil {
		return nil, err
	}

	token, err := auth.CreateToken(user.GetTokenData())
	if err != nil {
		return nil, err
	}

	return NewSignupOutput(token), nil
}
