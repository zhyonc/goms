package util

import "golang.org/x/crypto/bcrypt"

func Bcrypt(password string) (string, error) {
	b, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(b), err
}

func ComparePassword(dbPassword string, reqPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(dbPassword), []byte(reqPassword))
	return err == nil
}
