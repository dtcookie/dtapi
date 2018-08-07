package dtapi

import (
	"github.com/antihax/optional"
	dtapi "github.com/dtcookie/dtapi/libdtapiconf"
)

// DefaultWebApplicationConfigAPI TODO: documentation
type DefaultWebApplicationConfigAPI struct {
	confService
	DataPrivacySettings DefaultWebApplicationDataPrivacySettingsAPI
}

// DefaultWebApplicationDataPrivacySettingsAPI TODO: documentation
type DefaultWebApplicationDataPrivacySettingsAPI confService

/*
GET Retrieves the default web application configuration.
@return WebApplicationConfig
*/
func (api DefaultWebApplicationConfigAPI) Get() (dtapi.WebApplicationConfig, error) {
	result, _, err := api.client.WebApplicationConfigApi.GetDefaultApplication1(nil)
	return result, err
}

/*
CreateOrUpdate Update the default web application configuration.
 * @param c dtapi.WebApplicationConfig
*/
func (api DefaultWebApplicationConfigAPI) CreateOrUpdate(c dtapi.WebApplicationConfig) (bool, error) {
	return check204(api.client.WebApplicationConfigApi.CreateOrUpdateDefaultConfiguration1(nil, &dtapi.CreateOrUpdateDefaultConfiguration1Opts{
		WebApplicationConfig: optional.NewInterface(c),
	}))
}

/*
IsValid Validates new default web application configuration for the `PUT /default` request.
 * @param c dtapi.WebApplicationConfig
*/
func (api DefaultWebApplicationConfigAPI) IsValid(c dtapi.WebApplicationConfig) (bool, error) {
	return check204(api.client.WebApplicationConfigApi.ValidateDefaultConfiguration1(nil, &dtapi.ValidateDefaultConfiguration1Opts{
		WebApplicationConfig: optional.NewInterface(c),
	}))
}

/*
GET Retrieves the default web application's data privacy settings.
@return ApplicationDataPrivacy
*/
func (api DefaultWebApplicationDataPrivacySettingsAPI) GET() (dtapi.ApplicationDataPrivacy, error) {
	result, _, err := api.client.WebApplicationConfigApi.GetDefaultApplicationDataPrivacySettings1(nil)
	return result, err
}

/*
Update Update the default web application's data privacy settings.
 * @param c dtapi.ApplicationDataPrivacy
*/
func (api DefaultWebApplicationDataPrivacySettingsAPI) Update(c dtapi.ApplicationDataPrivacy) (bool, error) {
	return check204(api.client.WebApplicationConfigApi.UpdateDefaultApplicationDataPrivacySettings1(nil, &dtapi.UpdateDefaultApplicationDataPrivacySettings1Opts{
		ApplicationDataPrivacy: optional.NewInterface(c),
	}))
}

/*
isValid Validates new data privacy settings for the `PUT /default/dataPrivacy` request.
 * @param optional nil or *ValidateDefaultApplicationDataPrivacySettings1Opts - Optional Parameters:
 * @param "ApplicationDataPrivacy" (optional.Interface of ApplicationDataPrivacy) -
*/
func (api DefaultWebApplicationDataPrivacySettingsAPI) isValid(c dtapi.ApplicationDataPrivacy) (bool, error) {
	return check204(api.client.WebApplicationConfigApi.ValidateDefaultApplicationDataPrivacySettings1(nil, &dtapi.ValidateDefaultApplicationDataPrivacySettings1Opts{
		ApplicationDataPrivacy: optional.NewInterface(c),
	}))
}
