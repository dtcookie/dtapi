/*
 * Dynatrace Configuration API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package config


type DiskNameFilterDto struct {
	// Comparison operator.
	Operator string `json:"operator"`
	// Value to match.
	Value string `json:"value"`
}
