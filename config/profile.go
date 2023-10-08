package config

import (
	"flag"
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

func CheckApplicationProfile() {
	profilePtr := flag.String("profile", "", "Specify the application profile.")
	flag.Parse()

	profile := Profile(*profilePtr)

	switch profile {
	case Local, Test:
		err := loadProfileEnvironments(profile)

		if err != nil {
			log.Fatalf("Error loading .env file: %v", err)
		}

	default:
		fmt.Println("Invalid profile. Available profiles: ", Profiles)
		os.Exit(1)
	}
}

func loadProfileEnvironments(profile Profile) error {
	envFile := fmt.Sprintf("resources/%s.env", profile)
	return godotenv.Load(envFile)
}
