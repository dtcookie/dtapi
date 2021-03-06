/*
 * Dynatrace Environment API
 *
 * Documentation of the Dynatrace REST API. Refer to the [help page](https://www.dynatrace.com/support/help/shortlink/section-api) to read about use-cases and examples.
 *
 * API version: 1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package dtapienv


// Describes errors reported during a user session.
type UserSessionErrors struct {
	// The type of event or error. For example, 'Error' or 'Crash' for errors, and 'UserTag' for events.
	Type string `json:"type,omitempty"`
	// The name of the event or error.
	Name string `json:"name,omitempty"`
	// The DNS domain recorded for the user action. Domain is empty if not applicable.
	Domain string `json:"domain,omitempty"`
	// The time when the event or error recorded, in milliseconds since the epoch.
	StartTime int64 `json:"startTime,omitempty"`
	// The name of the application, based on the configured detection rules.
	Application string `json:"application,omitempty"`
	// The internal ID of the application. This information is useful when calling various REST APIs, for example as key for time series queries.
	InternalApplicationId string `json:"internalApplicationId,omitempty"`
}
