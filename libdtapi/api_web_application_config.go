package dtapi

import (
	"errors"

	"github.com/antihax/optional"
	dtapi "github.com/dtcookie/dtapi/libdtapiconf"
)

// WebApplicationConfigByIDAPI TODO: documentation
type WebApplicationConfigByIDAPI struct {
	client *dtapi.APIClient
	ID     string
}

// DataPrivacySettingsAPI TODO: documentation
type DataPrivacySettingsAPI struct {
	client *dtapi.APIClient
	ID     string
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
 * @param c dtapi.WebApplicationConfig
@return WebApplicationConfig
*/
func (api WebApplicationConfigByIDAPI) CreateOrUpdate(c dtapi.WebApplicationConfig) error {
	if (c.Identifier != "") && (c.Identifier != api.ID) {
		return errors.New("identifier of the webapp must either not be set or match the id of this api")
	}
	_, err := api.client.WebApplicationConfigApi.CreateOrUpdateConfiguration1(nil, api.ID, &dtapi.CreateOrUpdateConfiguration1Opts{
		WebApplicationConfig: optional.NewInterface(c),
	})
	// bodyBytes, _ := ioutil.ReadAll(httpResponse.Body)
	// fmt.Println(string(bodyBytes))
	return err
}

/*
Delete Delete the web application configuration with the given id.
*/
func (api WebApplicationConfigByIDAPI) Delete() (bool, error) {
	return check204(api.client.WebApplicationConfigApi.DeleteConfiguration1(nil, api.ID))
}

/*
Get Get the web application configuration with the given id.
 * @return WebApplicationConfig
*/
func (api WebApplicationConfigByIDAPI) Get() (dtapi.WebApplicationConfig, error) {
	result, _, err := api.client.WebApplicationConfigApi.GetConfiguration2(nil, api.ID)
	return result, err
}

/*
ValidateForCreateOrUpdate Validates new web application configuration for the `PUT /{id}` request.
 * @param c dtapi.WebApplicationConfig
*/
func (api WebApplicationConfigByIDAPI) isValid(c dtapi.WebApplicationConfig) (bool, error) {
	return check204(api.client.WebApplicationConfigApi.ValidateConfiguration3(nil, api.ID, &dtapi.ValidateConfiguration3Opts{
		WebApplicationConfig: optional.NewInterface(c),
	}))
}

/*
GET Get the web application's data privacy settings.
 * @param id
@return ApplicationDataPrivacy
*/
func (api DataPrivacySettingsAPI) GET() (dtapi.ApplicationDataPrivacy, error) {
	result, _, err := api.client.WebApplicationConfigApi.GetDataPrivacySettings1(nil, api.ID)
	return result, err
}

/*
UPDATE Updates the web application's data privacy settings.
 * @param id
 * @param optional nil or *UpdateDataPrivacySettings1Opts - Optional Parameters:
 * @param "ApplicationDataPrivacy" (optional.Interface of ApplicationDataPrivacy) -
*/
func (api DataPrivacySettingsAPI) UPDATE(p dtapi.ApplicationDataPrivacy) (bool, error) {
	return check204(api.client.WebApplicationConfigApi.UpdateDataPrivacySettings1(nil, api.ID, &dtapi.UpdateDataPrivacySettings1Opts{
		ApplicationDataPrivacy: optional.NewInterface(p),
	}))
}

/*
IsValid Validates new data privacy settings for the `PUT /{id}/dataPrivacy` request.
 * @param id
 * @param p dtapi.ApplicationDataPrivacy
 * @param "ApplicationDataPrivacy" (optional.Interface of ApplicationDataPrivacy) -
*/
func (api DataPrivacySettingsAPI) IsValid(p dtapi.ApplicationDataPrivacy) (bool, error) {
	return check204(api.client.WebApplicationConfigApi.ValidateDataPrivacySettings1(nil, api.ID, &dtapi.ValidateDataPrivacySettings1Opts{
		ApplicationDataPrivacy: optional.NewInterface(p),
	}))
}
