package config

import (
	"fmt"

	"github.com/spf13/viper"
)

func Initiator() {
	viper.AutomaticEnv()

	viper.BindEnv("DATABASE_URL")
	viper.BindEnv("DB_ENGINE")

	requiredEnv := []string{
		"DATABASE_URL",
		"DB_ENGINE",
	}

	for _, env := range requiredEnv {
		if viper.GetString(env) == "" {
			panic(fmt.Sprintf("%s environment variable is required", env))
		}
	}

	fmt.Println("Successfully read environment config")
}
