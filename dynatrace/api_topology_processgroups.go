package dynatrace

import (
	"github.com/antihax/optional"
	"github.com/dtcookie/dtapi/environment"
)

// TopologyProcessGroupAPI TODO: documentation
type TopologyProcessGroupAPI struct {
	service
}

// NewTopologyProcessGroupAPI TODO: documentation
func NewTopologyProcessGroupAPI(srvc service) TopologyProcessGroupAPI {
	return TopologyProcessGroupAPI{
		srvc,
	}
}

/*
FetchAll Lists all available hosts in your environment
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
func (api TopologyProcessGroupAPI) FetchAll() GetProcessGroupsBuilder {
	return GetProcessGroupsBuilder{
		client: api.client,
		opts:   environment.GetProcessGroupsOpts{},
	}
}

/*
Fetch Gets parameters of the specified host
 * @param entityID Dynatrace entity ID of the host you're inquiring.   You can find it in the URL of the corresponding host page, for example, `HOST-007`.
@return Host
*/
func (api TopologyProcessGroupAPI) Fetch(entityID string) (environment.ProcessGroup, error) {
	result, _, err := api.client.TopologySmartscapeProcessGroupApi.GetSingleProcessGroup(nil, entityID)
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
func (api TopologyProcessGroupAPI) Update(entityID string, tags []string) error {
	_, err := api.client.TopologySmartscapeProcessGroupApi.UpdateProcessGroup(nil, entityID, &environment.UpdateProcessGroupOpts{
		UpdateEntity: optional.NewInterface(environment.UpdateEntity{
			Tags: tags,
		}),
	})
	return err
}
