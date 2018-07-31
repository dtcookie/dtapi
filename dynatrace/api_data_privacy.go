package dynatrace

import (
	"github.com/dtcookie/dtapi/config"
)

// DataPrivacyAPI TODO: documentation
type DataPrivacyAPI configService

/*
Get Lists the global data privacy settings.
@return DataPrivacy
*/
func (api DataPrivacyAPI) Get() (config.DataPrivacy, error) {
	result, _, err := api.client.DataPrivacyApi.GetDataPrivacySettings2(nil)
	if err != nil {
		return result, err
	}
	return result, nil
}

/*
Update Updates the global data privacy settings.
This request updates global settings, affecting all your applications. To update application-specific data privacy settings, use the &#x60;PUT /applications/web/{id}/dataPrivacy&#x60; request.
 * @param optional nil or *UpdateDataPrivacySettings2Opts - Optional Parameters:
 * @param "DataPrivacy" (optional.Interface of DataPrivacy) -
*/
func (api DataPrivacyAPI) Update(opts *config.UpdateDataPrivacySettings2Opts) error {
	_, err := api.client.DataPrivacyApi.UpdateDataPrivacySettings2(nil, opts)
	return err
}

/*
Validate Validates new data privacy settings for the `PUT /dataPrivacy` request.
 * @param optional nil or *ValidateConfiguration4Opts - Optional Parameters:
 * @param "DataPrivacy" (optional.Interface of DataPrivacy) -
*/
func (api DataPrivacyAPI) Validate(opts *config.ValidateConfiguration4Opts) error {
	_, err := api.client.DataPrivacyApi.ValidateConfiguration4(nil, opts)
	return err
}
