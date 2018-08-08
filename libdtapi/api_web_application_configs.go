package dtapi

import (
	"fmt"

	"github.com/antihax/optional"
	dtapi "github.com/dtcookie/dtapi/libdtapiconf"
)

// WebApplicationConfigAPI deals with accessing
// and modifying the configuration of Web Applications
type WebApplicationConfigAPI struct {
	client *dtapi.APIClient
	// Default grants access and allows for modification
	// of the Default Web Application
	Default DefaultWebApplicationConfigAPI
}

// ForID provides an API for accessing and modifying the
// configuration of a specific Web application.
func (api WebApplicationConfigAPI) ForID(ID string) WebApplicationConfigByIDAPI {
	waca := WebApplicationConfigByIDAPI{}
	waca.client = api.client
	waca.ID = ID
	return waca
}

// Create creates a new web application configuration.
// The payload must not provide an id as that will be automatically assigned.
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

// ListConfigurations retrives id/name pairs for all currently configured web applications.
func (api WebApplicationConfigAPI) ListConfigurations() ([]dtapi.WebApplicationConfigStub, dtapi.ConfigurationMetadata, error) {
	result, _, err := api.client.WebApplicationConfigApi.ListConfigurations1(nil)
	return result.Values, result.Metadata, err
}

// ListDataPrivacySettings retrieves the Data Privacy Settings for all currently configured Web Applications.
func (api WebApplicationConfigAPI) ListDataPrivacySettings() ([]dtapi.ApplicationDataPrivacy, dtapi.ConfigurationMetadata, error) {
	result, _, err := api.client.WebApplicationConfigApi.ListDataPrivacySettings1(nil)
	return result.Values, result.Metadata, err
}

// isValid Validates new web application configuration for the `POST /` request.
func (api WebApplicationConfigAPI) isValid(conf dtapi.WebApplicationConfig) (bool, error) {
	return check204(api.client.WebApplicationConfigApi.ValidateConfiguration2(nil, &dtapi.ValidateConfiguration2Opts{
		WebApplicationConfig: optional.NewInterface(conf),
	}))
}
