package dtapi

import (
	"github.com/antihax/optional"
	dtapi "github.com/dtcookie/dtapi/libdtapiconf"
)

// DefaultWebApplicationConfigAPI grants access and allows
// for modification of the configuration of the
// Default Web Application
type DefaultWebApplicationConfigAPI struct {
	confService
	// DataPrivacySettings grants access and allows for modification
	// of the Data Privacy Settings of the Default Web Application
	DataPrivacySettings DefaultWebApplicationDataPrivacySettingsAPI
}

// DefaultWebApplicationDataPrivacySettingsAPI grants access
// and allows for modification of the Data Privacy Settings of the
// Default Web Application
type DefaultWebApplicationDataPrivacySettingsAPI confService

// Get Retrieves the default web application configuration.
func (api DefaultWebApplicationConfigAPI) Get() (dtapi.WebApplicationConfig, error) {
	result, _, err := api.client.WebApplicationConfigApi.GetDefaultApplication1(nil)
	return result, err
}

// Update updates the default web application configuration.
func (api DefaultWebApplicationConfigAPI) Update(c dtapi.WebApplicationConfig) (bool, error) {
	return check204(api.client.WebApplicationConfigApi.CreateOrUpdateDefaultConfiguration1(nil, &dtapi.CreateOrUpdateDefaultConfiguration1Opts{
		WebApplicationConfig: optional.NewInterface(c),
	}))
}

// IsValid validates new default web application configuration for the `PUT /default` request.
func (api DefaultWebApplicationConfigAPI) IsValid(c dtapi.WebApplicationConfig) (bool, error) {
	return check204(api.client.WebApplicationConfigApi.ValidateDefaultConfiguration1(nil, &dtapi.ValidateDefaultConfiguration1Opts{
		WebApplicationConfig: optional.NewInterface(c),
	}))
}

// Get retrieves the default web application's data privacy settings.
func (api DefaultWebApplicationDataPrivacySettingsAPI) Get() (dtapi.ApplicationDataPrivacy, error) {
	result, _, err := api.client.WebApplicationConfigApi.GetDefaultApplicationDataPrivacySettings1(nil)
	return result, err
}

// Update updates the default web application's data privacy settings.
func (api DefaultWebApplicationDataPrivacySettingsAPI) Update(c dtapi.ApplicationDataPrivacy) (bool, error) {
	return check204(api.client.WebApplicationConfigApi.UpdateDefaultApplicationDataPrivacySettings1(nil, &dtapi.UpdateDefaultApplicationDataPrivacySettings1Opts{
		ApplicationDataPrivacy: optional.NewInterface(c),
	}))
}

// isValid allows for validating the payload for updating the Default Web Application's Data Privacy Settings
// without actually invoking the Update.
func (api DefaultWebApplicationDataPrivacySettingsAPI) isValid(c dtapi.ApplicationDataPrivacy) (bool, error) {
	return check204(api.client.WebApplicationConfigApi.ValidateDefaultApplicationDataPrivacySettings1(nil, &dtapi.ValidateDefaultApplicationDataPrivacySettings1Opts{
		ApplicationDataPrivacy: optional.NewInterface(c),
	}))
}
