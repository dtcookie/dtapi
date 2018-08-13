package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	"code.cloudfoundry.org/cli/cf/manifest"
	"code.cloudfoundry.org/cli/plugin"
	dtcmd "github.com/dtcookie/dtapi/libdtcmd"
	log "github.com/dtcookie/dtapi/libdtlog"
	"github.com/fatih/color"
	"gopkg.in/yaml.v2"
)

// Version is the version of the plugin.
var Version = plugin.VersionType{
	Major: 1,
	Minor: 0,
	Build: 0,
}

// MinCliVersion is the minimum version of the CF CLI
// command line utility.
var MinCliVersion = plugin.VersionType{
	Major: 6,
	Minor: 7,
	Build: 0,
}

// CfCliDynatrace is the class of this plugin
type CfCliDynatrace struct{}

// NewTenant creates a new Tenant for accessing the Dynatrace REST API
//
// Parameters:
// - credentials			for authenticating against the Dynatrace Tenant
// - checkClusterVersion	if 'true' an initial check against the
// 							GetClusterVersion API is being performed.
func NewTenant(credentials dtcmd.Credentials, checkClusterVersion bool) (*dtcmd.Tenant, error) {
	tenant := dtcmd.NewTenant(credentials)

	if checkClusterVersion {
		if err := tenant.PrintClusterVersion(); err != nil {
			return nil, err
		}
	}
	return tenant, nil
}

// NewTenantPCF creates a new Tenant for accessing the Dynatrace REST API
//
// Parameters:
// - credentials			for authenticating against the Dynatrace Tenant
// - pcf					details about the PCF environment the user is currently
//							logged into. This is information is being used for
//							logging purposes.
// - checkClusterVersion	if 'true' an initial check against the
// 							GetClusterVersion API is being performed.
func NewTenantPCF(credentials dtcmd.Credentials, pcf PCF, checkClusterVersion bool) (*dtcmd.Tenant, error) {
	tenant := dtcmd.NewTenant(credentials)
	tenant.ActivityCallback = &cfActivityCallback{
		cfApp:   pcf.App,
		cfSpace: pcf.Space,
		cfOrg:   pcf.Org,
		cfUser:  pcf.User,
	}

	if checkClusterVersion {
		if err := tenant.PrintClusterVersion(); err != nil {
			return nil, err
		}
	}
	return tenant, nil
}

// ListWebApplications prints out a tabular list of the currently
// existing Web Applications in Dynatrace.
//
// Command Line Syntax:
//		cf dynatrace list
func (c *CfCliDynatrace) ListWebApplications(args []string) bool {
	var err error
	var tenant *dtcmd.Tenant
	var params *BaseParams

	if params = NewBaseParams(args); params == nil {
		return false
	}

	if tenant, err = NewTenant(params.Credentials, false); err != nil {
		return log.FailError(err)
	}
	if err = tenant.ListWebApplications(); err != nil {
		return log.FailError(err)
	}

	return true
}

// DeleteWebApplication deletes a Web Application in Dynatrace
//
// Command Line Syntax:
//		cf dynatrace delete <webapp>
func (c *CfCliDynatrace) DeleteWebApplication(args []string) bool {
	var err error
	var tenant *dtcmd.Tenant
	var params *WebAppParams

	if params = NewWebAppParams(args); params == nil {
		return false
	}

	if tenant, err = NewTenant(params.Credentials, false); err != nil {
		return log.FailError(err)
	}

	if err = tenant.DeleteNamedWebApplication(params.WebAppName); err != nil {
		return log.FailError(err)
	}
	return true
}

type Manifest struct {
	Applications []manifest.Application `yaml:"applications,omitempty"`
}

func contains(slice []string, item string) bool {
	for _, elem := range slice {
		if elem == item {
			return true
		}
	}
	return false
}

func handleAppStrProp(k string, v interface{}) {
	switch k {
	case "memory":
		fmt.Println(k + " found")
	case "name":
		fmt.Println(k + " found")
	case "path":
		fmt.Println(k + " found")
	case "random-route":
		fmt.Println(k + " found")
	case "services":
		fmt.Println(k + " found")
	default:
		fmt.Println(k + " ignored")
		return
	}
}

func handleAppProp(k interface{}, v interface{}) {
	switch tk := k.(type) {
	case string:
		handleAppStrProp(tk, v)
	default:
		return
	}
}

func handleApplication(v interface{}) {
	switch a := v.(type) {
	case map[interface{}]interface{}:
		for k, v := range a {
			handleAppProp(k, v)
		}
	default:
		fmt.Printf("unexpected type %T\n", a)
		return
	}

}

func handleApplications(v interface{}) {
	switch a := v.(type) {
	case []interface{}:
		for i := range a {
			handleApplication(a[i])
		}
	default:
		return
	}
}

func handleManifestStrProp(k string, v interface{}) {
	switch k {
	case "applications":
		handleApplications(v)
	default:
		return
	}
}

func handleManifestProp(k interface{}, v interface{}) {
	switch tk := k.(type) {
	case string:
		handleManifestStrProp(tk, v)
	default:
		return
	}
}

