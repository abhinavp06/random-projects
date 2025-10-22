package service

import (
	"abhinavp06/campaign-gateway/dto"
	"abhinavp06/campaign-gateway/util"
)

func PushToKafka(campaignId string, batch []dto.UserData, batchNumber int) {
	logger := util.GetLogger()

	logger.Info("Pushing user data to kafka", "total users", len(batch), "campaign ID", campaignId, "batch number", batchNumber)
}