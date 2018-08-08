package dtapi

import (
	"github.com/antihax/optional"
	dtapi "github.com/dtcookie/dtapi/libdtapienv"
)

// EventAPI is a convenience wrapper around
// the Services offered via
// "github.com/dtcookie/dtapi/libdtapienv" related to
// pushin Events into Dynatrace.
//
// The events REST endpoint enables 3rd party integrations to push custom events
// to one or more monitored entities via the API. The intent of this interface is
// to allow 3rd party systems, such as CI platforms (Jenkins, Bamboo, Electric Cloud, etc.)
// to provide additional detail for Dynatrace automated root cause analysis.
//
// The events API offers a set of semantically predefined event types that allow
// the Dynatrace problem correlation engine to correctly handle information provided
// by external systems.
// The predefined semantics of these event types allows for more precise root cause detection.
type EventAPI envService

// Get retrieves the details for an Event
func (api *EventAPI) Get(eventID string) (dtapi.EventRestEntry, error) {
	result, _, err := api.client.EventApi.GetEventById(nil, eventID)
	return result, err
}

// PostNaturalEvent pushes Custom Events to one or more monitored entities.
// Returns the IDs of the Events created by this push message.
func (api *EventAPI) PostNaturalEvent(eventPushMessage dtapi.EventPushMessage) ([]string, error) {
	result, _, err := api.client.EventApi.PostNaturalEvent(nil, &dtapi.PostNaturalEventOpts{
		EventPushMessage: optional.NewInterface(eventPushMessage),
	})
	return result.StoredIds, err
}

// QueryEvents queries for events without any specific parameterization.
func (api *EventAPI) QueryEvents() (dtapi.EventQueryResult, error) {
	return api.CreateQuery().QueryEvents()
}

// QueryEventsBuilder is a convenience object to set
// optional parameters for an even query.
type QueryEventsBuilder interface {
	// WithStartTimeStamp specifies the end timestamp in milliseconds since Unix epoch
	WithStartTimeStamp(timeStamp int64) QueryEventsBuilder
	// WithEndTimeStamp specifies the end timestamp in milliseconds since Unix epoch.
	// If not specified, the default timeframe is the last 30 days
	WithEndTimeStamp(timeStamp int64) QueryEventsBuilder
	// WithRelativeTime is the relative timeframe, back from the current time.
	WithRelativeTime(relativeTime string) QueryEventsBuilder
	// WithEventType filters the event feed based on a specific event type
	WithEventType(eventType string) QueryEventsBuilder
	// WithEntityID if specified, it means to only receive events for a given monitored entity, such as a host, process, or service
	WithEntityID(entityID string) QueryEventsBuilder
	// WithCursor If a query returns a cursor string this string can be used to fetch the next 150 events of a query.
	// Note that there is no need to specify additional parameters (for instance eventType, from or to) as the cursor string already contains all these parameters
	WithCursor(cursor string) QueryEventsBuilder
	// QueryEvents queries for events based on the parameterization previously defined
	// via 'WithStartTimeStamp', 'WithEndTimeStamp, 'WithRelativeTime', 'WithEventType', 'WithEntityID' and 'WithCursor'.
	QueryEvents() (dtapi.EventQueryResult, error)
}

// queryEventsBuilder is a convenience object to set
// optional parameters for an even query.
type queryEventsBuilder struct {
	QueryEventsBuilder
	client *dtapi.APIClient
	opts   *dtapi.QueryEventsOpts
}

// CreateQuery creates a Builder object to conveniently set
// optional parameters for an event query.
func (api EventAPI) CreateQuery() QueryEventsBuilder {
	builder := queryEventsBuilder{}
	builder.client = api.client
	return &builder
}

// WithStartTimeStamp specifies the end timestamp in milliseconds since Unix epoch
func (builder *queryEventsBuilder) WithStartTimeStamp(timeStamp int64) QueryEventsBuilder {
	builder.opts.From = optional.NewInt64(timeStamp)
	return builder
}

// WithEndTimeStamp specifies the end timestamp in milliseconds since Unix epoch.
// If not specified, the default timeframe is the last 30 days
func (builder *queryEventsBuilder) WithEndTimeStamp(timeStamp int64) QueryEventsBuilder {
	builder.opts.To = optional.NewInt64(timeStamp)
	return builder
}

// WithRelativeTime is the relative timeframe, back from the current time.
func (builder *queryEventsBuilder) WithRelativeTime(relativeTime string) QueryEventsBuilder {
	builder.opts.RelativeTime = optional.NewString(relativeTime)
	return builder
}

// WithEventType filters the event feed based on a specific event type
func (builder *queryEventsBuilder) WithEventType(eventType string) QueryEventsBuilder {
	builder.opts.EventType = optional.NewString(eventType)
	return builder
}

// WithEntityID if specified, it means to only receive events for a given monitored entity, such as a host, process, or service
func (builder *queryEventsBuilder) WithEntityID(entityID string) QueryEventsBuilder {
	builder.opts.EntityId = optional.NewString(entityID)
	return builder
}

// WithCursor If a query returns a cursor string this string can be used to fetch the next 150 events of a query.
// Note that there is no need to specify additional parameters (for instance eventType, from or to) as the cursor string already contains all these parameters
func (builder *queryEventsBuilder) WithCursor(cursor string) QueryEventsBuilder {
	builder.opts.Cursor = optional.NewString(cursor)
	return builder
}

// QueryEvents queries for events based on the parameterization previously defined
// via 'WithStartTimeStamp', 'WithEndTimeStamp, 'WithRelativeTime', 'WithEventType', 'WithEntityID' and 'WithCursor'.
func (builder *queryEventsBuilder) QueryEvents() (dtapi.EventQueryResult, error) {
	result, _, err := builder.client.EventApi.QueryEvents(nil, builder.opts)
	return result, err
}
