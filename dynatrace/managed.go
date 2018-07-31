package dynatrace

import "github.com/dtcookie/dtapi/environment"

type managedTenant struct {
	apiToken    string
	environment string
}

func (tenant managedTenant) envClient() *environment.APIClient {
	configuration := environment.NewConfiguration()
	configuration.BasePath = "https://" + tenant.environment + ".live.dynatrace.com/api/v1"
	configuration.AddDefaultHeader("Authorization", "Api-Token "+tenant.apiToken)
	return environment.NewAPIClient(configuration)
}
