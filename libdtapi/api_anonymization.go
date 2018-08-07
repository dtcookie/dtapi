package dtapi

import (
	"github.com/antihax/optional"
	dtapi "github.com/dtcookie/dtapi/libdtapienv"
)

// AnonymizationAPI TODO: documentation
type AnonymizationAPI struct {
	client    *dtapi.APIClient
	optionals dtapi.AnonymizeOpts
}

// WithIPs TODO: documentation
func (api *AnonymizationAPI) WithIPs(IPs []string) {
	api.optionals.Ips = optional.NewInterface(IPs)
}

// WithUserIDs TODO: documentation
func (api *AnonymizationAPI) WithUserIDs(userIDs []string) {
	api.optionals.UserIds = optional.NewInterface(userIDs)
}

// WithEndTimeStamp TODO: documentation
func (api *AnonymizationAPI) WithEndTimeStamp(timeStamp int64) {
	api.optionals.EndTimestamp = optional.NewInt64(timeStamp)
}

// WithStartTimeStamp TODO: documentation
func (api *AnonymizationAPI) WithStartTimeStamp(timeStamp int64) {
	api.optionals.StartTimestamp = optional.NewInt64(timeStamp)
}

// CREATE TODO: documentation
func (api *AnonymizationAPI) CREATE() (string, error) {
	result, _, err := api.client.AnonymizationApi.Anonymize(nil, &api.optionals)
	return result.RequestId, err
}

/*
STATUS Shows the progress of the specified anonymization job.
 * @param requestID the unique identifier of the anonymization job
@return int32 The progress of the anonymization job, percents. -1 if the job is waiting for execution
*/
func (api *AnonymizationAPI) STATUS(requestID string) (int32, error) {
	result, _, err := api.client.AnonymizationApi.GetStatus(nil, requestID)
	return result.Progress, err
}
