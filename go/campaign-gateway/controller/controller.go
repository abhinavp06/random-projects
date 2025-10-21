package controller

import (
	"abhinavp06/campaign-gateway/util"
	"abhinavp06/campaign-gateway/config"
	"net/http"
)

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	config.PingDatabase()
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

func RunCampaigns(w http.ResponseWriter, r *http.Request) {
	logger := util.GetLogger()
	logger.Info("Running campaigns...")
}