/*
 * Dynatrace Configuration API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package dtapiconf


type TagFilterDto struct {
	// The tag context.
	Context string `json:"context,omitempty"`
	// The tag key.
	Key string `json:"key"`
	// The tag value.
	Value string `json:"value,omitempty"`
}
