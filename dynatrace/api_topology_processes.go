package dynatrace

import (
	"github.com/dtcookie/dtapi/environment"
)

// TopologyProcessesAPI TODO: documentation
type TopologyProcessesAPI struct {
	service
}

// NewTopologyProcessesAPI TODO: documentation
func NewTopologyProcessesAPI(srvc service) TopologyProcessesAPI {
	return TopologyProcessesAPI{
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
func (api TopologyProcessesAPI) FetchAll() GetProcessesBuilder {
	return GetProcessesBuilder{
		client: api.client,
		opts:   environment.GetProcessesOpts{},
	}
}

/*
Fetch Gets parameters of the specified host
 * @param entityID Dynatrace entity ID of the host you're inquiring.   You can find it in the URL of the corresponding host page, for example, `HOST-007`.
@return Host
*/
func (api TopologyProcessesAPI) Fetch(entityID string) (environment.ProcessGroupInstance, error) {
	result, _, err := api.client.TopologySmartscapeProcessApi.GetSingleProcess(nil, entityID)
	if err != nil {
		return result, err
	}
	return result, nil
}
