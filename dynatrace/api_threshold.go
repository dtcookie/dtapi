package dynatrace

import (
	"github.com/dtcookie/dtapi/environment"
)

// ThresholdAPI TODO: documentation
type ThresholdAPI service

/*
CreateCustomThreshold Creates or updates an existing threshold for a plugin or a custom event.
 * @param thresholdID ID of the threshold to be created or updated.
 * @param optional nil or *CreateCustomThresholdOpts - Optional Parameters:
 * @param "ThresholdRegistrationMessage" (optional.Interface of environment.ThresholdRegistrationMessage) -  JSON body of the request, containing threshold parameters.
@return environment.Threshold
*/
func (api ThresholdAPI) CreateCustomThreshold(thresholdID string, opts *environment.CreateCustomThresholdOpts) (environment.Threshold, error) {
	result, _, err := api.client.ThresholdApi.CreateCustomThreshold(nil, thresholdID, opts)
	if err != nil {
		return result, err
	}
	return result, nil
}

/*
DeleteCustomThreshold Deletes the specified threshold.
 * @param thresholdID The ID of the threshold to be deleted.
*/
func (api ThresholdAPI) DeleteCustomThreshold(thresholdID string) error {
	_, err := api.client.ThresholdApi.DeleteCustomThreshold(nil, thresholdID)
	if err != nil {
		return err
	}
	return nil
}

/*
ReadCustomThresholds Gets all configured thresholds for plugins and custom events in a monitored environment.
 * @param optional nil or *ReadCustomThresholdsOpts - Optional Parameters:
 * @param "Filter" (optional.String) -  Filters thresholds by the source.
@return []Threshold
*/
func (api ThresholdAPI) ReadCustomThresholds(opts *environment.ReadCustomThresholdsOpts) ([]environment.Threshold, error) {
	result, _, err := api.client.ThresholdApi.ReadCustomThresholds(nil, opts)
	if err != nil {
		return result, err
	}
	return result, nil
}
