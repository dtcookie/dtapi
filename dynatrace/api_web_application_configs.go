package dynatrace

import (
	"fmt"

	"github.com/antihax/optional"
	"github.com/dtcookie/dtapi/config"
)

// WebApplicationConfigAPI TODO: documentation
type WebApplicationConfigAPI struct {
	configService
	Default DefaultWebApplicationConfigAPI
}

/*
ForID TODO: documentation
@return WebApplicationConfigByIDAPI
*/
func (api WebApplicationConfigAPI) ForID(ID string) WebApplicationConfigByIDAPI {
	return WebApplicationConfigByIDAPI{
		WebApplicationConfigAPI: api,
		ID: ID,
	}
}

/*
Create Create a new web application configuration. The payload must not provide an id as that will be automatically assigned.
@param conf config.WebApplicationConfig
@return config.WebApplicationConfig
@return WebApplicationConfigByIDAPI
*/
func (api WebApplicationConfigAPI) Create(conf config.WebApplicationConfig) (config.WebApplicationConfig, WebApplicationConfigByIDAPI, error) {
	result, httpResponse, err := api.client.WebApplicationConfigApi.CreateConfiguration1(nil, &config.CreateConfiguration1Opts{
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
@return []config.WebApplicationConfig
@return config.ConfigurationMetadata
*/
func (api WebApplicationConfigAPI) ListConfigurations() ([]config.WebApplicationConfig, config.ConfigurationMetadata, error) {
	result, _, err := api.client.WebApplicationConfigApi.ListConfigurations1(nil)
	return result.Values, result.Metadata, err
}

/*
ListDataPrivacySettings List all applications' data privacy settings.
@return []config.ApplicationDataPrivacy
@return config.ConfigurationMetadata
*/
func (api WebApplicationConfigAPI) ListDataPrivacySettings() ([]config.ApplicationDataPrivacy, config.ConfigurationMetadata, error) {
	result, _, err := api.client.WebApplicationConfigApi.ListDataPrivacySettings1(nil)
	return result.Values, result.Metadata, err
}

/*
isValid Validates new web application configuration for the `POST /` request.
 * @param conf nconfig.WebApplicationConfig
@return bool
*/
func (api WebApplicationConfigAPI) isValid(conf config.WebApplicationConfig) (bool, error) {
	return check204(api.client.WebApplicationConfigApi.ValidateConfiguration2(nil, &config.ValidateConfiguration2Opts{
		WebApplicationConfig: optional.NewInterface(conf),
	}))
}
