package dtapi

import (
	"github.com/antihax/optional"
	dtapi "github.com/dtcookie/dtapi/libdtapiconf"
)

// TechCustomServicesAPI TODO: documentation
type TechCustomServicesAPI struct {
	confService
	technology string
	All        TechAllCustomServicesAPI
}

// CustomServicesAPI TODO: documentation
type CustomServicesAPI struct {
	Java   TechCustomServicesAPI
	DotNet TechCustomServicesAPI
	PHP    TechCustomServicesAPI
}

// CustomServiceAPI TODO: documentation
type CustomServiceAPI struct {
	confService
	technology string
	ID         string
}

func newCustomServiceAPI(technology string, ID string, client *dtapi.APIClient) CustomServiceAPI {
	api := CustomServiceAPI{technology: technology, ID: ID}
	api.client = client
	return api
}

/*
ForID Provides services for addressing a specific custom service configuration
*/
func (api TechCustomServicesAPI) ForID(ID string) CustomServiceAPI {
	return newCustomServiceAPI(api.technology, ID, api.client)
}

/*
DELETE Deletes the this custom service
*/
func (api CustomServiceAPI) DELETE() error {
	_, err := api.client.CustomServicesApi.Delete(nil, api.technology, api.ID)
	return err
}

/*
GET Gets the definition of the custom service
@return CustomService
*/
func (api CustomServiceAPI) GET(withProcessGroupReferences bool) (dtapi.CustomService, error) {
	result, _, err := api.client.CustomServicesApi.GetItem(nil, api.technology, api.ID, &dtapi.GetItemOpts{
		IncludeProcessGroupReferences: optional.NewBool(withProcessGroupReferences),
	})
	return result, err
}

/*
UPDATE Updates the definition of the custom service
*/
func (api CustomServiceAPI) UPDATE(c dtapi.CustomService) error {
	_, err := api.client.CustomServicesApi.PutItem(nil, api.technology, api.ID, &dtapi.PutItemOpts{
		CustomService: optional.NewInterface(c),
	})
	return err
}

func newTechCustomServicesAPI(technology string, client *dtapi.APIClient) TechCustomServicesAPI {
	api := TechCustomServicesAPI{technology: technology}
	api.client = client
	api.All = newTechAllCustomServicesAPI(technology, client)
	return api
}

// NewCustomServicesAPI TODO: documentation
func NewCustomServicesAPI(client *dtapi.APIClient) CustomServicesAPI {
	return CustomServicesAPI{
		Java:   newTechCustomServicesAPI("java", client),
		DotNet: newTechCustomServicesAPI("dotNet", client),
		PHP:    newTechCustomServicesAPI("php", client),
	}
}

// TechAllCustomServicesAPI TODO: documentation
type TechAllCustomServicesAPI struct {
	confService
	technology string
}

func newTechAllCustomServicesAPI(technology string, client *dtapi.APIClient) TechAllCustomServicesAPI {
	api := TechAllCustomServicesAPI{technology: technology}
	api.client = client
	return api
}

/*
GET Lists all custom services of the specified technology along with their definitions
@return []dtapi.CustomService
*/
func (api TechAllCustomServicesAPI) GET(withProcessGroupReferences bool) ([]dtapi.CustomService, error) {
	result, _, err := api.client.CustomServicesApi.GetList(nil, api.technology, &dtapi.GetListOpts{
		IncludeProcessGroupReferences: optional.NewBool(withProcessGroupReferences),
	})
	return result.CustomServices, err
}

/*
UPDATE Updates all custom services for this technology
This request affects **all** the custom services of this technology. If a custom service is not presented in the body of the request, it will be deleted.
 * @param []dtapi.CustomService the definitions of the custom services to update.
*/
func (api TechAllCustomServicesAPI) UPDATE(services []dtapi.CustomService) error {
	_, err := api.client.CustomServicesApi.PutList(nil, api.technology, &dtapi.PutListOpts{
		CustomServiceList: optional.NewInterface(dtapi.CustomServiceList{
			CustomServices: services,
		}),
	})
	return err
}

/*
VALIDATE Validates custom services for the `PUT /customServices/` request.
 * @param []dtapi.CustomService -  the definitions of the custom services for this technology.

It must contain definitions of **all** custom services of the required technology.
*/
func (api TechAllCustomServicesAPI) VALIDATE(services []dtapi.CustomService) error {
	_, err := api.client.CustomServicesApi.ValidateList(nil, api.technology, &dtapi.ValidateListOpts{
		CustomServiceList: optional.NewInterface(dtapi.CustomServiceList{
			CustomServices: services,
		}),
	})
	return err
}

/*
CREATE Creates (HTTP POST) a custom service
In the body of the request, neither custom service nor its rules can have the ID. All IDs will be generated automatically by Dynatrace.
 * @param technology Technology of the new custom service.
 * @param optional nil or *PostOpts - Optional Parameters:
 * @param "CustomService" (optional.Interface of CustomService) -  JSON body of the request containing definition of the new custom service.

You must not specify the IDs for the custom service or any of its rules.
*/
func (api TechCustomServicesAPI) CREATE(c dtapi.CustomService) error {
	_, err := api.client.CustomServicesApi.Post(nil, api.technology, &dtapi.PostOpts{
		CustomService: optional.NewInterface(c),
	})
	return err
}

/*
VALIDATE Validate the new custom service for the `POST /customServices/{technology}` request
 * @param dtapi.CustomService the definition of the custom service to create
*/
func (api TechCustomServicesAPI) VALIDATE(service dtapi.CustomService) error {
	_, err := api.client.CustomServicesApi.ValidateItem(nil, api.technology, &dtapi.ValidateItemOpts{
		CustomService: optional.NewInterface(service),
	})
	return err
}
