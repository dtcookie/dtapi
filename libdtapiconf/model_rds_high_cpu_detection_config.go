/*
 * Dynatrace Configuration API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package dtapiconf


// Configuration of high CPU utilization detection on RDS
type RdsHighCpuDetectionConfig struct {
	// Is detection enabled.
	Enabled bool `json:"enabled"`
	CustomThresholds RdsHighCpuThresholds `json:"customThresholds,omitempty"`
}
