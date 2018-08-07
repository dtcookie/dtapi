package dtapi

import (
	"github.com/antihax/optional"
	dtapi "github.com/dtcookie/dtapi/libdtapiconf"
)

// DataPrivacyAPI TODO: documentation
type DataPrivacyAPI confService

/*
GET Lists the global data privacy settings.
@return DataPrivacy
*/
func (api DataPrivacyAPI) GET() (dtapi.DataPrivacy, error) {
	result, _, err := api.client.DataPrivacyApi.GetDataPrivacySettings2(nil)
	return result, err
}

/*
UPDATE Updates (HTTP PUT) the global data privacy settings.
This request updates global settings, affecting all your applications. To update application-specific data privacy settings, use the &#x60;PUT /applications/web/{id}/dataPrivacy&#x60; request.
 * @param optional nil or *UpdateDataPrivacySettings2Opts - Optional Parameters:
 * @param "DataPrivacy" (optional.Interface of DataPrivacy) -
*/
func (api DataPrivacyAPI) UPDATE(dataPrivacy dtapi.DataPrivacy) error {
	_, err := api.client.DataPrivacyApi.UpdateDataPrivacySettings2(nil, &dtapi.UpdateDataPrivacySettings2Opts{
		DataPrivacy: optional.NewInterface(dataPrivacy),
	})
	return err
}

/*
VALIDATE Validates new data privacy settings for the `PUT /dataPrivacy` request.
 * @param optional nil or *ValidateConfiguration4Opts - Optional Parameters:
 * @param "DataPrivacy" (optional.Interface of DataPrivacy) -
*/
func (api DataPrivacyAPI) VALIDATE(dataPrivacy dtapi.DataPrivacy) error {
	_, err := api.client.DataPrivacyApi.ValidateConfiguration4(nil, &dtapi.ValidateConfiguration4Opts{
		DataPrivacy: optional.NewInterface(dataPrivacy),
	})
	return err
}
