package controller

import (
	"abhinavp06/campaign-gateway/config"
	"abhinavp06/campaign-gateway/util"
	"abhinavp06/campaign-gateway/usecase"
	"net/http"
)

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	config.PingDatabase()
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

func TriggerCampaigns(w http.ResponseWriter, r *http.Request) {
	logger := util.GetLogger()
	logger.Info("Running campaigns...")

	usecase.RunCampaigns()

	logger.Info("Campaigns triggered successfully")
}