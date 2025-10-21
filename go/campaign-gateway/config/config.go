package config

import (
	"abhinavp06/campaign-gateway/util"
	"context"
	"os"
	"sync"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

type Config struct {
	Environment  string
	Port         string
	PgPool       *pgxpool.Pool
}

var (
	instance *Config
	once     sync.Once
)

func LoadConfig() {
	once.Do(func() {
		logger := util.GetLogger()
		err := godotenv.Load()
		if err != nil && !os.IsNotExist(err) {
			logger.Error("Error loading .env file", "error", err)
			os.Exit(1)
		}

		instance = &Config{
			Environment: os.Getenv("ENVIRONMENT"),
			Port:        os.Getenv("PORT"),
		}

		pool, err := pgxpool.New(context.Background(), os.Getenv("PG_URI"))

		if err != nil {
			logger.Error("Unable to create database pool", "error", err)
			os.Exit(1)
		}

		logger.Info("Connected to Postgres database pool successfully!")
		instance.PgPool = pool

		PingDatabase()
	})
}

func GetConfig() *Config {
	if instance == nil {
		panic("Config not initialized!")
	}
	return instance
}

func PingDatabase() {
	logger := util.GetLogger()
	logger.Info("Pinging database...")
	if err := instance.PgPool.Ping(context.Background()); err != nil {
		logger.Error("Unable to ping database", "error", err)
		os.Exit(1)
	} else {
		logger.Info("Successfully pinged database")
	}
}