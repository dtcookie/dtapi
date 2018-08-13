package dtapi

import (
	"github.com/antihax/optional"
	dtapi "github.com/dtcookie/dtapi/libdtapiconf"
)

// AutoTagsAPI is a convenience wrapper around
// the Services offered via
// "github.com/dtcookie/dtapi/libdtapiconf" related to
// configuring Auto Tagging
type AutoTagsAPI confService

// AutoTagAPI provides functionality for one specific
// Auto Tag
type AutoTagAPI struct {
	confService
	// the unique identifier of this custom service
	ID string
}

func newAutoTagAPI(ID string, client *dtapi.APIClient) *AutoTagAPI {
	api := AutoTagAPI{ID: ID}
	api.client = client
	return &api
}

// ForID provides an API for configuring a specific
// already existing Auto Tag Configuration, such as deleting
// or updating it.
func (api *AutoTagsAPI) ForID(ID string) *AutoTagAPI {
	return newAutoTagAPI(ID, api.client)
}

// Create TODO: documentation
func (api *AutoTagsAPI) Create(autoTag dtapi.AutoTag) (dtapi.AutoTagStub, error) {
	var err error
	var autoTagStub dtapi.AutoTagStub

	autoTagStub, _, err = api.client.AutomaticallyAppliedTagsApi.CreateAutoTag1(nil, &dtapi.CreateAutoTag1Opts{
		AutoTag: optional.NewInterface(autoTag),
	})
	return autoTagStub, err
}

// Update updates or creates (if it doesn't exist yet) and Auto Tag
func (api *AutoTagAPI) Update(autoTag dtapi.AutoTag) error {
	_, err := api.client.AutomaticallyAppliedTagsApi.CreateOrUpdateAutoTag1(nil, api.ID, &dtapi.CreateOrUpdateAutoTag1Opts{
		AutoTag: optional.NewInterface(autoTag),
	})
	return err
}

// Delete deletes the Auto Tag Configuration
func (api *AutoTagAPI) Delete() error {
	_, err := api.client.AutomaticallyAppliedTagsApi.DeleteAutoTag1(nil, api.ID)
	return err

}

// GetAll retrieves ID/name pairs for all currently configure Auto Tag Configurations
func (api *AutoTagsAPI) GetAll() ([]dtapi.AutoTagStub, error) {
	var err error
	var autoTags dtapi.AutoTagStubList
	autoTags, _, err = api.client.AutomaticallyAppliedTagsApi.GetAllAutoTagConfigs1(nil)
	return autoTags.Values, err
}

// Get retrieves the detailed configuration of the Auto Tag
func (api *AutoTagAPI) Get(includeProcessGroupRefs bool) (dtapi.AutoTag, error) {
	var err error
	var autoTag dtapi.AutoTag
	autoTag, _, err = api.client.AutomaticallyAppliedTagsApi.GetSingleAutoTagConfig1(nil, api.ID, &dtapi.GetSingleAutoTagConfig1Opts{
		IncludeProcessGroupReferences: optional.NewBool(includeProcessGroupRefs),
	})
	return autoTag, err
}

// Validate allows for validating an auto tag configuration before
// sending the actual Update request
func (api *AutoTagAPI) Validate(autoTag dtapi.AutoTag) error {
	_, err := api.client.AutomaticallyAppliedTagsApi.ValidateAutoTag1(nil, api.ID, &dtapi.ValidateAutoTag1Opts{
		AutoTag: optional.NewInterface(autoTag),
	})
	return err
}

// Validate allows for validating an auto tag configuration before
// sending the actual Create request
func (api *AutoTagsAPI) Validate(autoTag dtapi.AutoTag) error {
	_, err := api.client.AutomaticallyAppliedTagsApi.ValidateAutoTag2(nil, &dtapi.ValidateAutoTag2Opts{
		AutoTag: optional.NewInterface(autoTag),
	})
	return err
}
