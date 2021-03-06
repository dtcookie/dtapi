/*
 * Dynatrace Environment API
 *
 * Documentation of the Dynatrace REST API. Refer to the [help page](https://www.dynatrace.com/support/help/shortlink/section-api) to read about use-cases and examples.
 *
 * API version: 1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package dtapienv


// Allows filtering monitored entities by their type and tags.
type MonitoredEntityFilter struct {
	// The type of the Dynatrace entities, for example hosts or services, you want to pick up by matching.
	Type string `json:"type,omitempty"`
	// The tag you want to use for matching.   You can use custom tags from the UI, AWS tags, Cloud Foundry tags, OpenShift/Kubernetes, and tags based on environment variables.
	Tags []UniversalTag `json:"tags,omitempty"`
}
