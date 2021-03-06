/*
 * Dynatrace Environment API
 *
 * Documentation of the Dynatrace REST API. Refer to the [help page](https://www.dynatrace.com/support/help/shortlink/section-api) to read about use-cases and examples.
 *
 * API version: 1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package dtapienv


// Contains settings for a timeseries query.
type TimeseriesQueryMessage struct {
	// Case-sensitive identifier of the metric, where you want to read data points.   You can find a list of available built-in Dynatrace metrics in the Available timeseries section, and plugin-driven metrics in the Plugin timeseries section.  You can also execute a GET timeseries request, to obtain the list of available metrics.
	TimeseriesId string `json:"timeseriesId,omitempty"`
	// Defines which aggregation type is used for the resulting metric.  If the requested metric doesn't support the specified aggregation, the request will result in an error.
	AggregationType string `json:"aggregationType,omitempty"`
	// Start of timeframe in milliseconds since Unix epoch.
	StartTimestamp int64 `json:"startTimestamp,omitempty"`
	// End of timeframe in milliseconds since Unix epoch. Must be larger(later) than start timestamp. If the given timestamp is larger than the actual time, then the actual time is used instead.
	EndTimestamp int64 `json:"endTimestamp,omitempty"`
	// Used to predict future Values
	Predict bool `json:"predict,omitempty"`
	// Relative timeframe, back from the current time.
	RelativeTime string `json:"relativeTime,omitempty"`
	// Defines the type of result that the call should return. Valid result modes are: series: returns all the data points of the metric in the specified timeframe. total: returns one scalar value for the specified timeframe.   By default, the series mode is used.
	QueryMode string `json:"queryMode,omitempty"`
	// Filters requested data points by entities which should deliver them. You can specify several entities at once.   Allowed values are Dynatrace entity IDs. You can find them in the URL of the corresponding Dynatrace entity page, for example, HOST-007.   If the selected entity doesn't support the requested metric, the request will result in an error.
	Entities []string `json:"entities,omitempty"`
	// List of labels of entities that you want to fetch data for.
	Tags []string `json:"tags,omitempty"`
	// A filter is an object, containing map of filter keys and its values. Valid filter keys are: processType: Filters by process type. See Process types for allowed values. osType: Filters by operating system. See OS types for allowed values. serviceType: Filters by service type. See Service types for allowed values. technology: Filters by technology type. See Technology types for allowed values. webServiceName: Filters by web service name. webServiceNamespace: Filters by the web service namespace. host: Filters by entity ID of the host, for example HOST-007.
	Filters map[string]string `json:"filters,omitempty"`
	// In case of the percentile aggregation type, this parameter specifies which percentile of the selected response time metric should be delivered.   Valid values for percentile are between 1 and 99.   Please keep in mind that percentile export is only possible for response-time based metrics such as application and service response times.
	Percentile int32 `json:"percentile,omitempty"`
}
