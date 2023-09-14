package core

import "golang.org/x/crypto/bcrypt"

type User struct {
	ID uint `json:"id" gorm:"primary_key"`
	Audit
	Email    string   `json:"email" gorm:"type:varchar(255);unique;not null"`
	Password Password `json:"-" gorm:"type:varchar(255);not null"`
	UserRole UserRole `json:"user_role" gorm:"type:varchar(255);not null;default:'user'"`
	Enabled  bool     `json:"enabled" gorm:"index;default:true"`
}

func (u *User) GetID() uint {
	return u.ID
}

func (u *User) TableName() string {
	return "user"
}

func (u *User) SetPassword(password string) error {
	hashedPassword, err := HashPassword(password)

	if err != nil {
		return err
	}

	pointerToUser := &u
	(*pointerToUser).Password = hashedPassword

	return nil
}

func (u *User) IsPasswordValid(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}
