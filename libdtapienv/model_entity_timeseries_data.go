/*
 * Dynatrace Environment API
 *
 * Documentation of the Dynatrace REST API. Refer to the [help page](https://www.dynatrace.com/support/help/shortlink/section-api) to read about use-cases and examples.
 *
 * API version: 1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package dtapienv


// Information about a metric and its data points.
type EntityTimeseriesData struct {
	// Identifier of the metric, where you want to post data points.
	TimeseriesId string `json:"timeseriesId"`
	// Dimensions of the data points you're posting.   The key of the metric dimension must be defined earlier in the metric definition.
	Dimensions map[string]string `json:"dimensions,omitempty"`
	// List of data points.   Each data point is an array, containing the timestamp and the value.   Timestamp is UTC milliseconds reported as a number, for example: `1520523365557`.   Values can't be reported further than 2 hours into the past!   A custom metric must be registered **before** you can report a metric value. Therefore, the timestamp for reporting a value must be after the registration time of the metric.
	DataPoints [][]float32 `json:"dataPoints,omitempty"`
}
