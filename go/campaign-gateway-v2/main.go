package main

import (
	"abhinavp06/campaign-gateway/shared"
	"abhinavp06/campaign-gateway/db"

	"net/http"
	"os"
)

func main() {
	shared.InitLogger()
	shared.InitConfig()
	db.InitializePool()

	shared.Logger.Info("SERVER_STARTUP (LOG): starting campaign-gateway")

	port := ":" + shared.Configuration.Port

	err := http.ListenAndServe(port, nil)

	if err != nil {
		shared.Logger.Error("SERVER_STARTUP (ERROR): failed to start server", "error", err)
		os.Exit(1)
	}
}