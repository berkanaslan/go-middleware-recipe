package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Profile string

const (
	Local Profile = "local"
	Test  Profile = "test"
)

var Profiles = []Profile{Local, Test}

var ActiveProfile Profile

func CheckApplicationProfile() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: ./your_app <profile>")
		os.Exit(1)
	}

	profile := Profile(os.Args[1])

	switch profile {
	case Local, Test:
		err := loadProfileEnvironments(profile)

		if err != nil {
			log.Fatalf("Error loading .env file: %v", err)
		}

		ActiveProfile = profile
	default:
		fmt.Println("Invalid profile. Available profiles: ", Profiles)
		os.Exit(1)
	}
}

func loadProfileEnvironments(profile Profile) error {
	envFile := fmt.Sprintf("config/%s.env", profile)
	return godotenv.Load(envFile)
}
