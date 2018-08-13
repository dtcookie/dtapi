package dtapi

import (
	"github.com/antihax/optional"
	dtapi "github.com/dtcookie/dtapi/libdtapienv"
)

// TopologyApplicationAPI is a convenience wrapper around
// the Services offered via
// "github.com/dtcookie/dtapi/libdtapienv" related to
// Smartscape Topology
type TopologyApplicationAPI envService

// NewTopologyApplicationAPI TODO: documentation
func newTopologyApplicationAPI(client *dtapi.APIClient) TopologyApplicationAPI {
	api := TopologyApplicationAPI{}
	api.client = client
	return api
}

// CreateQuery creates a builder object to conveniently parameterize a query for applications.
func (api TopologyApplicationAPI) CreateQuery() GetApplicationsBuilder {
	return &getApplicationsBuilder{
		client: api.client,
		opts:   dtapi.GetApplicationsOpts{},
	}
}

// Get Gets parameters of the specified application
func (api TopologyApplicationAPI) Get(ID string) (dtapi.Application, error) {
	result, _, err := api.client.TopologySmartscapeApplicationApi.GetSingleApplication(nil, ID)
	return result, err
}

// Update updates parameters of the specified application
func (api TopologyApplicationAPI) Update(ID string, tags ...string) error {
	_, err := api.client.TopologySmartscapeApplicationApi.UpdateApplication(nil, ID, &dtapi.UpdateApplicationOpts{
		UpdateEntity: optional.NewInterface(dtapi.UpdateEntity{
			Tags: tags,
		}),
	})
	return err
}
