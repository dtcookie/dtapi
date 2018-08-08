package dtapi

import (
	"github.com/antihax/optional"
	dtapi "github.com/dtcookie/dtapi/libdtapienv"
)

// AnonymizationAPI is a cleaned up wrapper around the
// Anonymization Services offered via
// "github.com/dtcookie/dtapi/libdtapienv"
type AnonymizationAPI struct {
	client    *dtapi.APIClient
	optionals dtapi.AnonymizeOpts
}

// Create Creates a User Session Anonymization Job
//
// Note:
//	You can parameterize that Anonymization Job before
//	calling this function using the build functions
//	'WithStartTimeStamp', 'WithEndTimeStamp', 'WithIPs' and 'WithUserIDs'
//
// Returns:
//	- string	the 'requestID' for the created Anonymization Job
//				to be used to check its status later on
func (api *AnonymizationAPI) Create() (string, error) {
	result, _, err := api.client.AnonymizationApi.Anonymize(nil, &api.optionals)
	api.optionals = dtapi.AnonymizeOpts{}
	return result.RequestId, err
}

// Status Shows the progress of the Anonymization Job identified with
// the specified 'requestID'.
//
// Returns:
//	- int32	The progress of the Anonymization Job in percents.
//			-1 if the job is waiting for execution
func (api *AnonymizationAPI) Status(requestID string) (int32, error) {
	result, _, err := api.client.AnonymizationApi.GetStatus(nil, requestID)
	return result.Progress, err
}

// WithIPs allows for parameterizing the creation of an
// Anonymization Job with a list of IP Addresses to
// anonymize before invoking 'Create'.
func (api *AnonymizationAPI) WithIPs(IPs []string) {
	api.optionals.Ips = optional.NewInterface(IPs)
}

// WithUserIDs allows for parameterizing the creation of an
// Anonymization Job with a list of User IDs to
// anonymize before invoking 'Create'.
func (api *AnonymizationAPI) WithUserIDs(userIDs []string) {
	api.optionals.UserIds = optional.NewInterface(userIDs)
}

// WithEndTimeStamp allows for parameterizing the creation of an
// Anonymization Job with an end timestamp before invoking 'Create'.
// Parameters:
//	- timeStamp	is the end timestamp of the user session to anonymize,
//				in UTC milliseconds.
//				If 0, the current time is used.
func (api *AnonymizationAPI) WithEndTimeStamp(timeStamp int64) {
	api.optionals.EndTimestamp = optional.NewInt64(timeStamp)
}

// WithStartTimeStamp allows for parameterizing the creation of an
// Anonymization Job with a start timestamp before invoking 'Create'.
// Parameters:
//	- timeStamp	is the start timestamp of the user session to anonymize,
//				in UTC milliseconds.
//				If 0, the earliest available time is used.
func (api *AnonymizationAPI) WithStartTimeStamp(timeStamp int64) {
	api.optionals.StartTimestamp = optional.NewInt64(timeStamp)
}
