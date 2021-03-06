/*
 * Dynatrace Environment API
 *
 * Documentation of the Dynatrace REST API. Refer to the [help page](https://www.dynatrace.com/support/help/shortlink/section-api) to read about use-cases and examples.
 *
 * API version: 1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package dtapienv


type Application struct {
	// Dynatrace entity ID of the required entity.   You can find them in the URL of the corresponding Dynatrace page, for example, `HOST-007`.
	EntityId string `json:"entityId,omitempty"`
	// The name of the Dynatrace entity, displayed in the UI.
	DisplayName string `json:"displayName,omitempty"`
	// Customized name of the entity
	CustomizedName string `json:"customizedName,omitempty"`
	// Discovered name of the entity
	DiscoveredName string `json:"discoveredName,omitempty"`
	// Timestamp in UTC milliseconds when the entity was detected for the first time.
	FirstSeenTimestamp int64 `json:"firstSeenTimestamp,omitempty"`
	// Timestamp in UTC milliseconds when the entity was detected for the last time.
	LastSeenTimestamp int64 `json:"lastSeenTimestamp,omitempty"`
	// The list of entity tags.
	Tags []TagInfo `json:"tags,omitempty"`
	// The list of outgoing calls from the application.
	FromRelationships map[string]interface{} `json:"fromRelationships,omitempty"`
	// The list of incoming calls to the application.
	ToRelationships map[string]interface{} `json:"toRelationships,omitempty"`
	// 
	ApplicationType string `json:"applicationType,omitempty"`
	// 
	RuleAppliedMatchType string `json:"ruleAppliedMatchType,omitempty"`
	// 
	RuleAppliedPattern string `json:"ruleAppliedPattern,omitempty"`
	// 
	ApplicationMatchTarget string `json:"applicationMatchTarget,omitempty"`
}
