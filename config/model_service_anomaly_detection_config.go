/*
 * Dynatrace Configuration API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package config


type ServiceAnomalyDetectionConfig struct {
	Metadata ConfigurationMetadata `json:"metadata,omitempty"`
	ResponseTimeDegradation ResponseTimeDegradationDetectionConfig `json:"responseTimeDegradation"`
	FailureRateIncrease FailureRateIncreaseDetectionConfig `json:"failureRateIncrease"`
}
