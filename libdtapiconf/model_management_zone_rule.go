/*
 * Dynatrace Configuration API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package dtapiconf


// A rule defines when to apply a management zone. Each rule is evaluated independently of all other rules.
type ManagementZoneRule struct {
	// Type of entities for which the management zone rule should be evaluated and associated to if it matches.
	Type string `json:"type"`
	// Management zone evaluation rule enabled/disabled
	Enabled bool `json:"enabled"`
	// Types of propagation to be done on a matching management zone. Only some may be valid for the specified 'type'.
	PropagationTypes []string `json:"propagationTypes,omitempty"`
	// Actual conditions of the rule when the rule should match for an entity of the specified 'type'.Only if all conditions match for an entity, the evaluation is positive and the management zone is associated with the entity.
	Conditions []EntityRuleEngineCondition `json:"conditions"`
}
