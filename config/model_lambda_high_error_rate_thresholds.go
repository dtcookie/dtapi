/*
 * Dynatrace Configuration API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package config


// Custom thresholds for high error rate detection on Lambda service. If not present (null) then default thresholds will be used.
type LambdaHighErrorRateThresholds struct {
	// Detect high error rate when failed invocations rate is higher than N percent in 3 out of 5 samples.
	FailedInvocationsRate int32 `json:"failedInvocationsRate"`
}
