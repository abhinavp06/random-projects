package usecase

import (
	"abhinavp06/campaign-gateway/dto"
	"abhinavp06/campaign-gateway/repository"
	"abhinavp06/campaign-gateway/util"

	"sync"
)

func RunCampaigns() bool {
	var campaigns []dto.Campaign
	var err error
	logger := util.GetLogger()

	campaigns, err = repository.GetCampaigns()
	if err != nil {
		return false
	}

	var wg sync.WaitGroup

	
	for _, campaign := range campaigns {
		if(util.ValidateCronExpression(campaign.Cron)) {
			logger.Info("Running campaign", "campaign_name", campaign.Name)
			wg.Add(1)
			go func(c dto.Campaign) {
				defer wg.Done()
				repository.ProcessCampaignData(c);
			}(campaign)
		}
	}


	return true
}