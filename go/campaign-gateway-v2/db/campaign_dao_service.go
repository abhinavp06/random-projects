package db

import (
	"abhinavp06/campaign-gateway/shared"
	"context"

	"github.com/google/uuid"
)

func InsertCampaign(campaignName string, campaignCron string, campaignFilter string) bool {
    campaignId := uuid.New().String()


    query := "INSERT INTO campaigns (id, name, cron, filter) VALUES ('" + campaignId + "', '" + campaignName + "', '" + campaignCron + "', '" + campaignFilter + "');"

    _, err := PgPool.Exec(context.Background(), query)

    if(err != nil) {
        shared.Logger.Error("DB_OPERATION: failed to insert campaign record", "error", err) // TODO: better error logging - cover specific errors
        return false
    }

    return true
}