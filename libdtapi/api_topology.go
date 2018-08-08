package dtapi

import (
	"github.com/antihax/optional"
	dtapi "github.com/dtcookie/dtapi/libdtapienv"
)

// TopologyAPI is a convenience wrapper around
// the Services offered via
// "github.com/dtcookie/dtapi/libdtapienv" related to
// the Smartscape API
type TopologyAPI struct {
	envService
	// Applications offers functionality to query for and update Applications
	Applications TopologyApplicationAPI
	// Hosts offers functionality to query for and update Hosts
	Hosts TopologyHostAPI
	// ProcessGroups offers functionality to query for and update ProcessGroups
	ProcessGroups TopologyProcessGroupAPI
	// Processes offers functionality to query for and update Processes
	Processes TopologyProcessesAPI
	// Services offers functionality to query for and update Services
	Services TopologyServiceAPI
}

// CreateCustomDataPoints creates/updates a custom device, or reports metric data points to the custom device.
func (api *TopologyAPI) CreateCustomDataPoints(entityAlias string, pushMessage dtapi.CustomDevicePushMessage) (dtapi.CustomDevicePushResult, error) {
	result, _, err := api.client.TopologySmartscapeCustomDeviceApi.CreateCustomDataPoints(nil, entityAlias, &dtapi.CreateCustomDataPointsOpts{
		CustomDevicePushMessage: optional.NewInterface(pushMessage),
	})
	return result, err
}

func newTopologyAPI(client *dtapi.APIClient) *TopologyAPI {
	api := TopologyAPI{}
	api.client = client
	api.Applications = newTopologyApplicationAPI(client)
	api.Hosts = newTopologyHostAPI(client)
	api.ProcessGroups = newTopologyProcessGroupAPI(client)
	api.Processes = newTopologyProcessesAPI(client)
	api.Services = newTopologyServiceAPI(client)
	return &api
}
