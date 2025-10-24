package controller

import (
	"abhinavp06/campaign-gateway/controller/validator"
	"abhinavp06/campaign-gateway/db"
	"abhinavp06/campaign-gateway/types"
	"encoding/json"
	"net/http"
)

func CreateCampaign(w http.ResponseWriter, r *http.Request) {
	var campaign types.Campaign

	if err := json.NewDecoder(r.Body).Decode(&campaign); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	isValid, errors := validator.ValidateCampaignRequestBody(campaign)

	if !isValid {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]any{
			"error":  "Validation failed",
			"fields": errors,
		})
		return
	}

	isInserted := db.InsertCampaign(campaign.Name, campaign.Cron, *campaign.Filter)

	if !isInserted {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Failed to insert record"))
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("OK"))
}
