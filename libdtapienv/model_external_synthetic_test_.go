/*
 * Dynatrace Environment API
 *
 * Documentation of the Dynatrace REST API. Refer to the [help page](https://www.dynatrace.com/support/help/shortlink/section-api) to read about use-cases and examples.
 *
 * API version: 1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package dtapienv


// Synthetic test data.
type ExternalSyntheticTest struct {
	// Unique ID of the Synthetic test.
	Id string `json:"id"`
	// Test title.
	Title string `json:"title,omitempty"`
	// Test description.
	Description string `json:"description,omitempty"`
	// Test setup information, for example `browser`.
	TestSetup string `json:"testSetup,omitempty"`
	// Test expiration timestamp, UTC milliseconds in Dynatrace Cluster Time.
	ExpirationTimestamp int64 `json:"expirationTimestamp,omitempty"`
	// A URL to the test results.
	DrilldownLink string `json:"drilldownLink,omitempty"`
	// A URL to edit the test in the original UI.
	EditLink string `json:"editLink,omitempty"`
	// Test enabled/disabled. Default is `true`.   If `true`, set the **deleted** parameter to `false`.
	Enabled bool `json:"enabled,omitempty"`
	// Deleted test flag. Default is `false`.    If `true`, set the **enabled** parameter to `false`.
	Deleted bool `json:"deleted,omitempty"`
	// Locations where the Synthetic test runs.
	Locations []SyntheticTestLocation `json:"locations,omitempty"`
	// Steps of the Synthetic test.
	Steps []SyntheticTestStep `json:"steps,omitempty"`
	// Test timeout in milliseconds. If no test result is reported within this time, the availability state switches to unmonitored. Default is doubled test schedule interval.
	NoDataTimeout int32 `json:"noDataTimeout,omitempty"`
	// Test schedule in seconds. The is repeated with the specified interval at the external source.   Dynatrace expects test results with the specified interval. If you report test results to Dynatrace less often, adjust the **noDataTimeout** value accordingly.
	ScheduleIntervalInSeconds int32 `json:"scheduleIntervalInSeconds"`
}
