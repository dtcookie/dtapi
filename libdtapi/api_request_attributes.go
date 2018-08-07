package dtapi

import (
	"github.com/antihax/optional"
	dtapi "github.com/dtcookie/dtapi/libdtapiconf"
)

// RequestAttributesAPI TODO: documentation
type RequestAttributesAPI confService

/*
CREATE Creates (HTTP POST) a new request attribute.
The body must not provide an ID as IDs are automatically assigned.
 * @param optional nil or *CreateConfiguration2Opts - Optional Parameters:
 * @param "RequestAttribute" (optional.Interface of RequestAttribute) -
@return RequestAttributeStub
*/
func (api RequestAttributesAPI) CREATE(requestAttribute dtapi.RequestAttribute) (dtapi.RequestAttributeStub, error) {
	result, _, err := api.client.RequestAttributesApi.CreateConfiguration2(nil, &dtapi.CreateConfiguration2Opts{
		RequestAttribute: optional.NewInterface(requestAttribute),
	})
	return result, err
}

/*
UPDATE Updates (or creates) via HTTP POST an existing request attribute.
 * @param id The ID of the request attribute to be updated.   If you set the ID in the body as well, it must match this ID.
 * @param optional nil or *CreateOrUpdateConfiguration2Opts - Optional Parameters:
 * @param "RequestAttribute" (optional.Interface of RequestAttribute) -
*/
func (api RequestAttributesAPI) UPDATE(id string, requestAttribute dtapi.RequestAttribute) error {
	_, err := api.client.RequestAttributesApi.CreateOrUpdateConfiguration2(nil, id, &dtapi.CreateOrUpdateConfiguration2Opts{
		RequestAttribute: optional.NewInterface(requestAttribute),
	})
	return err
}

/*
DELETE Deletes the specified request attribute.
 * @param id The ID of the request attribute to be deleted.
*/
func (api RequestAttributesAPI) DELETE(id string) error {
	_, err := api.client.RequestAttributesApi.DeleteConfiguration2(nil, id)
	return err
}

/*
GET Shows the properties of the specified request attribute.
 * @param id The ID of the required request attribute.
 * @param optional nil or *GetConfiguration3Opts - Optional Parameters:
 * @param "IncludeProcessGroupReferences" (optional.Bool) -  Flag to include process group references to the response.    Process Group group references aren't compatible across environments.
@return RequestAttribute
*/
func (api RequestAttributesAPI) GET(id string, withProcessGroupReferences bool) (dtapi.RequestAttribute, error) {
	result, _, err := api.client.RequestAttributesApi.GetConfiguration3(nil, id, &dtapi.GetConfiguration3Opts{
		IncludeProcessGroupReferences: optional.NewBool(withProcessGroupReferences),
	})
	if err != nil {
		return result, err
	}
	return result, nil
}

/*
ALL Lists all available request attributes.
@return RequestAttributeStubListDto
*/
func (api RequestAttributesAPI) ALL() ([]dtapi.RequestAttributeStub, error) {
	result, _, err := api.client.RequestAttributesApi.ListConfigurations2(nil)
	return result.Values, err
}

/*
ValidateForUpdate Validate updates of existing request attribute for the `PUT /requestAttributes/{id}` request.
 * @param id The ID of the request attribute to be validated.
 * @param optional nil or *ValidateConfiguration5Opts - Optional Parameters:
 * @param "RequestAttribute" (optional.Interface of RequestAttribute) -
*/
func (api RequestAttributesAPI) ValidateForUpdate(requestAttribute dtapi.RequestAttribute) error {
	_, err := api.client.RequestAttributesApi.ValidateConfiguration5(nil, &dtapi.ValidateConfiguration5Opts{
		RequestAttribute: optional.NewInterface(requestAttribute),
	})
	return err
}

/*
ValidateForCreate Validates new request attributes for the `POST /requestAttributes` request.
 * @param  - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param optional nil or *ValidateConfiguration6Opts - Optional Parameters:
 * @param "RequestAttribute" (optional.Interface of RequestAttribute) -
*/
func (api RequestAttributesAPI) ValidateForCreate(id string, requestAttribute dtapi.RequestAttribute) error {
	_, err := api.client.RequestAttributesApi.ValidateConfiguration6(nil, id, &dtapi.ValidateConfiguration6Opts{
		RequestAttribute: optional.NewInterface(requestAttribute),
	})
	return err
}
