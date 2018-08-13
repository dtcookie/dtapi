package dtapi

import dtapi "github.com/dtcookie/dtapi/libdtapienv"

type managedTenant struct {
	apiToken    string
	environment string
}

func (tenant managedTenant) envClient() *dtapi.APIClient {
	configuration := dtapi.NewConfiguration()
	configuration.BasePath = "https://" + tenant.environment + ".live.dynatrace.com/api/v1"
	configuration.AddDefaultHeader("Authorization", "Api-Token "+tenant.apiToken)
	cl := dtapi.NewAPIClient(configuration)
	configuration.HTTPClient.Jar = cookieJar
	return cl
}
