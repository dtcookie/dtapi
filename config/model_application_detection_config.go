/*
 * Dynatrace Configuration API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package config


type ApplicationDetectionConfig struct {
	Metadata ConfigurationMetadata `json:"metadata,omitempty"`
	// Ordered list of application detection rules.
	DetectionRules []ApplicationDetectionRule `json:"detectionRules"`
}
