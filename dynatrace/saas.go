package dynatrace

import (
	"github.com/dtcookie/dtapi/config"
	"github.com/dtcookie/dtapi/environment"
)

func newSaasEnvClient(env string, apiToken string) *environment.APIClient {
	configuration := environment.NewConfiguration()
	configuration.BasePath = "https://" + env + ".live.dynatrace.com/api/v1"
	configuration.AddDefaultHeader("Authorization", "Api-Token "+apiToken)
	return environment.NewAPIClient(configuration)
}

func newSaasConfigClient(env string, apiToken string) *config.APIClient {
	configuration := config.NewConfiguration()
	configuration.BasePath = "https://" + env + ".live.dynatrace.com/api/config/v1"
	configuration.AddDefaultHeader("Authorization", "Api-Token "+apiToken)
	return config.NewAPIClient(configuration)
}
