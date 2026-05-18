package utils

import (
	"crypto/sha256"
	"errors"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func HashingPassword(password string) (string, error) {
	// hash 256 always produces 32 bytes
	hash := sha256.Sum256([]byte(password))
	preHashed := fmt.Sprintf("%x", hash) //  2bytes hex * 32 bytes  generate the hex string 64 hex char

	bytes, err := bcrypt.GenerateFromPassword([]byte(preHashed), 12)
	if err != nil {
		return "", errors.New("Something went wrong hashing")
	}
	return string(bytes), nil
}

func VerifyPassword(password string, hash string) bool {

	hashPassword := sha256.Sum256([]byte(password))

	preHashed := fmt.Sprintf("%x", hashPassword)

	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(preHashed))
	return err == nil
}
