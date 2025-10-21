package main

import (
	"abhinavp06/campaign-gateway/config"
	"abhinavp06/campaign-gateway/util"
)

func main() {
	// ** load logger and config at application startup
	util.InitLogger()
	logger := util.GetLogger()
	config.LoadConfig()

	logger.Info("Starting campaign-gateway")
	logger.Info("campaign-gateway is up and running!")
}