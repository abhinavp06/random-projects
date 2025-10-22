package repository

import (
	"abhinavp06/campaign-gateway/config"
	"abhinavp06/campaign-gateway/dto"
	"abhinavp06/campaign-gateway/service"
	"abhinavp06/campaign-gateway/util"

	"context"
	"fmt"
	"time"
)

const batchSize = 1000;

func GetCampaigns() ([]dto.Campaign, error) {
	cfg := config.GetConfig()
	logger := util.GetLogger()

	connectionPool := cfg.PgPool

	campaigns, err := connectionPool.Query(context.Background(), "SELECT * FROM campaigns WHERE enabled = true")

	if err != nil {
		logger.Error("Failed to fetch campaigns", "error", err)
		return nil, err
	}

	var campaignList []dto.Campaign

	for campaigns.Next() {
		var campaign dto.Campaign
		err := campaigns.Scan(&campaign.Id, &campaign.Name, &campaign.Cron, &campaign.Filter, &campaign.Enabled, &campaign.LastRun, &campaign.CreatedAt, &campaign.UpdatedAt)
		if err != nil {
			logger.Error("Failed to scan campaign row", "error", err)
			return nil, err
		}
		campaignList = append(campaignList, campaign)
	}

	return campaignList, nil
}

func ProcessCampaignData(campaign dto.Campaign) {
	cfg := config.GetConfig()
	logger := util.GetLogger()
	ctx := context.Background()

	offset := 0
	var batchNumber int = 1
	for {
		query := fmt.Sprintf("SELECT * FROM user_data WHERE %s LIMIT $1 OFFSET $2", campaign.Filter)
		rows, err := cfg.PgPool.Query(ctx, query, batchSize, offset)
		if err != nil {
			logger.Error("Failed to query user_data", "campaign", campaign.Name, "error", err)
		}

		var batch []dto.UserData
		rowCount := 0
		for rows.Next() {
			var ud dto.UserData
			err := rows.Scan(&ud.Id, &ud.Name, &ud.Age, &ud.Email, &ud.Mobile, &ud.Organization, &ud.JoiningDate, &ud.CreatedAt)
			if err != nil {
				logger.Error("Failed to scan user_data row", "error", err)
				rows.Close()
			}
			batch = append(batch, ud)
			rowCount++
		}
		rows.Close()

		if len(batch) > 0 {
			service.PushToKafka(campaign.Id, batch, batchNumber)
			batchNumber++
		}

		if rowCount < batchSize {
			break
		}
		offset += batchSize
	}
	_, err := cfg.PgPool.Exec(ctx, "UPDATE campaigns SET last_run = $1 WHERE id = $2", time.Now(), campaign.Id)
	if err != nil {
		logger.Error("Failed to update last_run", "campaign", campaign.Name, "error", err)
	}

	logger.Info("Campaign processed successfully", "campaign", campaign.Name)
}
