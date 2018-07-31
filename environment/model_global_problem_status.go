/*
 * Dynatrace Environment API
 *
 * Documentation of the Dynatrace REST API. Refer to the [help page](https://www.dynatrace.com/support/help/shortlink/section-api) to read about use-cases and examples.
 *
 * API version: 1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package environment


// Count of open problems in your environment.
type GlobalProblemStatus struct {
	// Total number of open problems in your environment.
	TotalOpenProblemsCount int32 `json:"totalOpenProblemsCount,omitempty"`
	// Number of open problems per impact level.
	OpenProblemCounts map[string]interface{} `json:"openProblemCounts,omitempty"`
}
