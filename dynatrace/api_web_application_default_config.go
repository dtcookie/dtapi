package dynatrace

import (
	"github.com/antihax/optional"
	"github.com/dtcookie/dtapi/config"
)

// DefaultWebApplicationConfigAPI TODO: documentation
type DefaultWebApplicationConfigAPI struct {
	configService
	DataPrivacySettings DefaultWebApplicationDataPrivacySettingsAPI
}

// DefaultWebApplicationDataPrivacySettingsAPI TODO: documentation
type DefaultWebApplicationDataPrivacySettingsAPI configService

/*
Fetch Retrieves the default web application configuration.
@return WebApplicationConfig
*/
func (api DefaultWebApplicationConfigAPI) Fetch() (config.WebApplicationConfig, error) {
	result, _, err := api.client.WebApplicationConfigApi.GetDefaultApplication1(nil)
	return result, err
}

/*
CreateOrUpdate Update the default web application configuration.
 * @param c config.WebApplicationConfig
*/
func (api DefaultWebApplicationConfigAPI) CreateOrUpdate(c config.WebApplicationConfig) (bool, error) {
	return check204(api.client.WebApplicationConfigApi.CreateOrUpdateDefaultConfiguration1(nil, &config.CreateOrUpdateDefaultConfiguration1Opts{
		WebApplicationConfig: optional.NewInterface(c),
	}))
}

/*
IsValid Validates new default web application configuration for the `PUT /default` request.
 * @param c config.WebApplicationConfig
*/
func (api DefaultWebApplicationConfigAPI) IsValid(c config.WebApplicationConfig) (bool, error) {
	return check204(api.client.WebApplicationConfigApi.ValidateDefaultConfiguration1(nil, &config.ValidateDefaultConfiguration1Opts{
		WebApplicationConfig: optional.NewInterface(c),
	}))
}

/*
Fetch Retrieves the default web application's data privacy settings.
@return ApplicationDataPrivacy
*/
func (api DefaultWebApplicationDataPrivacySettingsAPI) Fetch() (config.ApplicationDataPrivacy, error) {
	result, _, err := api.client.WebApplicationConfigApi.GetDefaultApplicationDataPrivacySettings1(nil)
	return result, err
}

/*
Update Update the default web application's data privacy settings.
 * @param c config.ApplicationDataPrivacy
*/
func (api DefaultWebApplicationDataPrivacySettingsAPI) Update(c config.ApplicationDataPrivacy) (bool, error) {
	return check204(api.client.WebApplicationConfigApi.UpdateDefaultApplicationDataPrivacySettings1(nil, &config.UpdateDefaultApplicationDataPrivacySettings1Opts{
		ApplicationDataPrivacy: optional.NewInterface(c),
	}))
}

/*
isValid Validates new data privacy settings for the `PUT /default/dataPrivacy` request.
 * @param optional nil or *ValidateDefaultApplicationDataPrivacySettings1Opts - Optional Parameters:
 * @param "ApplicationDataPrivacy" (optional.Interface of ApplicationDataPrivacy) -
*/
func (api DefaultWebApplicationDataPrivacySettingsAPI) isValid(c config.ApplicationDataPrivacy) (bool, error) {
	return check204(api.client.WebApplicationConfigApi.ValidateDefaultApplicationDataPrivacySettings1(nil, &config.ValidateDefaultApplicationDataPrivacySettings1Opts{
		ApplicationDataPrivacy: optional.NewInterface(c),
	}))
}
