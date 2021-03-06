/*
 * Dynatrace Configuration API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package dtapiconf


// Fixed thresholds for failure rate increase detection. Required if the **detectionMode** is `DETECT_USING_FIXED_THRESHOLDS`. Not applicable otherwise.
type FailureRateIncreaseThresholdConfig struct {
	// Failure rate during any 5-minute period to trigger an alert, %.
	Threshold int32 `json:"threshold"`
	// Sensitivity of the threshold.   With `low` sensitivity high statistical confidence is used, so brief violations (for example, due to a surge in load) won't trigger alerts.   With `high` sensitivity no statistical confidence is used. Each violation triggers alert.
	Sensitivity string `json:"sensitivity"`
}
