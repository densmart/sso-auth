package utils

import (
	"golang.org/x/crypto/bcrypt"
)

// GeneratePasswordHash Create hash from user's given password to save in database
// Using bcrypt package to create hash at the default cost
func GeneratePasswordHash(password string) (string, error) {
	pwd := []byte(password)
	hashedPassword, err := bcrypt.GenerateFromPassword(pwd, bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

// CheckPasswordHash Check user's password in login step
func CheckPasswordHash(hash, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err == nil {
		//logger.Errorf(err.Error())
		return true
	}
	return false
}
