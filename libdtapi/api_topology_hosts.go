package dtapi

import (
	"github.com/antihax/optional"
	dtapi "github.com/dtcookie/dtapi/libdtapienv"
)

// TopologyHostAPI TODO: documentation
type TopologyHostAPI envService

// NewTopologyHostAPI TODO: documentation
func NewTopologyHostAPI(client *dtapi.APIClient) TopologyHostAPI {
	api := TopologyHostAPI{}
	api.client = client
	return api
}

/*
CreateQuery Lists all available hosts in your environment
You can narrow down the output by specifying filtering parameters of the request.
 * @param optional nil or *GetHostsOpts - Optional Parameters:
 * @param "StartTimestamp" (optional.Int64) -  Start timestamp of the requested timeframe, in milliseconds (UTC).   If not set, then 72 hours behind from now is used.
 * @param "EndTimestamp" (optional.Int64) -  Start timestamp of the requested timeframe, in milliseconds (UTC).   If not set, then now is used.
 * @param "RelativeTime" (optional.String) -  Relative timeframe, back from now.
 * @param "Tag" (optional.Interface of []string) -  Filters the resulting set of hosts by the specified tag.    A host has to match ALL specified tags.
 * @param "ShowMonitoringCandidates" (optional.Bool) -  Include/exclude monitoring canditate to the response.   Monitoring candidates are network entities, which are detected but not monitored.
 * @param "Entity" (optional.Interface of []string) -  Only return specified hosts.
@return []Host
*/
func (api TopologyHostAPI) CreateQuery() *GetHostsBuilder {
	return &GetHostsBuilder{
		client: api.client,
		opts:   dtapi.GetHostsOpts{},
	}
}

/*
GET Gets parameters of the specified host
 * @param entityID Dynatrace entity ID of the host you're inquiring.   You can find it in the URL of the corresponding host page, for example, `HOST-007`.
@return Host
*/
func (api TopologyHostAPI) GET(entityID string) (dtapi.Host, error) {
	result, _, err := api.client.TopologySmartscapeHostApi.GetSingleHost(nil, entityID)
	if err != nil {
		return result, err
	}
	return result, nil
}

/*
Update Updates properties of the specified host
 * @param entityID Dynatrace entity ID of the host to be updated.   You can find it in the URL of the corresponding host page, for example, `HOST-007`.
 * @param optional nil or *UpdateHostOpts - Optional Parameters:
 * @param "UpdateEntity" (optional.Interface of UpdateEntity) -
*/
func (api TopologyHostAPI) Update(entityID string, tags []string) error {
	_, err := api.client.TopologySmartscapeHostApi.UpdateHost(nil, entityID, &dtapi.UpdateHostOpts{
		UpdateEntity: optional.NewInterface(dtapi.UpdateEntity{
			Tags: tags,
		}),
	})
	return err
}