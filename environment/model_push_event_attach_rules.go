/*
 * Dynatrace Environment API
 *
 * Documentation of the Dynatrace REST API. Refer to the [help page](https://www.dynatrace.com/support/help/shortlink/section-api) to read about use-cases and examples.
 *
 * API version: 1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package environment


// Defines which monitored entities a event is to be attached to
type PushEventAttachRules struct {
	// Array of entity identifiers, the event should be attached to. This can also just be a single value.
	EntityIds []string `json:"entityIds,omitempty"`
	// Array of tagRules which attach the event on all components that match the specific rules. This can also just be a single value.
	TagRule []TagMatchRule `json:"tagRule,omitempty"`
}
