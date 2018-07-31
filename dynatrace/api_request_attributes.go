package dynatrace

import (
	"github.com/dtcookie/dtapi/config"
)

// RequestAttributesAPI TODO: documentation
type RequestAttributesAPI configService

/*
CreateConfiguration2 Creates a new request attribute.
The body must not provide an ID as IDs are automatically assigned.
 * @param optional nil or *CreateConfiguration2Opts - Optional Parameters:
 * @param "RequestAttribute" (optional.Interface of RequestAttribute) -
@return RequestAttributeStub
*/
func (api RequestAttributesAPI) CreateConfiguration2(opts *config.CreateConfiguration2Opts) (config.RequestAttributeStub, error) {
	result, _, err := api.client.RequestAttributesApi.CreateConfiguration2(nil, opts)
	if err != nil {
		return result, err
	}
	return result, nil
}

/*
CreateOrUpdateConfiguration2 Updates an existing request attribute.
 * @param id The ID of the request attribute to be updated.   If you set the ID in the body as well, it must match this ID.
 * @param optional nil or *CreateOrUpdateConfiguration2Opts - Optional Parameters:
 * @param "RequestAttribute" (optional.Interface of RequestAttribute) -
*/
func (api RequestAttributesAPI) CreateOrUpdateConfiguration2(id string, opts *config.CreateOrUpdateConfiguration2Opts) error {
	_, err := api.client.RequestAttributesApi.CreateOrUpdateConfiguration2(nil, id, opts)
	return err
}

/*
DeleteConfiguration2 Deletes the specified request attribute.
 * @param id The ID of the request attribute to be deleted.
*/
func (api RequestAttributesAPI) DeleteConfiguration2(id string) error {
	_, err := api.client.RequestAttributesApi.DeleteConfiguration2(nil, id)
	return err
}

/*
GetConfiguration3 Shows the properties of the specified request attribute.
 * @param id The ID of the required request attribute.
 * @param optional nil or *GetConfiguration3Opts - Optional Parameters:
 * @param "IncludeProcessGroupReferences" (optional.Bool) -  Flag to include process group references to the response.    Process Group group references aren't compatible across environments.
@return RequestAttribute
*/
func (api RequestAttributesAPI) GetConfiguration3(id string, opts *config.GetConfiguration3Opts) (config.RequestAttribute, error) {
	result, _, err := api.client.RequestAttributesApi.GetConfiguration3(nil, id, opts)
	if err != nil {
		return result, err
	}
	return result, nil
}

/*
ListConfigurations2 Lists all available request attributes.
@return RequestAttributeStubListDto
*/
func (api RequestAttributesAPI) ListConfigurations2() (config.RequestAttributeStubListDto, error) {
	result, _, err := api.client.RequestAttributesApi.ListConfigurations2(nil)
	if err != nil {
		return result, err
	}
	return result, nil
}

/*
ValidateConfiguration5 Validate updates of existing request attribute for the `PUT /requestAttributes/{id}` request.
 * @param id The ID of the request attribute to be validated.
 * @param optional nil or *ValidateConfiguration5Opts - Optional Parameters:
 * @param "RequestAttribute" (optional.Interface of RequestAttribute) -
*/
func (api RequestAttributesAPI) ValidateConfiguration5(id string, opts *config.ValidateConfiguration5Opts) error {
	_, err := api.client.RequestAttributesApi.ValidateConfiguration5(nil, id, opts)
	return err
}

/*
ValidateConfiguration6 Validates new request attributes for the `POST /requestAttributes` request.
 * @param  - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param optional nil or *ValidateConfiguration6Opts - Optional Parameters:
 * @param "RequestAttribute" (optional.Interface of RequestAttribute) -
*/
func (api RequestAttributesAPI) ValidateConfiguration6(opts *config.ValidateConfiguration6Opts) error {
	_, err := api.client.RequestAttributesApi.ValidateConfiguration6(nil, opts)
	return err
}
