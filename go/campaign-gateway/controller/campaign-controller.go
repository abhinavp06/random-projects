package controller

import (
	"abhinavp06/campaign-gateway/util"
	"net/http"
)

func RunCampaigns(w http.ResponseWriter, r *http.Request) {
	logger := util.GetLogger()
	logger.Info("running campaigns...")
}