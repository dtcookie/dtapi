package dtapi

import (
	"github.com/antihax/optional"
	dtapi "github.com/dtcookie/dtapi/libdtapienv"
)

// TopologyAPI TODO: documentation
type TopologyAPI struct {
	envService
	Applications  TopologyApplicationAPI
	Hosts         TopologyHostAPI
	ProcessGroups TopologyProcessGroupAPI
	Processes     TopologyProcessesAPI
	Services      TopologyServiceAPI
}

// NewTopologyAPI TODO: documentation
func NewTopologyAPI(client *dtapi.APIClient) TopologyAPI {
	api := TopologyAPI{}
	api.client = client
	api.Applications = NewTopologyApplicationAPI(client)
	api.Hosts = NewTopologyHostAPI(client)
	api.ProcessGroups = NewTopologyProcessGroupAPI(client)
	api.Processes = NewTopologyProcessesAPI(client)
	api.Services = NewTopologyServiceAPI(client)
	return api
}

/*
CreateCustomDataPoints Creates/updates a custom device, or reports metric data points to the custom device.
 * @param entityAlias ID of the custom device.   If you use the ID of an existing device, the respective parameters will be updated.
 * @param optional nil or *CreateCustomDataPointsOpts - Optional Parameters:
 * @param "CustomDevicePushMessage" (optional.Interface of CustomDevicePushMessage) -  JSON body of the request containing entity's parameters.
@return CustomDevicePushResult
*/
func (api TopologyAPI) CreateCustomDataPoints(entityAlias string, pushMessage dtapi.CustomDevicePushMessage) (dtapi.CustomDevicePushResult, error) {
	result, _, err := api.client.TopologySmartscapeCustomDeviceApi.CreateCustomDataPoints(nil, entityAlias, &dtapi.CreateCustomDataPointsOpts{
		CustomDevicePushMessage: optional.NewInterface(pushMessage),
	})
	return result, err
}
