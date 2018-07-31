/*
 * Dynatrace Configuration API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package config


// Fixed thresholds for response time degradation detection. Required if the **detectionMode** is `DETECT_USING_FIXED_THRESHOLDS`. Not applicable otherwise.
type ResponseTimeDegradationThresholdConfig struct {
	// Response time during any 5-minute period to trigger an alert, milliseconds.
	ResponseTimeThresholdMilliseconds int32 `json:"responseTimeThresholdMilliseconds"`
	// Response time of the 10% slowest during any 5-minute period to trigger an alert, milliseconds.
	SlowestResponseTimeThresholdMilliseconds int32 `json:"slowestResponseTimeThresholdMilliseconds"`
	// Mininal service load to detect response time degradation.   Response time degradation of services with smaller load won't trigger alerts.
	LoadThreshold string `json:"loadThreshold"`
	// Sensitivity of the threshold.   With `low` sensitivity high statistical confidence is used, so brief violations (for example, due to a surge in load) won't trigger alerts.   With `high` sensitivity no statistical confidence is used. Each violation triggers alert.
	Sensitivity string `json:"sensitivity"`
}