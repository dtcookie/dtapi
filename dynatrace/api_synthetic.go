package dynatrace

import (
	"github.com/dtcookie/dtapi/environment"
)

// SyntheticAPI TODO: documentation
type SyntheticAPI service

/*
PushEvents Push information about external Synthetic Events
 * @param externalSyntheticEvents
*/
func (api SyntheticAPI) PushEvents(externalSyntheticEvents environment.ExternalSyntheticEvents) error {
	_, err := api.client.SyntheticApi.PushEvents(nil, externalSyntheticEvents)
	if err != nil {
		return err
	}
	return nil
}

/*
PushStateModification Modify the operation state of all external monitors
 * @param optional nil or *PushStateModificationOpts - Optional Parameters:
 * @param "StateModification" (optional.Interface of StateModification) -
@return string
*/
func (api SyntheticAPI) PushStateModification(opts *environment.PushStateModificationOpts) (string, error) {
	result, _, err := api.client.SyntheticApi.PushStateModification(nil, opts)
	if err != nil {
		return result, err
	}
	return result, nil
}

/*
TestResults Push Information about Synthetic Tests, Locations and Test Results.
 * @param externalSyntheticTests Information about Synthetic Tests, Locations and Test Results.
@return string
*/
func (api SyntheticAPI) TestResults(externalSyntheticTests environment.ExternalSyntheticTests) (string, error) {
	result, _, err := api.client.SyntheticApi.TestResults(nil, externalSyntheticTests)
	if err != nil {
		return result, err
	}
	return result, nil
}
