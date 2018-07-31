/*
 * Dynatrace Configuration API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package config


type ValueCondition struct {
	// Operator comparing the extracted value to the comparison value.
	Operator string `json:"operator"`
	// Negate the comparison.
	Negate bool `json:"negate"`
	// The value to compare to.
	Value string `json:"value"`
}