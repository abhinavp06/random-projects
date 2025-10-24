package validator

import "abhinavp06/campaign-gateway/types"

func ValidateCampaignRequestBody(campaign types.Campaign) (bool, map[string][]string) {
	errors := map[string][]string{
		"name":       {},
		"cron":       {},
		"filter":     {},
	}

	errors["name"] = validateCampaignName(campaign.Name)
	errors["cron"] = validateCampaignCron(campaign.Cron) // TODO: logic to validate a cron expression
	errors["filter"] = validateCampaignFilter(campaign.Filter) // TODO: logic to validate sql where clause

	hasErrors := false
	for _, errs := range errors {
		if len(errs) > 0 {
			hasErrors = true
			break
		}
	}

	if(hasErrors) {
		return false, errors
	}

	return true, nil
}

func validateCampaignName(campaignName string)  ([]string) {
	var errors []string

	if(campaignName == "") {
		errors = append(errors, "campaign name cannot be empty")
	}

	if(len(campaignName) < 6) {
		errors = append(errors, "campaign name must be greater than or equal to 6 characters")
	}

	if(len(campaignName) > 100) {
		errors = append(errors, "campaign name must be lesser than or equal to 100 characters")
	}

	// TODO: logic to validate campaign name input to avoid sql injection


	if(len(errors) > 0 ) {
		return errors 
	} else {
		return nil
	}
}

func validateCampaignCron(campaignCron string) ([]string) {
	return nil
}

func validateCampaignFilter(campaignFilter *string) ([]string) {
	return nil
}