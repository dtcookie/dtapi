package dtapi

import (
	"github.com/antihax/optional"
	dtapi "github.com/dtcookie/dtapi/libdtapienv"
)

// ThresholdAPI TODO: documentation
type ThresholdAPI envService

/*
CreateCustomThreshold Creates or updates an existing threshold for a plugin or a custom event.
 * @param thresholdID ID of the threshold to be created or updated.
 * @param optional nil or *CreateCustomThresholdOpts - Optional Parameters:
 * @param "ThresholdRegistrationMessage" (optional.Interface of dtapi.ThresholdRegistrationMessage) -  JSON body of the request, containing threshold parameters.
@return dtapi.Threshold
*/
func (api ThresholdAPI) CreateCustomThreshold(thresholdID string, registrationMessage dtapi.ThresholdRegistrationMessage) (dtapi.Threshold, error) {
	result, _, err := api.client.ThresholdApi.CreateCustomThreshold(nil, thresholdID, &dtapi.CreateCustomThresholdOpts{
		ThresholdRegistrationMessage: optional.NewInterface(registrationMessage),
	})
	return result, err
}

/*
DeleteCustomThreshold Deletes the specified threshold.
 * @param thresholdID The ID of the threshold to be deleted.
*/
func (api ThresholdAPI) DeleteCustomThreshold(thresholdID string) error {
	_, err := api.client.ThresholdApi.DeleteCustomThreshold(nil, thresholdID)
	return err
}

/*
ReadCustomThresholds Gets all configured thresholds for plugins and custom events in a monitored dtapi.
 * @param optional nil or *ReadCustomThresholdsOpts - Optional Parameters:
 * @param "Filter" (optional.String) -  Filters thresholds by the source.
@return []Threshold
*/
func (api ThresholdAPI) ReadCustomThresholds(opts *dtapi.ReadCustomThresholdsOpts) ([]dtapi.Threshold, error) {
	result, _, err := api.client.ThresholdApi.ReadCustomThresholds(nil, opts)
	return result, err
}
