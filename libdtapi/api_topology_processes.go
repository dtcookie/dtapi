package dtapi

import (
	dtapi "github.com/dtcookie/dtapi/libdtapienv"
)

// TopologyProcessesAPI is a convenience wrapper around
// the Services offered via
// "github.com/dtcookie/dtapi/libdtapienv" related to
// Smartscape Topology
type TopologyProcessesAPI envService

// CreateQuery creates a builder object for conveniently parameterizing a query for process group instances
func (api TopologyProcessesAPI) CreateQuery() GetProcessesBuilder {
	return &getProcessesBuilder{
		client: api.client,
		opts:   dtapi.GetProcessesOpts{},
	}
}

// Get Gets parameters of the specified process group instance
func (api TopologyProcessesAPI) Get(entityID string) (dtapi.ProcessGroupInstance, error) {
	result, _, err := api.client.TopologySmartscapeProcessApi.GetSingleProcess(nil, entityID)
	return result, err
}

func newTopologyProcessesAPI(client *dtapi.APIClient) TopologyProcessesAPI {
	api := TopologyProcessesAPI{}
	api.client = client
	return api
}
