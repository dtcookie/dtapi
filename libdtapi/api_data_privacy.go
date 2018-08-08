package dtapi

import (
	"github.com/antihax/optional"
	dtapi "github.com/dtcookie/dtapi/libdtapiconf"
)

// DataPrivacyAPI is a convenience wrapper around
// the Services offered via
// "github.com/dtcookie/dtapi/libdtapiconf" related to
// Data Privacy Configuration.
type DataPrivacyAPI confService

// Get retrieves the Global Data Privacy Settings.
func (api DataPrivacyAPI) Get() (dtapi.DataPrivacy, error) {
	result, _, err := api.client.DataPrivacyApi.GetDataPrivacySettings2(nil)
	return result, err
}

// Update updates (HTTP PUT) the Global Data Privacy Settings
func (api DataPrivacyAPI) Update(dataPrivacy dtapi.DataPrivacy) error {
	_, err := api.client.DataPrivacyApi.UpdateDataPrivacySettings2(nil, &dtapi.UpdateDataPrivacySettings2Opts{
		DataPrivacy: optional.NewInterface(dataPrivacy),
	})
	return err
}

// Validate allows to validate the configuration for Global Data Privacy before
// invoking the actual Update.
func (api DataPrivacyAPI) Validate(dataPrivacy dtapi.DataPrivacy) error {
	_, err := api.client.DataPrivacyApi.ValidateConfiguration4(nil, &dtapi.ValidateConfiguration4Opts{
		DataPrivacy: optional.NewInterface(dataPrivacy),
	})
	return err
}
