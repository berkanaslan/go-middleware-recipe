package core

import "golang.org/x/crypto/bcrypt"

type Password string

func HashPassword(password string) (Password, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return "", err
	}

	return Password(hashedPassword), nil
}
