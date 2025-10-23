package shared

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Environment string
	Port string
	PgUri string
}

var Configuration *Config = &Config{}

func InitConfig() {
	err := godotenv.Load()
	if err != nil && !os.IsNotExist(err) {
		Logger.Error("Error loading .env. Exiting....", "error", err)
		os.Exit(1)
	}

	Configuration.Environment = os.Getenv("ENVIRONMENT")
	Configuration.Port = os.Getenv("PORT")
	Configuration.PgUri = os.Getenv("PG_URI")
}