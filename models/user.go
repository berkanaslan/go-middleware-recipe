package models

import (
	"go-middleware-recipe/models/base"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	base.Entity
	base.Audit
	Email    string        `json:"email" gorm:"type:varchar(255)"`
	Password string        `json:"password" gorm:"type:varchar(255)"`
	UserRole base.UserRole `json:"user_role" gorm:"type:varchar(255)"`
	Enabled  bool          `json:"enabled"`
}

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return "", err
	}

	return string(hashedPassword), nil
}

func (u *User) SetPassword(password string) error {
	hashedPassword, err := HashPassword(password)

	if err != nil {
		return err
	}

	u.Password = hashedPassword

	return nil
}

func (u *User) CheckPasswordHash(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}
