package dtapi

import (
	"github.com/antihax/optional"
	dtapi "github.com/dtcookie/dtapi/libdtapienv"
)

// TopologyHostAPI is a convenience wrapper around
// the Services offered via
// "github.com/dtcookie/dtapi/libdtapienv" related to
// Smartscape Topology
type TopologyHostAPI envService

func newTopologyHostAPI(client *dtapi.APIClient) TopologyHostAPI {
	api := TopologyHostAPI{}
	api.client = client
	return api
}

// CreateQuery creates a builder object for conveniently parameterizing a query for hosts
func (api TopologyHostAPI) CreateQuery() GetHostsBuilder {
	return &getHostsBuilder{
		client: api.client,
		opts:   dtapi.GetHostsOpts{},
	}
}

// Get retrieves details about the specified host
func (api TopologyHostAPI) Get(ID string) (dtapi.Host, error) {
	result, _, err := api.client.TopologySmartscapeHostApi.GetSingleHost(nil, ID)
	if err != nil {
		return result, err
	}
	return result, nil
}

// Update specifies the tags for a specific host
func (api TopologyHostAPI) Update(ID string, tags []string) error {
	_, err := api.client.TopologySmartscapeHostApi.UpdateHost(nil, ID, &dtapi.UpdateHostOpts{
		UpdateEntity: optional.NewInterface(dtapi.UpdateEntity{
			Tags: tags,
		}),
	})
	return err
}