func cast(v interface{}) map[interface{}][]interface{} {
	switch w := v.(type) {
	case map[interface{}][]interface{}:
		return w
	default:
		fmt.Println(fmt.Sprintf("unknown %T", w))
		return nil
	}
}

func (c *CfCliDynatrace) Push(cli plugin.CliConnection, args []string) bool {
	var err error
	var output []string
	var pushArguments *PushArguments
	// var mani *manifest.Manifest

	// var userGUID string

	// if userGUID, err = cli.UserGuid(); err != nil {
	// 	log.FailError(err)
	// 	return false
	// }

	// metaDataServiceName := "dynatrace-metadata-" + userGUID
	if pushArguments, err = ParsePushFlags(args); err != nil {
		log.FailError(err)
		return false
	}

	var bytes []byte

	if bytes, err = ioutil.ReadFile(pushArguments.Manifest); err != nil {
		log.FailError(err)
		return false
	}

	// var mani Manifest
	mmap := make(map[interface{}]interface{})

	if err = yaml.Unmarshal(bytes, &mmap); err != nil {
		log.FailError(err)
		return false
	}

	applications := cast(mmap["applications"])

	//var application map[interface{}]interface{}
	for k, application := range applications {
		fmt.Println(fmt.Sprintf("key: %s", k))
		fmt.Println(fmt.Sprintf("val: %s", application))
	}

	return true

	// for k, v := range mmap {
	// 	handleManifestProp(k, v)
	// }

	// var applications []manifest.Application
	// applications = mani.Applications
	// var application *manifest.Application
	// for i := range applications {
	// 	application = &applications[i]
	// 	if !contains(application.Services, metaDataServiceName) {
	// 		application.Services = append(application.Services, metaDataServiceName)
	// 	}
	// }
	if bytes, err = yaml.Marshal(mmap); err != nil {
		log.FailError(err)
		return false
	}
	fmt.Println(string(bytes))

	// str := string(b) // convert content to a 'string'

	// fmt.Println(str) // print the content as a 'string'

	return true

	pushArgs := append(args, "push")
	if output, err = cli.CliCommand(pushArgs...); err != nil {
		log.FailError(err)
		return false
	}
	for _, line := range output {
		fmt.Println(line)
	}
	return true
	// var err error
	// var tenant *dtcmd.Tenant
	// var params *PCFDomainParams

	// if params = NewPCFDomainParams(cli, args); params == nil {
	// 	return false
	// }

	// if tenant, err = NewTenantPCF(params.Credentials, params.PCF, false); err != nil {
	// 	return log.FailError(err)
	// }

	// if _, err = tenant.BindWebApplication(params.WebAppName, params.Domains, params.PCF.User); err != nil {
	// 	return log.FailError(err)
	// }

	return true
}

// Bind creates or updates a Dynatrace Web Application based on the
// routes a PCF Application can get reached at.
//
// Command Line Syntax:
//		cf dynatrace bind <app> [-name <webapp>]
//
// 			<app>		the name of the PCF application.
//			<webapp>	the name Dynatrace Web Application.
//							May contain the placeholders
//								{cf:app}, {cf:space} and {cf:org}
//							If not specified the name of the PCF Application
//							is being used.
func (c *CfCliDynatrace) Bind(cli plugin.CliConnection, args []string) bool {
	var err error
	var tenant *dtcmd.Tenant
	var params *PCFDomainParams

	if params = NewPCFDomainParams(cli, args); params == nil {
		return false
	}

	if tenant, err = NewTenantPCF(params.Credentials, params.PCF, false); err != nil {
		return log.FailError(err)
	}

	if _, err = tenant.BindWebApplication(params.WebAppName, params.Domains, params.PCF.User); err != nil {
		return log.FailError(err)
	}

	return true
}

// Unbind removes the Application Detection Rules of a Dynatrace Web Application
// that are matching with the corresponding PCF Application.
// The Dynatrace Web Application won't get deleted.
//
// Command Line Syntax:
//		cf dynatrace unbind <app> [-name <webapp>]
//
// 			<app>		the name of the PCF application.
//			<webapp>	the name Dynatrace Web Application.
//							May contain the placeholders
//								{cf:app}, {cf:space} and {cf:org}
//							If not specified the name of the PCF Application
//							is being used.
func (c *CfCliDynatrace) Unbind(cliConnection plugin.CliConnection, args []string) bool {
	var err error
	var tenant *dtcmd.Tenant
	var params *PCFDomainParams

	if params = NewPCFDomainParams(cliConnection, args); params == nil {
		return false
	}

	if tenant, err = NewTenantPCF(params.Credentials, params.PCF, false); err != nil {
		return log.FailError(err)
	}

	if _, err = tenant.UnBindWebApplication(params.WebAppName, params.Domains); err != nil {
		return log.FailError(err)
	}

	return true
}

type cfActivityCallback struct {
	dtcmd.ActivityCallback
	cfApp   string
	cfOrg   string
	cfSpace string
	cfUser  string
}

