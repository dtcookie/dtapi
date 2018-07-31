/*
 * Dynatrace Environment API
 *
 * Documentation of the Dynatrace REST API. Refer to the [help page](https://www.dynatrace.com/support/help/shortlink/section-api) to read about use-cases and examples.
 *
 * API version: 1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package environment


// Additional data on the event severity.
type EventSeverity struct {
	// The metric used in the event severity calculation.
	Context string `json:"context,omitempty"`
	// The value of the severity.
	Value float64 `json:"value,omitempty"`
	// The unit of the severity value.
	Unit string `json:"unit,omitempty"`
}
