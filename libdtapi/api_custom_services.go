package dtapi

import (
	"github.com/antihax/optional"
	dtapi "github.com/dtcookie/dtapi/libdtapiconf"
)

// CustomServicesAPI is a convenience wrapper around
// the Services offered via
// "github.com/dtcookie/dtapi/libdtapiconf" related to
// Custom Service Configuration.
//
// This API does not offer functionaliy directly, but instead
// allows for access to technology specific APIs through
// its members 'Java', 'DotNet' and 'PHP'.
type CustomServicesAPI struct {
	// Java offers Custom Services Configuration for
	// for Java
	Java *techCustomServicesAPI
	// DotNet offers Custom Services Configuration for
	// for DotNet
	DotNet *techCustomServicesAPI
	// PHP offers Custom Services Configuration for
	// for PHP
	PHP *techCustomServicesAPI
}

// NewCustomServicesAPI creates and initializes an object of type
// CustomServicesAPI.
// The properties 'Java', 'DotNet' and 'PHP' will be populated
// with working APIs.
func NewCustomServicesAPI(client *dtapi.APIClient) CustomServicesAPI {
	return CustomServicesAPI{
		Java:   newTechCustomServicesAPI("java", client),
		DotNet: newTechCustomServicesAPI("dotNet", client),
		PHP:    newTechCustomServicesAPI("php", client),
	}
}

// techCustomServicesAPI offers functionality
// for Custom Service Configuration restricted
// to a specific technology type.
//
// It's not necessary to create objects of that type
// manually.
// Instead just use either one of
// 	- CustomServicesAPI.Java
// 	- CustomServicesAPI.DotNet
// 	- CustomServicesAPI.PHP
//
// The property 'All' offers functionality
// for batch processing restricted to that
// specific technology type.
type techCustomServicesAPI struct {
	confService
	technology string
	All        *techAllCustomServicesAPI
}

func newTechCustomServicesAPI(technology string, client *dtapi.APIClient) *techCustomServicesAPI {
	api := techCustomServicesAPI{technology: technology}
	api.client = client
	api.All = newTechAllCustomServicesAPI(technology, client)
	return &api
}

// techAllCustomServicesAPI provides functionality
// for batch processing Custom Service Configuration
// restricted to a specific technology type.
type techAllCustomServicesAPI struct {
	confService
	technology string
}

func newTechAllCustomServicesAPI(technology string, client *dtapi.APIClient) *techAllCustomServicesAPI {
	api := techAllCustomServicesAPI{technology: technology}
	api.client = client
	return &api
}

// customServiceAPI offers functionality for
// configuring a specific Custom Service
// restricted down not only to a technology
// but the actual unique identifier of this
// Custom Service Config
type customServiceAPI struct {
	confService
	technology string
	// the unique identifier of this custom service
	ID string
}

func newCustomServiceAPI(technology string, ID string, client *dtapi.APIClient) *customServiceAPI {
	api := customServiceAPI{technology: technology, ID: ID}
	api.client = client
	return &api
}

// ForID provides an API for configuring a specific
// already existing Custom Service, such as deleting
// or updating it.
func (api *techCustomServicesAPI) ForID(ID string) *customServiceAPI {
	return newCustomServiceAPI(api.technology, ID, api.client)
}

// Delete deletes the custom service
func (api *customServiceAPI) Delete() error {
	_, err := api.client.CustomServicesApi.Delete(nil, api.technology, api.ID)
	return err
}

// Get retrieves the details of the configuration of this Custom Service
// Parameters:
// 	- withProcessGroupReferences	if 'true' the resulting 'CustomService' object
//									will be populated with references to process
//									process groups.
func (api *customServiceAPI) Get(withProcessGroupReferences bool) (dtapi.CustomService, error) {
	result, _, err := api.client.CustomServicesApi.GetItem(nil, api.technology, api.ID, &dtapi.GetItemOpts{
		IncludeProcessGroupReferences: optional.NewBool(withProcessGroupReferences),
	})
	return result, err
}

// Update updates the configuration of the Custom Service
//
// In case the ID of the Custom Service encoded into its configuration
// doesn't match this APIs ID, the update will fail.
//
// TODO: Perform ID check before sending the request.
func (api *customServiceAPI) Update(c dtapi.CustomService) error {
	_, err := api.client.CustomServicesApi.PutItem(nil, api.technology, api.ID, &dtapi.PutItemOpts{
		CustomService: optional.NewInterface(c),
	})
	return err
}

// Get retrieves a list of all currently configured Custom Service Configurations for
// the technology type of this specific API.
//
// Process Group References will be populated in the response optionally.
func (api *techAllCustomServicesAPI) Get(withProcessGroupReferences bool) ([]dtapi.CustomService, error) {
	result, _, err := api.client.CustomServicesApi.GetList(nil, api.technology, &dtapi.GetListOpts{
		IncludeProcessGroupReferences: optional.NewBool(withProcessGroupReferences),
	})
	return result.CustomServices, err
}

// Update updates all the Custom Service configurations for the
// technology type specific to this API.
//
// NOTE:	This update affects ALL the Custom Services for this technology.
//			A Custom Service not present in the specified service list is going to get deleted!
func (api *techAllCustomServicesAPI) Update(services []dtapi.CustomService) error {
	_, err := api.client.CustomServicesApi.PutList(nil, api.technology, &dtapi.PutListOpts{
		CustomServiceList: optional.NewInterface(dtapi.CustomServiceList{
			CustomServices: services,
		}),
	})
	return err
}

// Validate allows for validating the configuration before invoking the actual
// update.
//
// NOTE:	Similar to the function 'Update' this validation is also assuming
//			that the specified service list covers ALL the services that are supposed
//			to be configured for the technology type this API is related to.
func (api *techAllCustomServicesAPI) Validate(services []dtapi.CustomService) error {
	_, err := api.client.CustomServicesApi.ValidateList(nil, api.technology, &dtapi.ValidateListOpts{
		CustomServiceList: optional.NewInterface(dtapi.CustomServiceList{
			CustomServices: services,
		}),
	})
	return err
}

// Create creates (HTTP Post) a new Custom Service for the technology type
// this API is related to.
//
// NOTE:	You must not specify the IDs for the custom service or any of its rules.
//
// TODO:	Perform the check for populated IDs offline before sending the request.
func (api techCustomServicesAPI) Create(c dtapi.CustomService) error {
	_, err := api.client.CustomServicesApi.Post(nil, api.technology, &dtapi.PostOpts{
		CustomService: optional.NewInterface(c),
	})
	return err
}

// Validate allows for validating the correctness of a Custom Service before
// invoking the actual Create
func (api techCustomServicesAPI) Validate(service dtapi.CustomService) error {
	_, err := api.client.CustomServicesApi.ValidateItem(nil, api.technology, &dtapi.ValidateItemOpts{
		CustomService: optional.NewInterface(service),
	})
	return err
}
