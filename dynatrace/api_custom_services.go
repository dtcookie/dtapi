package dynatrace

import (
	"github.com/dtcookie/dtapi/config"
)

// CustomServicesAPI TODO: documentation
type CustomServicesAPI configService

/*
Delete Deletes the specified custom service
 * @param technology Technology of the custom service to delete.
 * @param id The ID of the custom service to delete.
*/
func (api CustomServicesAPI) Delete(technology string, id string) error {
	_, err := api.client.CustomServicesApi.Delete(nil, technology, id)
	return err
}

/*
GetItem Gets the definition of the specified custom service
 * @param technology Technology of the custom service you're inquiring.
 * @param id The ID of the custom service you're inquiring.
 * @param optional nil or *GetItemOpts - Optional Parameters:
 * @param "IncludeProcessGroupReferences" (optional.Bool) -  Flag to include process group references to the response.    Process group references aren't compatible across environments.   `false` is used if the value is not set.
@return CustomService
*/
func (api CustomServicesAPI) GetItem(technology string, id string, opts *config.GetItemOpts) (config.CustomService, error) {
	result, _, err := api.client.CustomServicesApi.GetItem(nil, technology, id, opts)
	if err != nil {
		return result, err
	}
	return result, nil
}

/*
GetList Lists all custom services of the specified technology along with their definitions
 * @param technology Technology of the required custom services.
 * @param optional nil or *GetListOpts - Optional Parameters:
 * @param "IncludeProcessGroupReferences" (optional.Bool) -  Flag to include process group references to the response.    Process group references aren't compatible across environments.   `false` is used if the value is not set.
@return CustomServiceList
*/
func (api CustomServicesAPI) GetList(technology string, opts *config.GetListOpts) (config.CustomServiceList, error) {
	result, _, err := api.client.CustomServicesApi.GetList(nil, technology, opts)
	if err != nil {
		return result, err
	}
	return result, nil
}

/*
Post Creates a custom service
In the body of the request, neither custom service nor its rules can have the ID. All IDs will be generated automatically by Dynatrace.
 * @param technology Technology of the new custom service.
 * @param optional nil or *PostOpts - Optional Parameters:
 * @param "CustomService" (optional.Interface of CustomService) -  JSON body of the request containing definition of the new custom service.

You must not specify the IDs for the custom service or any of its rules.
*/
func (api CustomServicesAPI) Post(technology string, opts *config.PostOpts) error {
	_, err := api.client.CustomServicesApi.Post(nil, technology, opts)
	return err
}

/*
PutItem Updates the specified custom service.
 * @param technology Technology of the custom service to update.
 * @param id The ID of the custom service to update.   The ID of the custom service in the body of the request must match this ID.
 * @param optional nil or *PutItemOpts - Optional Parameters:
 * @param "CustomService" (optional.Interface of CustomService) -  JSON body of the request containing updated definition of the custom service.
*/
func (api CustomServicesAPI) PutItem(technology string, id string, opts *config.PutItemOpts) error {
	_, err := api.client.CustomServicesApi.PutItem(nil, technology, id, opts)
	return err
}

/*
PutList Updates all custom services of the specified technology
This request affects **all** the custom services of the specified technology. If a custom service is not presented in the body of the request, it will be deleted.
 * @param technology Technology of custom services to update.
 * @param optional nil or *PutListOpts - Optional Parameters:
 * @param "CustomServiceList" (optional.Interface of CustomServiceList) -  JSON body of the request containing updated definitions of custom services.
*/
func (api CustomServicesAPI) PutList(technology string, opts *config.PutListOpts) error {
	_, err := api.client.CustomServicesApi.PutList(nil, technology, opts)
	return err
}

/*
ValidateItem Validate the new custom service for the `POST /customServices/{technology}` request
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param technology Technology of the custom service to validate.
 * @param optional nil or *ValidateItemOpts - Optional Parameters:
 * @param "CustomService" (optional.Interface of CustomService) -  JSON body of the request containing definition of the custom service to validate.
*/
func (api CustomServicesAPI) ValidateItem(technology string, opts *config.ValidateItemOpts) error {
	_, err := api.client.CustomServicesApi.ValidateItem(nil, technology, opts)
	return err
}

/*
ValidateList Validates custom services for the `PUT /customServices/` request.
 * @param technology Technology of the custom services to validate.
 * @param optional nil or *ValidateListOpts - Optional Parameters:
 * @param "CustomServiceList" (optional.Interface of CustomServiceList) -  JSON body of the request containing definitions of the custom services to validate.

It must contain definitions of **all** custom services of the required technology.
*/
func (api CustomServicesAPI) ValidateList(technology string, opts *config.ValidateListOpts) error {
	_, err := api.client.CustomServicesApi.ValidateList(nil, technology, opts)
	return err
}
