package main

import (
	"fmt"
	"strings"

	"code.cloudfoundry.org/cli/plugin"
	"code.cloudfoundry.org/cli/plugin/models"
	log "github.com/dtcookie/dtapi/libdtlog"
)

// Flags optional parameters that are specified using
// the '-flag value' syntax.
//
// In general these parameters should be optional arguments
type Flags struct {
	// Name will be populated if the argument
	// '-name <value>' has been specified.
	//
	// A usage example is to explicitly specify the name
	// of the web application to create instead of relying
	// on the fallback, which is the name of the PCF Application
	Name string
	// Force is a switch, specified with the argument '-force'.
	// A usage example is to enforce automatic removal of
	// application detection rules
	Force bool
}

// Credentials hold the necessary information in order to access
// the REST API of a Dynatrace Tenant
type Credentials struct {
	// Environment specifies the unique identifier of a
	// Dynatrace Tenant.
	// For a SaaS Tenant, this is the first part of the Web URL
	// Example: https://<environment>.live.dynatrace.com
	Environment string
	// APIToken specifies the access token to be used to
	// authenticate against the Dynatrace REST API.
	APIToken string
}

// BaseParams hold the credentials to authenticate
// against the Dynatrace REST API and collect optional
// flags
type BaseParams struct {
	// Credentials for authenticating against the
	// Dynatrace REST API
	Credentials Credentials
	// Flags are optional arguments, but are not
	// necessarily required in order for a command
	// to execute
	Flags Flags
}

// WebAppParams hold the same information as BaseParams
// In addition it holds the name of a Dynatrace Web
// Application.
type WebAppParams struct {
	BaseParams
	//
	WebAppName string
}

// PCF holds information about the PCF environment the
// user is required to be logged into.
type PCF struct {
	// App is the PCF App to work query information from.
	App string
	// Org is the PCF Org the PCF App belongs to.
	Org string
	// Space is the Space Org the PCF App belongs to.
	Space string
	// User is the currently logged in user.
	User string
}

// PCFParams hold the same information as BaseParams
//
// Furthermore the fully populated object provides
// information about the PCF environment the user is
// required to be logged into.
//
// In addition it holds the resolved name of the Dynatrace
// Web Application the operation is supposed to run against.
type PCFParams struct {
	BaseParams
	// information about the PCF environment the
	// user is required to be logged into.
	PCF PCF
	// The name of the resolved Dynatrace Web Application
	// the operation is supposed to run against.
	//
	// If not explicitly specified the name of the Web
	// Application is identical to the PCF Application
	// (which is a mandatory argument).
	// If specified via '-name <webapp>' the value may
	// contain the placeholders '{cf:app}', '{cf:space}'
	// and '{cf:org}' which will get automatically resolved.
	WebAppName string
}

// PCFDomainParams hold the same information as PCFParams
//
// In addition it provides the domain names the PCF
// Application can be reached at.
type PCFDomainParams struct {
	PCFParams
	// Domains are the fully qualified domain names
	// the PCF Application can be reached at.
	Domains []string
}

// NewBaseParams creates a new BaseParams object and
// resolves the required command line arguments as well
// as optional flags.
//
// Required arguments are:
// * -environment <environment>
//   - unless available as environment variable 'DT_ENVIRONMENT'
// * -api-token <api-token>
//   - unless available as environment variable 'DT_APITOKEN'
//
// In case the required arguments cannot get resolved this function
// returns 'nil' after printing out an error message.
func NewBaseParams(args []string) *BaseParams {
	var err error
	params := BaseParams{}

	cmdflags := commandFlags{}

	if err = cmdflags.parse(args); err != nil {
		log.FailError(err)
		return nil
	}
	if !cmdflags.validate() {
		return nil
	}

	params.Credentials.APIToken = cmdflags.APIToken
	params.Credentials.Environment = cmdflags.Environment

	params.Flags.Name = cmdflags.Name
	params.Flags.Force = cmdflags.Force

	return &params
}

// NewWebAppParams creates a new WebAppParams object and
// resolves the required command line arguments as well
// as optional flags.
//
// Required arguments are:
// * -environment <environment>
//   - unless available as environment variable 'DT_ENVIRONMENT'
// * -api-token <api-token>
//   - unless available as environment variable 'DT_APITOKEN'
// * <webappname>
//   - the name of the Dynatrace Web Application the operation
//     is supposed to work against.
//
// In case the required arguments cannot get resolved this function
// returns 'nil' after printing out an error message.
func NewWebAppParams(args []string) *WebAppParams {
	baseParams := NewBaseParams(args)
	if baseParams == nil {
		return nil
	}
	params := WebAppParams{}

	params.Credentials = baseParams.Credentials
	params.Flags = baseParams.Flags

	var ok bool

	if args, ok = evalWebAppName(&params, args); !ok {
		log.Fail()
		log.Println(log.Red("expected:"), " delete ", log.Cyan("<webapp>"))
		log.Println("  where ", log.Cyan("<webapp>"), " is the name of your dynatrace web application")
		return nil
	}
	return &params
}

