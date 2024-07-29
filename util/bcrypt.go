package util

import (
	"golang.org/x/crypto/bcrypt"
)

func Bcrypt(password string) (string, error) {
	b, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(b), err
}

func ComparePassword(enableBcryptPassword bool, dbPassword string, reqPassword string) bool {
	if enableBcryptPassword {
		err := bcrypt.CompareHashAndPassword([]byte(dbPassword), []byte(reqPassword))
		if err != nil {
			return false
		}
	} else if dbPassword != reqPassword {
		return false
	}
	return true
}
