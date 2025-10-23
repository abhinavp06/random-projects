package controller

import (
	"abhinavp06/campaign-gateway/db"
	"net/http"
)

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	db.PingDatabase()

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}