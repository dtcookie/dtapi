/*
 * Dynatrace Configuration API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package dtapiconf


type Error struct {
	Code int32 `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
	ConstraintViolations []ConstraintViolation `json:"constraintViolations,omitempty"`
}