package auth

import (
	"crypto/rand"
	"os"
	"strconv"

	"golang.org/x/crypto/argon2"
)

type PasswordDTO struct {
	Hash []byte
	Salt []byte
}

//DigestPassword digests the password to be stored in the database.
func DigestPassword(password string) (*PasswordDTO, error) {
	salt, err := generateSalt()
	if err != nil {
		return nil, err
	}
	hash := argon2.IDKey([]byte(password), salt, 2, uint32(15360), uint8(1), uint32(16))

	return &PasswordDTO{
		hash,
		salt,
	}, nil
}

//ValidatePassword validates a plaintext password to a digested password and checks if they match
func ValidatePassword(password string, digest *PasswordDTO) bool {
	salt := digest.Salt

	hash := argon2.IDKey([]byte(password), salt, 2, uint32(15360), uint8(1), uint32(16))
	digestedHash := digest.Hash

	if len(digestedHash) != len(hash) {
		return false
	}

	for i := range hash {
		if hash[i] != digestedHash[i] {
			return false
		}
	}
	return true
}

func generateSalt() ([]byte, error) {
	saltLength := os.Getenv("SALT_LENGTH")
	if saltLength == "" {
		saltLength = "16"
	}
	saltLen, err := strconv.Atoi(saltLength)
	if err != nil {
		return nil, err
	}

	salt := make([]byte, saltLen)
	_, err = rand.Read(salt)
	if err != nil {
		return nil, err
	}

	return salt, nil
}
