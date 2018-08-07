package main

import (
	"fmt"
	"strings"

	"code.cloudfoundry.org/cli/plugin"
	"code.cloudfoundry.org/cli/plugin/models"
	log "github.com/dtcookie/dtapi/libdtlog"
)

// Flags TODO: documentation
type Flags struct {
	Name  string
	Force bool
}

// Credentials TODO: documentation
type Credentials struct {
	Environment string
	APIToken    string
}

// BaseParams TODO: documentation
type BaseParams struct {
	Credentials Credentials
	Flags       Flags
}

// NewBaseParams TODO: documentation
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

// WebAppParams TODO: documentation
type WebAppParams struct {
	BaseParams
	WebAppName string
}

// NewWebAppParams TODO: documentation
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

func evalWebAppName(params *WebAppParams, args []string) ([]string, bool) {
	appName, args := argp.consume(args)
	params.WebAppName = appName
	if appName == "" {
		return args, false
	}
	return args, true
}

// PCF TODO: documentation
type PCF struct {
	App   string
	Org   string
	Space string
	User  string
}

// PCFParams TODO: documentation
type PCFParams struct {
	BaseParams
	PCF        PCF
	WebAppName string
}

func resolveVars(name string, pcf PCF) string {
	name = strings.Replace(name, "{cf:app}", pcf.App, -1)
	name = strings.Replace(name, "{cf:space}", pcf.Space, -1)
	name = strings.Replace(name, "{cf:org}", pcf.Org, -1)
	return name
}

// NewPCFParams TODO: documentation
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

// PCFDomainParams TODO: documentation
type PCFDomainParams struct {
	PCFParams
	Domains []string
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

// NewPCFDomainParams TODO: documentation
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
