package main

import (
	"encoding/json"
	"fmt"

	"github.com/dtcookie/dtapi/dynatrace"
)

// ToString TODO: documentation
func ToString(v interface{}) string {
	bytes, err := json.MarshalIndent(v, "", "    ")
	if err != nil {
		return ""
	}
	return string(bytes)
}

func main() {

	credentials := Credentials{}
	credentials.Read("credentials.json")

	tenant := dynatrace.NewTenant(credentials.Environment, credentials.APIToken)
	version, err := tenant.APIs.Cluster.Version()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(version)

	clusterTime, err := tenant.APIs.Cluster.ClusterTime()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(clusterTime)

	webAppConfig, err := tenant.APIs.WebApplications.Default.Fetch()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(ToString(webAppConfig))

	webAppConfig.XhrActionKeyPerformanceMetric = "asdfasdf"
	isValid, err := tenant.APIs.WebApplications.Default.IsValid(webAppConfig)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(fmt.Sprintf("isValid: %t", isValid))

	builder := tenant.APIs.UserSessionQL.AsTree("")
	result, err := builder.Fetch()
	fmt.Println(result)

	asdf := tenant.APIs.CustomServices.Java.ForID("")
	fmt.Println(asdf)
}
