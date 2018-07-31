package dynatrace

import (
	"github.com/dtcookie/dtapi/config"
)

// ApplicationDetectionConfigAPI TODO: documentation
type ApplicationDetectionConfigAPI configService

/*
Get Lists all available application detection rules.
@return ApplicationDetectionConfig
*/
func (api ApplicationDetectionConfigAPI) Get() (config.ApplicationDetectionConfig, error) {
	result, _, err := api.client.ApplicationDetectionConfigApi.GetConfiguration(nil)
	if err != nil {
		return result, err
	}
	return result, nil
}

/*
Update Updates an application detection configuration.
All the application detection rules are stored in a single configuration. When you want to edit a rule or create a new one, you have to put **all** existing rules in the request body, otherwise you&#39;ll lose them. Execute the &#x60;GET /applicationDetection&#x60; request to view the list of existing rules. Then modify an existing rule or add a new one, and keep the rest of the rules as they are.   The order of rules is important. Dynatrace evaluates rules from top to bottom. To add a new rule to the configuration, execute the &#x60;GET /applicationDetection&#x60; request, and place the new rule in the required position among the existing rules.    Validate the configuration with the &#x60;POST /validator&#x60; request before submitting it.
 * @param optional nil or *UpdateConfigurationOpts - Optional Parameters:
 * @param "ApplicationDetectionConfig" (optional.Interface of ApplicationDetectionConfig) -
*/
func (api ApplicationDetectionConfigAPI) Update(opts *config.UpdateConfigurationOpts) error {
	_, err := api.client.ApplicationDetectionConfigApi.UpdateConfiguration(nil, opts)
	return err
}

/*
Validate Validates the application detection configuration for the `PUT /applicationDetection` request.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param optional nil or *ValidateConfigurationOpts - Optional Parameters:
 * @param "ApplicationDetectionConfig" (optional.Interface of ApplicationDetectionConfig) -
*/
func (api ApplicationDetectionConfigAPI) Validate(opts *config.ValidateConfigurationOpts) error {
	_, err := api.client.ApplicationDetectionConfigApi.ValidateConfiguration(nil, opts)
	return err
}

// func testme(a int, b int, fn func(int, int) int) int {
// 	return fn(a, b)
// }

// func add(a int, b int) int {
// 	return a + b
// }

// func tested() {
// 	result := testme(1, 2, add)
// 	fmt.Println(result)
// }
