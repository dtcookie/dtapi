package dynatrace

import (
	"errors"

	"github.com/antihax/optional"
	"github.com/dtcookie/dtapi/config"
)

// WebApplicationConfigByIDAPI TODO: documentation
type WebApplicationConfigByIDAPI struct {
	WebApplicationConfigAPI
	ID string
}

// DataPrivacySettingsAPI TODO: documentation
type DataPrivacySettingsAPI struct {
	configService
	ID string
}

/*
DataPrivacySettings Provides an REST API Client that deals with DataPrivacySettings for a Web Application
@teturn DefaultWebApplicationConfigAPI
*/
func (api WebApplicationConfigByIDAPI) DataPrivacySettings() DataPrivacySettingsAPI {
	a := DataPrivacySettingsAPI{}
	a.client = api.client
	a.ID = api.ID
	return a
}

/*
CreateOrUpdate Store the configured web application configuration. If the payload has the id property set, it must match the id in the URL.
 * @param c config.WebApplicationConfig
@return WebApplicationConfig
*/
func (api WebApplicationConfigByIDAPI) CreateOrUpdate(c config.WebApplicationConfig) (config.WebApplicationConfig, error) {
	if (c.Identifier != "") && (c.Identifier != api.ID) {
		return config.WebApplicationConfig{}, errors.New("identifier of the webapp must either not be set or match the id of this api")
	}
	result, _, err := api.client.WebApplicationConfigApi.CreateOrUpdateConfiguration1(nil, api.ID, &config.CreateOrUpdateConfiguration1Opts{
		WebApplicationConfig: optional.NewInterface(c),
	})
	return result, err
}

/*
Delete Delete the web application configuration with the given id.
*/
func (api WebApplicationConfigByIDAPI) Delete() (bool, error) {
	return check204(api.client.WebApplicationConfigApi.DeleteConfiguration1(nil, api.ID))
}

/*
Fetch Get the web application configuration with the given id.
 * @return WebApplicationConfig
*/
func (api WebApplicationConfigByIDAPI) Fetch() (config.WebApplicationConfig, error) {
	result, _, err := api.client.WebApplicationConfigApi.GetConfiguration2(nil, api.ID)
	return result, err
}

/*
ValidateForCreateOrUpdate Validates new web application configuration for the `PUT /{id}` request.
 * @param c config.WebApplicationConfig
*/
func (api WebApplicationConfigByIDAPI) isValid(c config.WebApplicationConfig) (bool, error) {
	return check204(api.client.WebApplicationConfigApi.ValidateConfiguration3(nil, api.ID, &config.ValidateConfiguration3Opts{
		WebApplicationConfig: optional.NewInterface(c),
	}))
}

/*
Fetch Get the web application's data privacy settings.
 * @param id
@return ApplicationDataPrivacy
*/
func (api DataPrivacySettingsAPI) Fetch() (config.ApplicationDataPrivacy, error) {
	result, _, err := api.client.WebApplicationConfigApi.GetDataPrivacySettings1(nil, api.ID)
	return result, err
}

/*
Update Updates the web application's data privacy settings.
 * @param id
 * @param optional nil or *UpdateDataPrivacySettings1Opts - Optional Parameters:
 * @param "ApplicationDataPrivacy" (optional.Interface of ApplicationDataPrivacy) -
*/
func (api DataPrivacySettingsAPI) Update(p config.ApplicationDataPrivacy) (bool, error) {
	return check204(api.client.WebApplicationConfigApi.UpdateDataPrivacySettings1(nil, api.ID, &config.UpdateDataPrivacySettings1Opts{
		ApplicationDataPrivacy: optional.NewInterface(p),
	}))
}

/*
IsValid Validates new data privacy settings for the `PUT /{id}/dataPrivacy` request.
 * @param id
 * @param p config.ApplicationDataPrivacy
 * @param "ApplicationDataPrivacy" (optional.Interface of ApplicationDataPrivacy) -
*/
func (api DataPrivacySettingsAPI) IsValid(p config.ApplicationDataPrivacy) (bool, error) {
	return check204(api.client.WebApplicationConfigApi.ValidateDataPrivacySettings1(nil, api.ID, &config.ValidateDataPrivacySettings1Opts{
		ApplicationDataPrivacy: optional.NewInterface(p),
	}))
}
