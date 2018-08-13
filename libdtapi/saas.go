package dtapi

import (
	dtapiconf "github.com/dtcookie/dtapi/libdtapiconf"
	dtapienv "github.com/dtcookie/dtapi/libdtapienv"
)

func newSaasEnvClient(env string, apiToken string) *dtapienv.APIClient {
	configuration := dtapienv.NewConfiguration()
	configuration.BasePath = "https://" + env + ".live.dynatrace.com/api/v1"
	configuration.AddDefaultHeader("Authorization", "Api-Token "+apiToken)
	cl := dtapienv.NewAPIClient(configuration)
	configuration.HTTPClient.Jar = cookieJar
	return cl
}

func newSaasConfigClient(env string, apiToken string) *dtapiconf.APIClient {
	configuration := dtapiconf.NewConfiguration()
	configuration.BasePath = "https://" + env + ".live.dynatrace.com/api/config/v1"
	configuration.AddDefaultHeader("Authorization", "Api-Token "+apiToken)
	cl := dtapiconf.NewAPIClient(configuration)
	configuration.HTTPClient.Jar = cookieJar
	return cl
}
