package dtapi

import (
	"github.com/antihax/optional"
	dtapi "github.com/dtcookie/dtapi/libdtapiconf"
)

// RequestAttributesAPI is a convenience wrapper around
// the Services offered via
// "github.com/dtcookie/dtapi/libdtapiconf" related to
// configuring Request Attributes
type RequestAttributesAPI confService

// Create creates (HTTP POST) a new request attribute.
//
// Note:
//	The request attribute configuration must NOT contain any IDs
//	as these are being determined during creation.
//
// TODO:
//	perform ID check for body offline before sending request
//
// Returns an ID/Name pair for the created Request Attribute
func (api RequestAttributesAPI) Create(attribute dtapi.RequestAttribute) (dtapi.RequestAttributeStub, error) {
	result, _, err := api.client.RequestAttributesApi.CreateConfiguration2(nil, &dtapi.CreateConfiguration2Opts{
		RequestAttribute: optional.NewInterface(attribute),
	})
	return result, err
}

// Update updates (or creates) via HTTP POST an existing request attribute.
//
// TODO:
//	perform ID check for body offline before sending request
//
// Parameters:
//	- ID	is the unique identifier of the request attribute to be updated.
//			If you set the ID in the attribute as well, it must match this ID.
func (api RequestAttributesAPI) Update(ID string, attribute dtapi.RequestAttribute) error {
	_, err := api.client.RequestAttributesApi.CreateOrUpdateConfiguration2(nil, ID, &dtapi.CreateOrUpdateConfiguration2Opts{
		RequestAttribute: optional.NewInterface(attribute),
	})
	return err
}

// Delete deletes the specified request attribute.
func (api RequestAttributesAPI) Delete(ID string) error {
	_, err := api.client.RequestAttributesApi.DeleteConfiguration2(nil, ID)
	return err
}

// Get retrieves the configuration a request attribute.
// Optionally populates the process group references within the result object.
func (api RequestAttributesAPI) Get(ID string, withProcessGroupReferences bool) (dtapi.RequestAttribute, error) {
	result, _, err := api.client.RequestAttributesApi.GetConfiguration3(nil, ID, &dtapi.GetConfiguration3Opts{
		IncludeProcessGroupReferences: optional.NewBool(withProcessGroupReferences),
	})
	if err != nil {
		return result, err
	}
	return result, nil
}

// All retrieves the configurations of all currently configured request attributes.
func (api RequestAttributesAPI) All() ([]dtapi.RequestAttributeStub, error) {
	result, _, err := api.client.RequestAttributesApi.ListConfigurations2(nil)
	return result.Values, err
}

// ValidateForUpdate allows for validating the configuration of a request attribute
// before invoking the actual Update.
func (api RequestAttributesAPI) ValidateForUpdate(requestAttribute dtapi.RequestAttribute) error {
	_, err := api.client.RequestAttributesApi.ValidateConfiguration5(nil, &dtapi.ValidateConfiguration5Opts{
		RequestAttribute: optional.NewInterface(requestAttribute),
	})
	return err
}

// ValidateForCreate allows for validating the configuration of a request attribute
// before invoking the actual Create.
func (api RequestAttributesAPI) ValidateForCreate(ID string, requestAttribute dtapi.RequestAttribute) error {
	_, err := api.client.RequestAttributesApi.ValidateConfiguration6(nil, ID, &dtapi.ValidateConfiguration6Opts{
		RequestAttribute: optional.NewInterface(requestAttribute),
	})
	return err
}
