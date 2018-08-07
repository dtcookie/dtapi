/*
 * Dynatrace Environment API
 *
 * Documentation of the Dynatrace REST API. Refer to the [help page](https://www.dynatrace.com/support/help/shortlink/section-api) to read about use-cases and examples.
 *
 * API version: 1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package dtapienv


// Synthetic location.
type SyntheticLocation struct {
	// Unique ID of the location.
	Id string `json:"id"`
	// Location name displayed in the UI.
	Name string `json:"name,omitempty"`
	// IP address of the location.
	Ip string `json:"ip"`
}