/*
 * Dynatrace Environment API
 *
 * Documentation of the Dynatrace REST API. Refer to the [help page](https://www.dynatrace.com/support/help/shortlink/section-api) to read about use-cases and examples.
 *
 * API version: 1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package environment


// Representation of a metric configuration with all its parameters and optionally data points.
type TimeseriesQueryResult struct {
	// The name of the metric in the user interface.
	DisplayName string `json:"displayName,omitempty"`
	// Fine metric division, for example process group and process ID for some process-related metric.
	Dimensions []string `json:"dimensions,omitempty"`
	// Units of the metric.
	Unit string `json:"unit,omitempty"`
	// The feature, where the metric originates.
	DetailedSource string `json:"detailedSource,omitempty"`
	// The ID of the plugin, where the metric originates.
	PluginId string `json:"pluginId,omitempty"`
	// Technology type definition. Used to group metrics under a logical technology name.
	Types []string `json:"types,omitempty"`
	DataResult TimeseriesDataPointQueryResult `json:"dataResult,omitempty"`
	// The list of allowed aggregations for this metric.
	AggregationTypes []string `json:"aggregationTypes,omitempty"`
	// The feature, where the metric originates.
	Filter string `json:"filter,omitempty"`
	// The ID of the metric.
	TimeseriesId string `json:"timeseriesId,omitempty"`
}
