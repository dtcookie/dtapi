package dtapi

import (
	"errors"

	"github.com/antihax/optional"
	dtapi "github.com/dtcookie/dtapi/libdtapiconf"
)

// WebApplicationConfigByIDAPI is an API
// for accessing and modifying the configuration
// of a specific Web Application
type WebApplicationConfigByIDAPI struct {
	client *dtapi.APIClient
	ID     string
}

// DataPrivacySettingsAPI is an API for accessing
// and modifying the Data Privacy Settings of
// a specific Web Application.
type DataPrivacySettingsAPI struct {
	client *dtapi.APIClient
	ID     string
}

// DataPrivacySettings Provides an REST API Client that deals with DataPrivacySettings for a Web Application
func (api WebApplicationConfigByIDAPI) DataPrivacySettings() DataPrivacySettingsAPI {
	a := DataPrivacySettingsAPI{}
	a.client = api.client
	a.ID = api.ID
	return a
}

// Update stores the configured web application configuration.
// If no such application exists yet, it will be created.
// If the payload has the id property set, it must match the id in the URL.
func (api WebApplicationConfigByIDAPI) Update(c dtapi.WebApplicationConfig) error {
	if (c.Identifier != "") && (c.Identifier != api.ID) {
		return errors.New("identifier of the webapp must either not be set or match the id of this api")
	}
	_, err := api.client.WebApplicationConfigApi.CreateOrUpdateConfiguration1(nil, api.ID, &dtapi.CreateOrUpdateConfiguration1Opts{
		WebApplicationConfig: optional.NewInterface(c),
	})
	return err
}

// Delete deletes the web application
func (api WebApplicationConfigByIDAPI) Delete() (bool, error) {
	return check204(api.client.WebApplicationConfigApi.DeleteConfiguration1(nil, api.ID))
}

// Get retrieves the web applications configuration
func (api WebApplicationConfigByIDAPI) Get() (dtapi.WebApplicationConfig, error) {
	result, _, err := api.client.WebApplicationConfigApi.GetConfiguration2(nil, api.ID)
	return result, err
}

// IsValid allows for validating the configuration of a Web Application before
// invoking the actual Create or Update
func (api WebApplicationConfigByIDAPI) IsValid(c dtapi.WebApplicationConfig) (bool, error) {
	return check204(api.client.WebApplicationConfigApi.ValidateConfiguration3(nil, api.ID, &dtapi.ValidateConfiguration3Opts{
		WebApplicationConfig: optional.NewInterface(c),
	}))
}

// Get retrieves the web application's data privacy settings.
func (api DataPrivacySettingsAPI) Get() (dtapi.ApplicationDataPrivacy, error) {
	result, _, err := api.client.WebApplicationConfigApi.GetDataPrivacySettings1(nil, api.ID)
	return result, err
}

// Update Updates the web application's data privacy settings.
func (api DataPrivacySettingsAPI) Update(p dtapi.ApplicationDataPrivacy) (bool, error) {
	return check204(api.client.WebApplicationConfigApi.UpdateDataPrivacySettings1(nil, api.ID, &dtapi.UpdateDataPrivacySettings1Opts{
		ApplicationDataPrivacy: optional.NewInterface(p),
	}))
}

// IsValid allows for validating the configuration of a Web Application's Data Privacy Settings before
// invoking the actual Update
func (api DataPrivacySettingsAPI) IsValid(p dtapi.ApplicationDataPrivacy) (bool, error) {
	return check204(api.client.WebApplicationConfigApi.ValidateDataPrivacySettings1(nil, api.ID, &dtapi.ValidateDataPrivacySettings1Opts{
		ApplicationDataPrivacy: optional.NewInterface(p),
	}))
}
