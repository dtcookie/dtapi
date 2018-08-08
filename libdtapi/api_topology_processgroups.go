package dtapi

import (
	"github.com/antihax/optional"
	dtapi "github.com/dtcookie/dtapi/libdtapienv"
)

// TopologyProcessGroupAPI is a convenience wrapper around
// the Services offered via
// "github.com/dtcookie/dtapi/libdtapienv" related to
// Smartscape Topology
type TopologyProcessGroupAPI envService

// CreateQuery creates a builder object for conveniently parameterizing a query for process groups
func (api TopologyProcessGroupAPI) CreateQuery() GetProcessGroupsBuilder {
	return &getProcessGroupsBuilder{
		client: api.client,
		opts:   dtapi.GetProcessGroupsOpts{},
	}
}

// Get Gets parameters of the specified process group
func (api TopologyProcessGroupAPI) Get(entityID string) (dtapi.ProcessGroup, error) {
	result, _, err := api.client.TopologySmartscapeProcessGroupApi.GetSingleProcessGroup(nil, entityID)
	return result, err
}

// Update Updates properties of the specified process group
func (api TopologyProcessGroupAPI) Update(entityID string, tags []string) error {
	_, err := api.client.TopologySmartscapeProcessGroupApi.UpdateProcessGroup(nil, entityID, &dtapi.UpdateProcessGroupOpts{
		UpdateEntity: optional.NewInterface(dtapi.UpdateEntity{
			Tags: tags,
		}),
	})
	return err
}

func newTopologyProcessGroupAPI(client *dtapi.APIClient) TopologyProcessGroupAPI {
	api := TopologyProcessGroupAPI{}
	api.client = client
	return api
}
