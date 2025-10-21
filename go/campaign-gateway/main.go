package main

import (
	"abhinavp06/campaign-gateway/config"
	"abhinavp06/campaign-gateway/util"

	"net/http"
)

func main() {
	// ** load logger and config at application startup
	util.InitLogger()
	logger := util.GetLogger()
	config.LoadConfig()
	cfg := config.GetConfig()

	port := cfg.Port
	if port == "" {
		port = "3010"
	}

	if port[0] != ':' {
		port = ":" + port
	}

	logger.Info("Starting campaign-gateway")
	logger.Info("campaign-gateway is up and running on port", "port", port)
	
	err := http.ListenAndServe(port, nil)
	if err != nil {
		logger.Error("Server failed to start", "error", err)
	}
}