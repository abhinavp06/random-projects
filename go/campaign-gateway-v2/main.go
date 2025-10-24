package main

import (
	"abhinavp06/campaign-gateway/controller"
	"abhinavp06/campaign-gateway/db"
	"abhinavp06/campaign-gateway/shared"

	"net/http"
	"os"
)

func main() {
	shared.InitLogger()
	shared.InitConfig()
	db.InitializePool()

	shared.Logger.Info("SERVER_STARTUP (LOG): starting campaign-gateway")

	port := ":" + shared.Configuration.Port

	http.HandleFunc("GET /health", controller.HealthCheck)
	http.HandleFunc("POST /campaigns", controller.CreateCampaign)

	err := http.ListenAndServe(port, nil)

	if err != nil {
		shared.Logger.Error("SERVER_STARTUP (ERROR): failed to start server", "error", err)
		os.Exit(1)
	}
}