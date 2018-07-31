/*
 * Dynatrace Configuration API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package config


// Custom thresholds for high number of backend connection errors detection on ELB. If not present (null) then default thresholds will be used.
type ElbHighConnectionErrorsThresholds struct {
	// Detect high number of connection errors when number of backend connection errors is higher than N per minute in 3 out of 5 samples.
	ConnectionErrorsPerMinute int32 `json:"connectionErrorsPerMinute"`
}
