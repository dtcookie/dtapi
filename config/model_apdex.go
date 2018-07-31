/*
 * Dynatrace Configuration API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package config


// Defines the apdex settings of an application.
type Apdex struct {
	// Apdex user satisfaction threshold (decimal number with one fractional digit).
	Threshold float32 `json:"threshold"`
	// Consider JavaScript errors in Apdex calculations
	ConsiderJavaScriptErrors bool `json:"considerJavaScriptErrors"`
}