// NewPCFParams creates a new PCFParams object and
// resolves the required command line arguments as well
// as optional flags.
//
// Required arguments are:
// * <app>
//   - the name of the PCF Application to work with
// * -environment <environment>
//   - unless available as environment variable 'DT_ENVIRONMENT'
// * -api-token <api-token>
//   - unless available as environment variable 'DT_APITOKEN'
//
// Optional arguments are:
// * -namme <webapp>
//   - the name of the Dynatrace Web Application the operation
//     is supposed to work against.
//   - If not explicitly specified the name of the Web
//     Application is identical to the PCF Application
//     (which is a mandatory argument).
//     If specified via '-name <webapp>' the value may
//     contain the placeholders '{cf:app}', '{cf:space}'
//     and '{cf:org}' which will get automatically resolved.
//
// In case the required arguments cannot get resolved this function
// returns 'nil' after printing out an error message.
func NewPCFParams(cli plugin.CliConnection, args []string) *PCFParams {
	baseParams := NewBaseParams(args)
	if baseParams == nil {
		return nil
	}
	params := PCFParams{}

	params.Credentials = baseParams.Credentials
	params.Flags = baseParams.Flags

	var err error
	var ok bool

	if args, ok = evalAppName(&params, args); !ok {
		return nil
	}
	if !evalOrg(&params, cli) {
		return nil
	}
	if !evalSpace(&params, cli) {
		return nil
	}
	if !checkAppExists(&params, cli) {
		return nil
	}

	if params.PCF.User, err = cli.Username(); err != nil {
		log.FailError(err)
		return nil
	}

	params.WebAppName = params.PCF.App

	if params.Flags.Name != "" {
		params.WebAppName = resolveVars(params.Flags.Name, params.PCF)

	}
	return &params
}

// NewPCFDomainParams creates a new PCFDomainParams object and
// resolves the required command line arguments as well
// as optional flags.
//
// In addition it also resolves the routes the PCF Application
// can get reached at.
//
// Required arguments are:
// * <app>
//   - the name of the PCF Application to work with
// * -environment <environment>
//   - unless available as environment variable 'DT_ENVIRONMENT'
// * -api-token <api-token>
//   - unless available as environment variable 'DT_APITOKEN'
//
// Optional arguments are:
// * -namme <webapp>
//   - the name of the Dynatrace Web Application the operation
//     is supposed to work against.
//   - If not explicitly specified the name of the Web
//     Application is identical to the PCF Application
//     (which is a mandatory argument).
//     If specified via '-name <webapp>' the value may
//     contain the placeholders '{cf:app}', '{cf:space}'
//     and '{cf:org}' which will get automatically resolved.
//
// In case the required arguments cannot get resolved this function
// returns 'nil' after printing out an error message.
func NewPCFDomainParams(cli plugin.CliConnection, args []string) *PCFDomainParams {
	baseParams := NewPCFParams(cli, args)
	if baseParams == nil {
		return nil
	}
	params := PCFDomainParams{}

	params.Credentials = baseParams.Credentials
	params.Flags = baseParams.Flags

	params.PCF = baseParams.PCF

	if resolveDomains(&params, cli) {
		return &params
	}

	return nil
}

func evalWebAppName(params *WebAppParams, args []string) ([]string, bool) {
	appName, args := argp.consume(args)
	params.WebAppName = appName
	if appName == "" {
		return args, false
	}
	return args, true
}

func resolveVars(name string, pcf PCF) string {
	name = strings.Replace(name, "{cf:app}", pcf.App, -1)
	name = strings.Replace(name, "{cf:space}", pcf.Space, -1)
	name = strings.Replace(name, "{cf:org}", pcf.Org, -1)
	return name
}

func evalAppName(params *PCFParams, args []string) ([]string, bool) {
	appName, args := argp.consume(args)
	params.PCF.App = appName
	if appName == "" {
		log.Fail()
		log.Println(log.Red("expected:"), " bind ", log.Cyan("<app>"))
		log.Println("  where ", log.Cyan("<app>"), " is the name of your cloudfoundy application")
		return args, false
	}
	return args, true
}

func evalOrg(params *PCFParams, cli plugin.CliConnection) bool {
	org, _ := cli.GetCurrentOrg()
	params.PCF.Org = org.Name
	if org.Name == "" {
		log.Fail("No org and space targeted, use '", log.Red("cf target -o ORG -s SPACE"), "' to target an org and space.")
		return false
	}
	return true
}

func evalSpace(params *PCFParams, cli plugin.CliConnection) bool {
	space, _ := cli.GetCurrentSpace()
	params.PCF.Space = space.Name
	if space.Name == "" {
		log.Fail("No space targeted, use '", log.Red("cf target -s"), "' to target a space.")
		return false
	}
	return true
}

func checkAppExists(params *PCFParams, cli plugin.CliConnection) bool {
	var err error
	// var appModel plugin_models.GetAppModel

	if _, err = cli.GetApp(params.PCF.App); err == nil {
		// params.appModel = appModel
		return true
	}
	if strings.HasSuffix(err.Error(), "not found") {
		return log.Fail("App ", log.Cyan(params.PCF.App), " not found in org ", log.Cyan(params.PCF.Org), " / space ", log.Cyan(params.PCF.Space))
	}
	return log.Fail(err.Error())
}

func resolveDomains(params *PCFDomainParams, cli plugin.CliConnection) bool {
	var appModel plugin_models.GetAppModel
	var err error

	if appModel, err = cli.GetApp(params.PCF.App); err != nil {
		return false
	}

	domains := make([]string, 0)
	for _, route := range appModel.Routes {
		domain := fmt.Sprintf("%s.%s", route.Host, route.Domain.Name)
		domains = append(domains, domain)
	}
	params.Domains = domains
	return true
}
