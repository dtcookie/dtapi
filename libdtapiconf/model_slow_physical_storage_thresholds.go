/*
 * Dynatrace Configuration API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package dtapiconf


// Custom thresholds for slow running physical storage device detection. If not present (null) then default thresholds will be used.
type SlowPhysicalStorageThresholds struct {
	// Detect slow running physical storage device always when read/write latency is higher than N milliseconds in 4 out of 5 samples.
	AvgReadWriteLatency int32 `json:"avgReadWriteLatency"`
	// Detect slow running physical storage device always when peak value for read/write latency is higher than N milliseconds in 4 out of 5 samples.
	PeakReadWriteLatency int32 `json:"peakReadWriteLatency"`
}