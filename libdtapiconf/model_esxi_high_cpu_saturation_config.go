/*
 * Dynatrace Configuration API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package dtapiconf


// Configuration of high CPU saturation detection on ESXi host
type EsxiHighCpuSaturationConfig struct {
	// Is detection enabled.
	Enabled bool `json:"enabled"`
	CustomThresholds EsxiHighCpuThresholds `json:"customThresholds,omitempty"`
}
