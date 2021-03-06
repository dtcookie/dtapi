/*
 * Dynatrace Environment API
 *
 * Documentation of the Dynatrace REST API. Refer to the [help page](https://www.dynatrace.com/support/help/shortlink/section-api) to read about use-cases and examples.
 *
 * API version: 1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package dtapienv


// An object defining date, time, and recurrence of the maintenance window.
type MaintenanceWindowSchedule struct {
	// Recurrence of the schedule.
	Type string `json:"type"`
	// Time zone of the start and end time. Default time zone is UTC.    You can user either UTC offset `UTC+01:00` format or the IANA Time Zone Database format.
	TimezoneId string `json:"timezoneId,omitempty"`
	// Start date and time of the maintenance window in the `yyyy-MM-dd HH:mm` format.
	MaintenanceStart string `json:"maintenanceStart"`
	// Start date and time of the maintenance window in the `yyyy-MM-dd HH:mm` format.
	MaintenanceEnd string `json:"maintenanceEnd"`
	Recurrence MaintenanceWindowRecurrence `json:"recurrence,omitempty"`
}
