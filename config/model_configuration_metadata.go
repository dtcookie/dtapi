/*
 * Dynatrace Configuration API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package config


// Metadata useful for debugging
type ConfigurationMetadata struct {
	// Dynatrace server version.
	ClusterVersion string `json:"clusterVersion,omitempty"`
	// Sorted list containing the version numbers of the configuration.
	ConfigurationVersions []int64 `json:"configurationVersions,omitempty"`
}
