package dtapi

import (
	"fmt"

	"github.com/antihax/optional"
	dtapi "github.com/dtcookie/dtapi/libdtapiconf"
)

// WebApplicationConfigAPI TODO: documentation
type WebApplicationConfigAPI struct {
	client  *dtapi.APIClient
	Default DefaultWebApplicationConfigAPI
}

/*
ForID TODO: documentation
@return WebApplicationConfigByIDAPI
*/
func (api WebApplicationConfigAPI) ForID(ID string) WebApplicationConfigByIDAPI {
	waca := WebApplicationConfigByIDAPI{}
	waca.client = api.client
	waca.ID = ID
	return waca
}

/*
Create Create a new web application configuration. The payload must not provide an id as that will be automatically assigned.
@param conf dtapi.WebApplicationConfig
@return dtapi.WebApplicationConfig
@return WebApplicationConfigByIDAPI
*/
func (api WebApplicationConfigAPI) Create(conf dtapi.WebApplicationConfig) (dtapi.WebApplicationConfigStub, WebApplicationConfigByIDAPI, error) {
	result, httpResponse, err := api.client.WebApplicationConfigApi.CreateConfiguration1(nil, &dtapi.CreateConfiguration1Opts{
		WebApplicationConfig: optional.NewInterface(conf),
	})
	if err != nil {
		return result, WebApplicationConfigByIDAPI{}, err
	}
	if httpResponse.StatusCode == 201 {
		return result, api.ForID(result.Identifier), err
	}
	return result, WebApplicationConfigByIDAPI{}, fmt.Errorf("invalid input (http status code %d", httpResponse.StatusCode)
}

/*
ListConfigurations List all application configurations.
@return []dtapi.WebApplicationConfig
@return dtapi.ConfigurationMetadata
*/
func (api WebApplicationConfigAPI) ListConfigurations() ([]dtapi.WebApplicationConfigStub, dtapi.ConfigurationMetadata, error) {
	result, _, err := api.client.WebApplicationConfigApi.ListConfigurations1(nil)
	return result.Values, result.Metadata, err
}

/*
ListDataPrivacySettings List all applications' data privacy settings.
@return []dtapi.ApplicationDataPrivacy
@return dtapi.ConfigurationMetadata
*/
func (api WebApplicationConfigAPI) ListDataPrivacySettings() ([]dtapi.ApplicationDataPrivacy, dtapi.ConfigurationMetadata, error) {
	result, _, err := api.client.WebApplicationConfigApi.ListDataPrivacySettings1(nil)
	return result.Values, result.Metadata, err
}

/*
isValid Validates new web application configuration for the `POST /` request.
 * @param conf ndtapi.WebApplicationConfig
@return bool
*/
func (api WebApplicationConfigAPI) isValid(conf dtapi.WebApplicationConfig) (bool, error) {
	return check204(api.client.WebApplicationConfigApi.ValidateConfiguration2(nil, &dtapi.ValidateConfiguration2Opts{
		WebApplicationConfig: optional.NewInterface(conf),
	}))
}
