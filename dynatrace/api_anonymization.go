package dynatrace

import (
	"github.com/antihax/optional"
	"github.com/dtcookie/dtapi/environment"
)

// AnonymizationAPI TODO: documentation
type AnonymizationAPI service

// CreateAnonymizationJobOP TODO: documentation
type CreateAnonymizationJobOP struct {
	api       AnonymizationAPI
	optionals environment.AnonymizeOpts
}

// StartTimeStamp TODO: documentation
func (op CreateAnonymizationJobOP) StartTimeStamp(timeStamp int64) {
	op.optionals.StartTimestamp = optional.NewInt64(timeStamp)
}

// Execute TODO: documentation
func (op CreateAnonymizationJobOP) Execute() (string, error) {
	return op.api.Create(&op.optionals)
}

/*
Create Creates user session anonymization job.
The job anonymizes all user sessions in the specified timeframe by masking the userID and IP address.   To identify user sessions to be anonymized you can specify either userID, or IP address, or both. If you specify both the **OR** logic applies.   Regardless how you&#39;re identifying the user session, both userID and IP are masked. You can&#39;t undo the anonymization.
 * @param optional nil or *AnonymizeOpts - Optional Parameters:
 * @param "StartTimestamp" (optional.Int64) -  The start timestamp of the user session to anonymize, in UTC milliseconds. If 0, the earliest available time is used.
 * @param "EndTimestamp" (optional.Int64) -  The end timestamp of the user session to anonymize, in UTC milliseconds. If 0, the current time is used.
 * @param "UserIds" (optional.Interface of []string) -  UserID of the user to anonymize.    You can specify several IDs, in the following format: `userIds=user1&userIds=user2`.
 * @param "Ips" (optional.Interface of []string) -  IP address of the user to anonymize. All user sessions from this IP will be anonymized.   You can specify several IPs, in a same way as user IDs.
@return AnonymizationIdResult
*/
func (api AnonymizationAPI) Create(options *environment.AnonymizeOpts) (string, error) {
	result, _, err := api.client.AnonymizationApi.Anonymize(nil, options)
	if err != nil {
		return result.RequestId, err
	}
	return result.RequestId, nil
}

/*
Status Shows the progress of the specified anonymization job.
 * @param requestID
@return AnonymizationProgressResult
*/
func (api AnonymizationAPI) Status(requestID string) (int32, error) {
	result, _, err := api.client.AnonymizationApi.GetStatus(nil, requestID)
	if err != nil {
		return result.Progress, err
	}
	return result.Progress, nil
}
