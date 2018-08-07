package dtapi

import (
	"github.com/antihax/optional"
	dtapi "github.com/dtcookie/dtapi/libdtapienv"
)

// EventAPI TODO: documentation
type EventAPI envService

// GetEventByID TODO: documentation
func (api EventAPI) GetEventByID(eventID string) (dtapi.EventRestEntry, error) {
	result, _, err := api.client.EventApi.GetEventById(nil, eventID)
	return result, err
}

// PostNaturalEvent TODO: documentation
func (api EventAPI) PostNaturalEvent(eventPushMessage dtapi.EventPushMessage) ([]string, error) {
	result, _, err := api.client.EventApi.PostNaturalEvent(nil, &dtapi.PostNaturalEventOpts{
		EventPushMessage: optional.NewInterface(eventPushMessage),
	})
	return result.StoredIds, err
}

// CreateQuery TODO: documentation
func (api EventAPI) CreateQuery() *QueryEventsBuilder {
	builder := QueryEventsBuilder{}
	builder.client = api.client
	return &builder
}

// QueryEventsBuilder TODO: documentation
type QueryEventsBuilder struct {
	client *dtapi.APIClient
	opts   *dtapi.QueryEventsOpts
}

// WithStartTimeStamp TODO: documentation
func (builder *QueryEventsBuilder) WithStartTimeStamp(timeStamp int64) *QueryEventsBuilder {
	builder.opts.From = optional.NewInt64(timeStamp)
	return builder
}

// WithEndTimeStamp TODO: documentation
func (builder *QueryEventsBuilder) WithEndTimeStamp(timeStamp int64) *QueryEventsBuilder {
	builder.opts.To = optional.NewInt64(timeStamp)
	return builder
}

// WithRelativeTime TODO: documentation
func (builder *QueryEventsBuilder) WithRelativeTime(relativeTime string) *QueryEventsBuilder {
	builder.opts.RelativeTime = optional.NewString(relativeTime)
	return builder
}

// WithEventType TODO: documentation
func (builder *QueryEventsBuilder) WithEventType(eventType string) *QueryEventsBuilder {
	builder.opts.EventType = optional.NewString(eventType)
	return builder
}

// WithEntityID TODO: documentation
func (builder *QueryEventsBuilder) WithEntityID(entityID string) *QueryEventsBuilder {
	builder.opts.EntityId = optional.NewString(entityID)
	return builder
}

// WithCursor TODO: documentation
func (builder *QueryEventsBuilder) WithCursor(cursor string) *QueryEventsBuilder {
	builder.opts.Cursor = optional.NewString(cursor)
	return builder
}

// QueryEvents TODO: documentation
func (builder *QueryEventsBuilder) QueryEvents() (dtapi.EventQueryResult, error) {
	result, _, err := builder.client.EventApi.QueryEvents(nil, builder.opts)
	return result, err
}
