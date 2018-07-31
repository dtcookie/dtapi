package dynatrace

import (
	"github.com/antihax/optional"
	"github.com/dtcookie/dtapi/environment"
)

// TopologyApplicationAPI TODO: documentation
type TopologyApplicationAPI struct {
	service
}

// NewTopologyApplicationAPI TODO: documentation
func NewTopologyApplicationAPI(srvc service) TopologyApplicationAPI {
	return TopologyApplicationAPI{
		srvc,
	}
}

/*
FetchAll Gets the list of all applications in your environment along with their parameters
You can optionally specify timeframe, to filter the output only to applications, active in specified time.
 * @param optional nil or *GetApplicationsOpts - Optional Parameters:
 * @param "StartTimestamp" (optional.Int64) -  Start timestamp of the requested timeframe, in milliseconds (UTC).   If no timeframe specified the 72 hours behind from now is used.
 * @param "EndTimestamp" (optional.Int64) -  End timestamp of the requested timeframe, in milliseconds (UTC).   If no timeframe specified then now is used.
 * @param "RelativeTime" (optional.String) -  Relative timeframe, back from now.
 * @param "Tag" (optional.Interface of []string) -  Filters the resulting set of applications by the specified tag.    An application has to match ALL specified tags.
 * @param "Entity" (optional.Interface of []string) -  Only return specified applications.
@return []Application
*/
func (api TopologyApplicationAPI) FetchAll() GetApplicationsBuilder {
	return GetApplicationsBuilder{
		client: api.client,
		opts:   environment.GetApplicationsOpts{},
	}
}

/*
FetchByID Gets parameters of the specified application
 * @param entityID Dynatrace entity ID of the application you're inquiring.   You can find them in the URL of the corresponding application page, for example, `APPLICATION-007`.
@return Application
*/
func (api TopologyApplicationAPI) FetchByID(entityID string) (environment.Application, error) {
	result, _, err := api.client.TopologySmartscapeApplicationApi.GetSingleApplication(nil, entityID)
	return result, err
}

/*
Update Updates parameters of the specified application
 * @param entityID Dynatrace entity ID of the application to be updated.   You can find them in the URL of the corresponding application page, for example, `APPLICATION-007`.
 * @param tags []string
*/
func (api TopologyApplicationAPI) Update(entityID string, tags []string) error {
	_, err := api.client.TopologySmartscapeApplicationApi.UpdateApplication(nil, entityID, &environment.UpdateApplicationOpts{
		UpdateEntity: optional.NewInterface(environment.UpdateEntity{
			Tags: tags,
		}),
	})
	return err
}
