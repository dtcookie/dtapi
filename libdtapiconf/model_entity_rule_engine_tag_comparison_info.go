/*
 * Dynatrace Configuration API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package dtapiconf


type EntityRuleEngineTagComparisonInfo struct {
	// Operator comparing the extracted value to the comparison value.
	Operator string `json:"operator"`
	Value Tag `json:"value,omitempty"`
	// Negate the comparison.
	Negate bool `json:"negate"`
}
