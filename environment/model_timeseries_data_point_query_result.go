/*
 * Dynatrace Environment API
 *
 * Documentation of the Dynatrace REST API. Refer to the [help page](https://www.dynatrace.com/support/help/shortlink/section-api) to read about use-cases and examples.
 *
 * API version: 1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package environment


// List of datapoints, as well as attributes describing the unit of the returned data points, the result resolution of the metric and the selected type of aggregation.
type TimeseriesDataPointQueryResult struct {
	// Metric data points   A JSON object that maps the ID of the entity that delivered the data points and an array, which consists of arrays of the data point float values.  May contain more that one entity ID per record (for example, a host and its network interface). In such cases, entity IDs are separated by commas. A datapoint contains a value and a timestamp, at which the value was recorded. There are three versions of data points: Numeric datapoint: contains a numeric value Enum datapoint: contains an enum value (currently only for availability timeseries) Prediction datapoint: Similar to the numeric datapoint, but it contains a confidence interval, within which the future values are expected to be
	DataPoints map[string][][]float32 `json:"dataPoints,omitempty"`
	// Unit of data points.
	Unit string `json:"unit,omitempty"`
	// The type of data points aggregation.
	AggregationType string `json:"aggregationType,omitempty"`
	// Resolution of data points.
	ResolutionInMillisUTC int64 `json:"resolutionInMillisUTC,omitempty"`
	// Entities where the data points originate.  A JSON object that maps the entity ID in Dynatrace and the actual name of the entity.
	Entities map[string]string `json:"entities,omitempty"`
	// The ID of the metric.
	TimeseriesId string `json:"timeseriesId,omitempty"`
}
