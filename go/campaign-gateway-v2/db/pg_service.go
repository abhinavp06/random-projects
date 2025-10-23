package db

import (
	"abhinavp06/campaign-gateway/shared"
	"context"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

var PgPool *pgxpool.Pool = &pgxpool.Pool{}

func InitializePool() {
	shared.Logger.Info("SERVER_STARTUP: connecting to database")
	
	pool, err := pgxpool.New(context.Background(), shared.Configuration.PgUri)
	if err != nil {
		shared.Logger.Error("SERVER_STARTUP: failed to connect to database", "error", err)
		os.Exit(1)
	}

	PgPool = pool

	shared.Logger.Info("SERVER_STARTUP: connected to database successfully")
	PingDatabase()
}

func PingDatabase() {
	shared.Logger.Info("HEALTH_CHECK: pinging database")
		if err := PgPool.Ping(context.Background()); err != nil {
		shared.Logger.Error("HEALTH_CHECK: unable to ping database", "error", err)
		os.Exit(1) // TODO: figure out if i should exit here or do something else if a ping fails
	} else {
		shared.Logger.Info("HEALTH_CHECK: successfully pinged database")
	}
}