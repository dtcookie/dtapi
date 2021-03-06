/*
 * Dynatrace Configuration API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package dtapiconf


type EntityRuleEngineDatabaseTopologyComparisonInfo struct {
	// Operator comparing the extracted value to the comparison value.
	Operator string `json:"operator"`
	// Holds the actual value which should be used for comparison. Other parameters of this condition define how this value should be compared to actually stored values.
	Value string `json:"value,omitempty"`
	// Negate the comparison.
	Negate bool `json:"negate"`
}
