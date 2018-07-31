package dynatrace

import (
	"github.com/antihax/optional"
	"github.com/dtcookie/dtapi/environment"
)

// TopologyServiceAPI TODO: documentation
type TopologyServiceAPI struct {
	service
}

// NewTopologyServiceAPI TODO: documentation
func NewTopologyServiceAPI(srvc service) TopologyServiceAPI {
	return TopologyServiceAPI{
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
func (api TopologyServiceAPI) FetchAll() GetServiceBuilder {
	return GetServiceBuilder{
		client: api.client,
		opts:   environment.GetServicesOpts{},
	}
}

/*
Fetch Gets parameters of the specified host
 * @param entityID Dynatrace entity ID of the host you're inquiring.   You can find it in the URL of the corresponding host page, for example, `HOST-007`.
@return Host
*/
func (api TopologyServiceAPI) Fetch(entityID string) (environment.Service, error) {
	result, _, err := api.client.TopologySmartscapeServiceApi.GetSingleService(nil, entityID)
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
func (api TopologyServiceAPI) Update(entityID string, tags []string) error {
	_, err := api.client.TopologySmartscapeServiceApi.UpdateService(nil, entityID, &environment.UpdateServiceOpts{
		UpdateEntity: optional.NewInterface(environment.UpdateEntity{
			Tags: tags,
		}),
	})
	return err
}