// OnCreateWebApp prints out, in addition to the default activity callback of a tenant,
// the PCF app, org and space when a web application is getting created.
func (callback *cfActivityCallback) OnCreateWebApp(name string) {
	log.Println("... creating web application ", log.Cyan(name), " for app ", log.Cyan(callback.cfApp), " in org ", log.Cyan(callback.cfOrg), " / space ", log.Cyan(callback.cfSpace), " as ", log.Cyan(callback.cfUser))
}

// Run this is how the plugin is actually getting invoked.
func (c *CfCliDynatrace) Run(cliConnection plugin.CliConnection, args []string) {
	// Ensure that we called the command dynatrace
	command, args := argp.consume(args)

	if command != "dynatrace" {
		return
	}

	isLoggedIn, err := cliConnection.IsLoggedIn()
	if err != nil {
		log.Fail(err.Error())
		return
	}
	if !isLoggedIn {
		log.Fail("Not logged in. Use '", log.Red("cf login"), "' to log in.")
		return
	}

	isOk := false

	subCommand, args := argp.consume(args)

	switch subCommand {
	case "bind":
		isOk = c.Bind(cliConnection, args)
	case "unbind":
		isOk = c.Unbind(cliConnection, args)
	case "delete":
		isOk = c.DeleteWebApplication(args)
	case "list":
		isOk = c.ListWebApplications(args)
	case "push":
		isOk = c.Push(cliConnection, args)
	default:
		log.Println("Incorrect Usage: the required argument <operation> was not provided")
		log.Println("   Valid argument values for the operation are: ", log.Cyan("bind"), " | ", log.Cyan("unbind"), " | ", log.Cyan("delete"), " | ", log.Cyan("list"))
		log.Fail()
		fmt.Println("USAGE:")
		fmt.Println("   " + c.GetMetadata().Commands[0].UsageDetails.Usage)
	}

	if isOk {
		color.New(color.FgGreen, color.Bold).Println("OK")
		// } else {
		// 	color.New(color.FgRed, color.Bold).Println("FAILED")
	}
}

// GetMetadata Metadata about the plugin
func (c *CfCliDynatrace) GetMetadata() plugin.PluginMetadata {
	// usage := []string{
	// 	"cf dynatrace bind <app> [-name <webapp>] [-environment <environmentID>] [-api-token <api-token>]",
	// 	"         to create or update a web application in dynatrace to cover the routes of the given cloudfoundry application",
	// 	"      where",
	// 	"         <app>    is the name of an application deployed in cloudfoundry",
	// 	"         <webapp> is the name of the web application in dynatrace to create of update",
	// 	"            which may optionally contain the placeholders '{cf:app}', '{cf:space}' and '{cf:org}'",
	// 	"            referring to the cloudfoundry application, space and org",
	// 	"   cf dynatrace unbind <app> [-name <webapp>] [-environment <environmentID>] [-api-token <api-token>]",
	// 	"         to remove the application detection rule that binds the given web application to the given route",
	// 	"      where",
	// 	"         <app>    is the name of an application deployed in cloudfoundry",
	// 	"         <webapp> is the name of the web application in dynatrace to create of update",
	// 	"            which may optionally contain the placeholders '{cf:app}', '{cf:space}' and '{cf:org}'",
	// 	"            referring to the cloudfoundry application, space and org",
	// 	"   cf dynatrace delete <webapp> [-environment <environmentID>] [-api-token <api-token>]",
	// 	"      where <webapp> is a web application in dynatrace",
	// 	"   cf dynatrace list [-environment <environmentID>] [-api-token <api-token>]",
	// }
	usage := []string{
		"",
		"   cf dynatrace bind <app> [-name <webapp>] [-environment <environmentID>] [-api-token <api-token>]",
		"     to create or update a web application in dynatrace to cover the routes of the given cloudfoundry application",
		"",
		"   cf dynatrace unbind <app> [-name <webapp>] [-environment <environmentID>] [-api-token <api-token>]",
		"     to remove the application detection rule that binds the given web application to the given route",
		"",
		"   cf dynatrace delete <webapp> [-environment <environmentID>] [-api-token <api-token>]",
		"     to delete a web application (including the application detection rules)",
		"",
		"   cf dynatrace list [-environment <environmentID>] [-api-token <api-token>]",
		"     to list the currently configured web applications",
		"",
		"   The arguments -environment and -api-token are optional if the environment variables",
		"       DT_ENVIRONMENT and DT_APITOKEN",
		"   are defined.",
	}
	usageStr := strings.Join(usage, "\n")

	return plugin.PluginMetadata{
		Name:          "dynatrace",
		Version:       Version,
		MinCliVersion: MinCliVersion,
		Commands: []plugin.Command{
			{
				Name:     "dynatrace",
				HelpText: "Manages Dynatrace Web Applications for apps deployed in cloudfoundry",
				UsageDetails: plugin.Usage{
					Usage: usageStr,
				},
			},
		},
	}
}

func main() {
	plugin.Start(new(CfCliDynatrace))
}
