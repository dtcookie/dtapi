/*
 * Dynatrace Environment API
 *
 * Documentation of the Dynatrace REST API. Refer to the [help page](https://www.dynatrace.com/support/help/shortlink/section-api) to read about use-cases and examples.
 *
 * API version: 1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package environment


// Allows to attach an event to an entity that fits the specified criteria
type TagMatchRule struct {
	// List of ME types the event can be attached to
	MeTypes []string `json:"meTypes,omitempty"`
	// List of Tags an entity has to match
	Tags []TagInfo `json:"tags,omitempty"`
}