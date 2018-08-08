package dtapi

import (
	"github.com/antihax/optional"
	dtapi "github.com/dtcookie/dtapi/libdtapienv"
)

// TopologyServiceAPI is a convenience wrapper around
// the Services offered via
// "github.com/dtcookie/dtapi/libdtapienv" related to
// Smartscape Topology
type TopologyServiceAPI envService

// CreateQuery creates a builder object for conveniently parameterizing a query for Services
func (api TopologyServiceAPI) CreateQuery() GetServiceBuilder {
	return &getServiceBuilder{
		client: api.client,
		opts:   dtapi.GetServicesOpts{},
	}
}

// Get retrieves parameters of the specified Service
func (api TopologyServiceAPI) Get(entityID string) (dtapi.Service, error) {
	result, _, err := api.client.TopologySmartscapeServiceApi.GetSingleService(nil, entityID)
	return result, err
}

// Update Updates properties of the specified Service
func (api TopologyServiceAPI) Update(entityID string, tags []string) error {
	_, err := api.client.TopologySmartscapeServiceApi.UpdateService(nil, entityID, &dtapi.UpdateServiceOpts{
		UpdateEntity: optional.NewInterface(dtapi.UpdateEntity{
			Tags: tags,
		}),
	})
	return err
}

func newTopologyServiceAPI(client *dtapi.APIClient) TopologyServiceAPI {
	api := TopologyServiceAPI{}
	api.client = client
	return api
}
