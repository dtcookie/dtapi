/*
 * Dynatrace Configuration API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package dtapiconf


type ConstraintViolation struct {
	Path string `json:"path,omitempty"`
	Message string `json:"message,omitempty"`
	ParameterLocation string `json:"parameterLocation,omitempty"`
	Location string `json:"location,omitempty"`
}
