/*
 * Dynatrace Configuration API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package dtapiconf


// Configuration of anomaly detection for AWS
type AwsAnomalyDetectionConfig struct {
	Metadata ConfigurationMetadata `json:"metadata,omitempty"`
	RdsHighCpuDetection RdsHighCpuDetectionConfig `json:"rdsHighCpuDetection"`
	RdsHighWriteReadLatencyDetection RdsHighWriteReadLatencyDetectionConfig `json:"rdsHighWriteReadLatencyDetection"`
	RdsLowStorageDetection RdsLowStorageDetectionConfig `json:"rdsLowStorageDetection"`
	RdsHighMemoryDetection RdsHighMemoryDetectionConfig `json:"rdsHighMemoryDetection"`
	ElbHighConnectionErrorsDetection ElbHighConnectionErrorsDetectionConfig `json:"elbHighConnectionErrorsDetection"`
	RdsRestartsSequenceDetection RdsRestartsSequenceDetectionConfig `json:"rdsRestartsSequenceDetection"`
	LambdaHighErrorRateDetection LambdaHighErrorRateDetectionConfig `json:"lambdaHighErrorRateDetection"`
}
