package repository

import (
	"go-middleware-recipe/database"
	"go-middleware-recipe/model/core"
)

type UserRepository struct {
	Impl Impl[core.User]
}

func NewUserRepository() *UserRepository {
	return &UserRepository{Impl: Impl[core.User]{}}
}

func (r *UserRepository) FindByEmail(email string) (core.User, error) {
	var user core.User
	tx := database.DBConn.Where("email = ?", email).First(&user)

	return user, tx.Error
}
