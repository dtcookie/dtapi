/*
 * Dynatrace Configuration API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package config


// Configuration of low disk inodes number detection
type DiskLowInodesDetectionConfig struct {
	// Is detection enabled.
	Enabled bool `json:"enabled"`
	CustomThresholds DiskLowInodesThresholds `json:"customThresholds,omitempty"`
}
