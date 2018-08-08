package dtapi

import (
	dtapi "github.com/dtcookie/dtapi/libdtapienv"
)

// SyntheticAPI is a convenience wrapper around
// the Services offered via
// "github.com/dtcookie/dtapi/libdtapienv" related to
// configuring the Synthetic API
type SyntheticAPI envService

// PushEvents pushes updates on open events to Dynatrace or closes previously open events.
func (api SyntheticAPI) PushEvents(externalSyntheticEvents dtapi.ExternalSyntheticEvents) error {
	_, err := api.client.SyntheticApi.PushEvents(nil, externalSyntheticEvents)
	return err
}

// PushStateModification modifies the operation state of all external monitors
func (api SyntheticAPI) PushStateModification(state string) (string, error) {
	result, _, err := api.client.SyntheticApi.PushStateModification(nil, dtapi.StateModification{
		State: state,
	})
	return result, err
}

// TestResults pushes information about external Synthetic tests, locations, and test results to Dynatrace.
func (api SyntheticAPI) TestResults(externalSyntheticTests dtapi.ExternalSyntheticTests) (string, error) {
	result, _, err := api.client.SyntheticApi.TestResults(nil, externalSyntheticTests)
	return result, err
}
