/*
 * Dynatrace Configuration API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package config


type CapturedMethod struct {
	Method MethodReference `json:"method"`
	// What to capture from the method.
	Capture string `json:"capture"`
	// The index of the argument to capture. Set `0` to capture the return value, `1` or higher to capture a mehtod argument.    Required if the **capture** is set to `ARGUMENT`.   Not applicable in other cases.
	ArgumentIndex int32 `json:"argumentIndex,omitempty"`
	// The getter chain to apply to the captured object. It is required in one of the following cases:   The **capture** is set to `THIS`.    The **capture** is set to `ARGUMENT`, and the argument is not a primitive, a primitive wrapper class, a string, or an array.    Not applicable in other cases.
	DeepObjectAccess string `json:"deepObjectAccess,omitempty"`
}
