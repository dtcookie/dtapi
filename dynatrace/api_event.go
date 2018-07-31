package dynatrace

import (
	"github.com/dtcookie/dtapi/environment"
)

// EventAPI TODO: documentation
type EventAPI service

// GetEventByID TODO: documentation
func (api EventAPI) GetEventByID(eventID string) (environment.EventRestEntry, error) {
	result, _, err := api.client.EventApi.GetEventById(nil, eventID)
	if err != nil {
		return result, err
	}
	return result, nil
}

// PostNaturalEvent TODO: documentation
func (api EventAPI) PostNaturalEvent(opts *environment.PostNaturalEventOpts) ([]string, error) {
	result, _, err := api.client.EventApi.PostNaturalEvent(nil, opts)
	if err != nil {
		return result.StoredIds, err
	}
	return result.StoredIds, nil
}

// QueryEvents TODO: documentation
func (api EventAPI) QueryEvents(opts *environment.QueryEventsOpts) (environment.EventQueryResult, error) {
	result, _, err := api.client.EventApi.QueryEvents(nil, opts)
	if err != nil {
		return result, err
	}
	return result, nil
}
