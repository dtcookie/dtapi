package dtapi

import (
	dtapi "github.com/dtcookie/dtapi/libdtapienv"
)

// SyntheticAPI TODO: documentation
type SyntheticAPI envService

/*
PushEvents Push information about external Synthetic Events
 * @param externalSyntheticEvents
*/
func (api SyntheticAPI) PushEvents(externalSyntheticEvents dtapi.ExternalSyntheticEvents) error {
	_, err := api.client.SyntheticApi.PushEvents(nil, externalSyntheticEvents)
	return err
}

/*
PushStateModification Modify the operation state of all external monitors
 * @param optional nil or *PushStateModificationOpts - Optional Parameters:
 * @param "StateModification" (optional.Interface of StateModification) -
@return string
*/
func (api SyntheticAPI) PushStateModification(opts dtapi.StateModification) (string, error) {
	result, _, err := api.client.SyntheticApi.PushStateModification(nil, opts)
	return result, err
}

/*
TestResults Push Information about Synthetic Tests, Locations and Test Results.
 * @param externalSyntheticTests Information about Synthetic Tests, Locations and Test Results.
@return string
*/
func (api SyntheticAPI) TestResults(externalSyntheticTests dtapi.ExternalSyntheticTests) (string, error) {
	result, _, err := api.client.SyntheticApi.TestResults(nil, externalSyntheticTests)
	return result, err
}
