package dynatrace

import (
	"github.com/dtcookie/dtapi/environment"
)

// TopologyAPI TODO: documentation
type TopologyAPI struct {
	service
	Applications  TopologyApplicationAPI
	Hosts         TopologyHostAPI
	ProcessGroups TopologyProcessGroupAPI
	Processes     TopologyProcessesAPI
	Services      TopologyServiceAPI
}

// NewTopologyAPI TODO: documentation
func NewTopologyAPI(srvc service) TopologyAPI {
	topologyAPI := TopologyAPI{
		service:       srvc,
		Applications:  NewTopologyApplicationAPI(srvc),
		Hosts:         NewTopologyHostAPI(srvc),
		ProcessGroups: NewTopologyProcessGroupAPI(srvc),
		Processes:     NewTopologyProcessesAPI(srvc),
		Services:      NewTopologyServiceAPI(srvc),
	}
	return topologyAPI
}

/*
CreateCustomDataPoints Creates/updates a custom device, or reports metric data points to the custom device.
 * @param entityAlias ID of the custom device.   If you use the ID of an existing device, the respective parameters will be updated.
 * @param optional nil or *CreateCustomDataPointsOpts - Optional Parameters:
 * @param "CustomDevicePushMessage" (optional.Interface of CustomDevicePushMessage) -  JSON body of the request containing entity's parameters.
@return CustomDevicePushResult
*/
func (api TopologyAPI) CreateCustomDataPoints(entityAlias string, opts *environment.CreateCustomDataPointsOpts) (environment.CustomDevicePushResult, error) {
	result, _, err := api.client.TopologySmartscapeCustomDeviceApi.CreateCustomDataPoints(nil, entityAlias, opts)
	if err != nil {
		return result, err
	}
	return result, nil
}
