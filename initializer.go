package main

import (
	"go-middleware-recipe/model/core"
	"go-middleware-recipe/repository"
	"log"
)

var userRepository = repository.NewUserRepository()

func InitializeData() {
	createDefaultUser()
}

func createDefaultUser() {
	log.Default().Println("Default user creation started.")

	if count, _ := userRepository.Impl.Count(); count > 0 {
		log.Default().Println("Default user creation skipped.")
		return
	}

	superAdmin := core.User{
		Email:    "superadmin@localhost",
		UserRole: core.Admin,
		Enabled:  true,
	}

	_ = superAdmin.SetPassword("password")
	_, _ = userRepository.Impl.Create(superAdmin)
}
