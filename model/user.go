package model

import (
	"encoding/base64"
	"expense-bucket-api/auth"
	"expense-bucket-api/dto"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `gorm:"not null"`
	Email    string `gorm:"uniqueIndex"`
	Password string `gorm:"not null"`
	Salt     []byte `gorm:"not null"`
}

func (u *User) BeforeSave(tx *gorm.DB) error {
	if u.Salt != nil {
		return nil
	}
	passwordOutput, err := auth.DigestPassword(u.Password)
	if err != nil {
		return err
	}
	password := base64.StdEncoding.EncodeToString(passwordOutput.Hash)
	if err != nil {
		return err
	}
	u.Salt = passwordOutput.Salt
	u.Password = password

	return nil
}

func (u User) GetTokenData() auth.TokenData {
	return auth.TokenData{
		Email: u.Email,
		Name:  u.Name,
	}
}

func (u *User) ToDTO() *dto.User {
	return &dto.User{
		Name:      u.Name,
		Email:     u.Email,
		CreatedAt: u.CreatedAt,
	}
}
