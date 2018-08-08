package dtapi

import (
	"github.com/antihax/optional"
	dtapi "github.com/dtcookie/dtapi/libdtapiconf"
)

// ApplicationDetectionConfigAPI is a cleaned up wrapper around the
// Configuration Service for Application Detection offered via
// "github.com/dtcookie/dtapi/libdtapiconf"
type ApplicationDetectionConfigAPI confService

// Get retrieves the current configuration for Application Detection
func (api *ApplicationDetectionConfigAPI) Get() (dtapi.ApplicationDetectionConfig, error) {
	result, _, err := api.client.ApplicationDetectionConfigApi.GetConfiguration(nil)
	return result, err
}

// Update updates (HTTP PUT) the application detection configuration.
func (api *ApplicationDetectionConfigAPI) Update(c dtapi.ApplicationDetectionConfig) error {
	_, err := api.client.ApplicationDetectionConfigApi.UpdateConfiguration(nil, &dtapi.UpdateConfigurationOpts{
		ApplicationDetectionConfig: optional.NewInterface(c),
	})
	return err
}

// Validate allows for validating a configuration for Application Detection
// before performing the actual Update.
func (api *ApplicationDetectionConfigAPI) Validate(c dtapi.ApplicationDetectionConfig) error {
	_, err := api.client.ApplicationDetectionConfigApi.ValidateConfiguration(nil, &dtapi.ValidateConfigurationOpts{
		ApplicationDetectionConfig: optional.NewInterface(c),
	})
	return err
}
