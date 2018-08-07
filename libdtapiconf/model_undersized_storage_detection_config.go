/*
 * Dynatrace Configuration API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package dtapiconf


// Configuration of undersized storage device detection
type UndersizedStorageDetectionConfig struct {
	// Is detection enabled.
	Enabled bool `json:"enabled"`
	CustomThresholds UndersizedStorageThresholds `json:"customThresholds,omitempty"`
}