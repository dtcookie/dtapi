package dtapi

import (
	dtapiconf "github.com/dtcookie/dtapi/libdtapiconf"
	dtapienv "github.com/dtcookie/dtapi/libdtapienv"
)

func newSaasEnvClient(env string, apiToken string) *dtapienv.APIClient {
	configuration := dtapienv.NewConfiguration()
	configuration.BasePath = "https://" + env + ".live.dynatrace.com/api/v1"
	configuration.AddDefaultHeader("Authorization", "Api-Token "+apiToken)
	return dtapienv.NewAPIClient(configuration)
}

func newSaasConfigClient(env string, apiToken string) *dtapiconf.APIClient {
	configuration := dtapiconf.NewConfiguration()
	configuration.BasePath = "https://" + env + ".live.dynatrace.com/api/config/v1"
	configuration.AddDefaultHeader("Authorization", "Api-Token "+apiToken)
	return dtapiconf.NewAPIClient(configuration)
}
