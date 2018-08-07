package main

import (
	"fmt"
	"strings"

	"code.cloudfoundry.org/cli/plugin"
	dtcmd "github.com/dtcookie/dtapi/libdtcmd"
	log "github.com/dtcookie/dtapi/libdtlog"
	"github.com/fatih/color"
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
func NewTenant(credentials Credentials, checkClusterVersion bool) (*dtcmd.Tenant, error) {
	tenant := dtcmd.NewTenant(credentials.Environment, credentials.APIToken)

	if checkClusterVersion {
		if err := tenant.CheckClusterVersion(); err != nil {
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
func NewTenantPCF(credentials Credentials, pcf PCF, checkClusterVersion bool) (*dtcmd.Tenant, error) {
	tenant := dtcmd.NewTenant(credentials.Environment, credentials.APIToken)
	tenant.ActivityCallback = &cfActivityCallback{
		cfApp:   pcf.App,
		cfSpace: pcf.Space,
		cfOrg:   pcf.Org,
		cfUser:  pcf.User,
	}

	if checkClusterVersion {
		if err := tenant.CheckClusterVersion(); err != nil {
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
	if err = tenant.WebApplications.List(); err != nil {
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

	if tenant.WebApplications.NamedDelete(params.WebAppName) != nil {
		return tenant.LogLastErr()
	}
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

	if _, err = tenant.BindWebApplication(params.WebAppName, params.Domains); err != nil {
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
	dtcmd.DefaultActivityCallback
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
	usage := []string{
		"cf dynatrace bind <app> [-name <webapp>] [-environment <environmentID>] [-api-token <api-token>]",
		"         to create or update a web application in dynatrace to cover the routes of the given cloudfoundry application",
		"      where",
		"         <app>    is the name of an application deployed in cloudfoundry",
		"         <webapp> is the name of the web application in dynatrace to create of update",
		"            which may optionally contain the placeholders '{cf:app}', '{cf:space}' and '{cf:org}'",
		"            referring to the cloudfoundry application, space and org",
		"   cf dynatrace unbind <app> [-name <webapp>] [-environment <environmentID>] [-api-token <api-token>]",
		"         to remove the application detection rule that binds the given web application to the given route",
		"      where",
		"         <app>    is the name of an application deployed in cloudfoundry",
		"         <webapp> is the name of the web application in dynatrace to create of update",
		"            which may optionally contain the placeholders '{cf:app}', '{cf:space}' and '{cf:org}'",
		"            referring to the cloudfoundry application, space and org",
		"   cf dynatrace delete <webapp> [-environment <environmentID>] [-api-token <api-token>]",
		"      where <webapp> is a web application in dynatrace",
		"   cf dynatrace list [-environment <environmentID>] [-api-token <api-token>]",
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
