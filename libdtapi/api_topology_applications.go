package dtapi

import (
	"github.com/antihax/optional"
	dtapi "github.com/dtcookie/dtapi/libdtapienv"
)

// TopologyApplicationAPI TODO: documentation
type TopologyApplicationAPI envService

// NewTopologyApplicationAPI TODO: documentation
func NewTopologyApplicationAPI(client *dtapi.APIClient) TopologyApplicationAPI {
	api := TopologyApplicationAPI{}
	api.client = client
	return api
}

/*
CreateQuery Gets the list of all applications in your environment along with their parameters
You can optionally specify timeframe, to filter the output only to applications, active in specified time.
 * @param optional nil or *GetApplicationsOpts - Optional Parameters:
 * @param "StartTimestamp" (optional.Int64) -  Start timestamp of the requested timeframe, in milliseconds (UTC).   If no timeframe specified the 72 hours behind from now is used.
 * @param "EndTimestamp" (optional.Int64) -  End timestamp of the requested timeframe, in milliseconds (UTC).   If no timeframe specified then now is used.
 * @param "RelativeTime" (optional.String) -  Relative timeframe, back from now.
 * @param "Tag" (optional.Interface of []string) -  Filters the resulting set of applications by the specified tag.    An application has to match ALL specified tags.
 * @param "Entity" (optional.Interface of []string) -  Only return specified applications.
@return []Application
*/
func (api TopologyApplicationAPI) CreateQuery() *GetApplicationsBuilder {
	return &GetApplicationsBuilder{
		client: api.client,
		opts:   dtapi.GetApplicationsOpts{},
	}
}

/*
GET Gets parameters of the specified application
 * @param entityID Dynatrace entity ID of the application you're inquiring.   You can find them in the URL of the corresponding application page, for example, `APPLICATION-007`.
@return Application
*/
func (api TopologyApplicationAPI) GET(entityID string) (dtapi.Application, error) {
	result, _, err := api.client.TopologySmartscapeApplicationApi.GetSingleApplication(nil, entityID)
	return result, err
}

/*
UPDATE Updates parameters of the specified application
 * @param entityID Dynatrace entity ID of the application to be updated.   You can find them in the URL of the corresponding application page, for example, `APPLICATION-007`.
 * @param tags []string
*/
func (api TopologyApplicationAPI) UPDATE(entityID string, tags []string) error {
	_, err := api.client.TopologySmartscapeApplicationApi.UpdateApplication(nil, entityID, &dtapi.UpdateApplicationOpts{
		UpdateEntity: optional.NewInterface(dtapi.UpdateEntity{
			Tags: tags,
		}),
	})
	return err
}
