package config

import (
	"fmt"
	"os"
	"sync"

	"github.com/joho/godotenv"
)


type Config struct {
	Environment string
	Port        string
}


var (
	instance *Config
	once     sync.Once
)


func LoadConfig() {
	once.Do(func() {
		err := godotenv.Load()
		if err != nil {
			fmt.Println("Error loading .env file")
		}

		instance = &Config{
			Environment: os.Getenv("ENVIRONMENT"),
			Port:        os.Getenv("PORT"),
		}
	})
}

func GetConfig() *Config {
	if instance == nil {
		panic("Config not initialized! Call LoadConfig() first in main.go")
	}
	return instance
}