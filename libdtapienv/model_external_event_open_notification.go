/*
 * Dynatrace Environment API
 *
 * Documentation of the Dynatrace REST API. Refer to the [help page](https://www.dynatrace.com/support/help/shortlink/section-api) to read about use-cases and examples.
 *
 * API version: 1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package dtapienv


// Open Synthetic event.
type ExternalEventOpenNotification struct {
	// Tne ID of the Synthetic test.
	TestId string `json:"testId,omitempty"`
	// The unique ID of the event.
	EventId string `json:"eventId,omitempty"`
	// The name of the event, displayed in the UI.
	Name string `json:"name,omitempty"`
	// The type of the event.
	EventType string `json:"eventType,omitempty"`
	// IDs of locations where the event happens.
	LocationIds []string `json:"locationIds,omitempty"`
	// Start timestamp of the event.
	StartTimestamp int64 `json:"startTimestamp,omitempty"`
}