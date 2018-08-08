package dtapi

import (
	"github.com/antihax/optional"
	dtapi "github.com/dtcookie/dtapi/libdtapienv"
)

// ThresholdAPI is a convenience wrapper around
// the Services offered via
// "github.com/dtcookie/dtapi/libdtapienv" related to
// configuring Thresholds
type ThresholdAPI envService

// CreateCustomThreshold creates or updates an existing threshold for a plugin or a custom event.
func (api ThresholdAPI) CreateCustomThreshold(thresholdID string, registrationMessage dtapi.ThresholdRegistrationMessage) (dtapi.Threshold, error) {
	result, _, err := api.client.ThresholdApi.CreateCustomThreshold(nil, thresholdID, &dtapi.CreateCustomThresholdOpts{
		ThresholdRegistrationMessage: optional.NewInterface(registrationMessage),
	})
	return result, err
}

// DeleteCustomThreshold deletes the specified threshold.
func (api ThresholdAPI) DeleteCustomThreshold(thresholdID string) error {
	_, err := api.client.ThresholdApi.DeleteCustomThreshold(nil, thresholdID)
	return err
}

// ReadCustomThresholds retrieves all configured thresholds for plugins and custom events in a monitored dtapi.
func (api ThresholdAPI) ReadCustomThresholds(opts *dtapi.ReadCustomThresholdsOpts) ([]dtapi.Threshold, error) {
	result, _, err := api.client.ThresholdApi.ReadCustomThresholds(nil, opts)
	return result, err
}
