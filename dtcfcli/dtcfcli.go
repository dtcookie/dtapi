package main

import (
	"flag"
	"fmt"

	"code.cloudfoundry.org/cli/plugin"
	"github.com/fatih/color"
)

// CfCliDynatrace is the struct implementing the interface defined by the core CLI. It can
// be found at  "code.cloudfoundry.org/cli/plugin/plugin.go"
type CfCliDynatrace struct{}

// Usage Usage
func (c *CfCliDynatrace) Usage() {
}

// ApplicationRule ApplicationRule
func (c *CfCliDynatrace) ApplicationRule(cliConnection plugin.CliConnection, args []string) bool {
	for _, arg := range args {
		fmt.Println("arg: " + arg)
	}

	appName := ""
	flagSet := flag.NewFlagSet("dynatrace", flag.ExitOnError)
	flagSet.StringVar(&appName, "cfapp", "", "the name of the cloudfoundry app")
	flagSet.Parse(args[1:])
	fmt.Println("appName: " + appName)

	appModels, err := cliConnection.GetApps()
	if err != nil {
		fmt.Println(err)
	} else {
		for _, appModel := range appModels {
			fmt.Println(appModel.Name)
			for _, route := range appModel.Routes {
				domain := fmt.Sprintf("%s.%s", route.Host, route.Domain.Name)
				fmt.Println("   " + domain)
			}
		}
	}

	return true
}

// Run must be implemented by any plugin because it is part of the
// plugin interface defined by the core CLI.
//
// Run(....) is the entry point when the core CLI is invoking a command defined
// by a plugin. The first parameter, plugin.CliConnection, is a struct that can
// be used to invoke cli commands. The second paramter, args, is a slice of
// strings. args[0] will be the name of the command, and will be followed by
// any additional arguments a cli user typed in.
//
// Any error handling should be handled with the plugin itself (this means printing
// user facing errors). The CLI will exit 0 if the plugin exits 0 and will exit
// 1 should the plugin exits nonzero.
func (c *CfCliDynatrace) Run(cliConnection plugin.CliConnection, args []string) {

	// Ensure that we called the command dynatrace
	command, args := argp.consume(args)

	if command != "dynatrace" {
		return
	}

	fmt.Println("Running the dynatrace plugin")

	isOk := false

	subCommand, args := argp.consume(args)

	switch subCommand {
	case "application-rule":
		isOk = c.ApplicationRule(cliConnection, args)
	default:
		fmt.Print("expected: ")
		colors.cyan("application-rule")
		fmt.Println("")
	}

	if isOk {
		color.New(color.FgGreen, color.Bold).Println("OK")
	} else {
		color.New(color.FgRed, color.Bold).Println("FAILED")
	}
}

// GetMetadata must be implemented as part of the plugin interface
// defined by the core CLI.
//
// GetMetadata() returns a PluginMetadata struct. The first field, Name,
// determines the name of the plugin which should generally be without spaces.
// If there are spaces in the name a user will need to properly quote the name
// during uninstall otherwise the name will be treated as seperate arguments.
// The second value is a slice of Command structs. Our slice only contains one
// Command Struct, but could contain any number of them. The first field Name
// defines the command `cf basic-plugin-command` once installed into the CLI. The
// second field, HelpText, is used by the core CLI to display help information
// to the user in the core commands `cf help`, `cf`, or `cf -h`.
func (c *CfCliDynatrace) GetMetadata() plugin.PluginMetadata {
	return plugin.PluginMetadata{
		Name: "dynatrace",
		Version: plugin.VersionType{
			Major: 1,
			Minor: 0,
			Build: 0,
		},
		MinCliVersion: plugin.VersionType{
			Major: 6,
			Minor: 7,
			Build: 0,
		},
		Commands: []plugin.Command{
			{
				Name:     "dynatrace",
				HelpText: "dynatrace's help text",

				// UsageDetails is optional
				// It is used to show help of usage of each command
				UsageDetails: plugin.Usage{
					Usage: "dynatrace\n   cf dynatrace",
				},
			},
		},
	}
}

// Unlike most Go programs, the `Main()` function will not be used to run all of the
// commands provided in your plugin. Main will be used to initialize the plugin
// process, as well as any dependencies you might require for your
// plugin.
func main() {
	// Any initialization for your plugin can be handled here
	//
	// Note: to run the plugin.Start method, we pass in a pointer to the struct
	// implementing the interface defined at "code.cloudfoundry.org/cli/plugin/plugin.go"
	//
	// Note: The plugin's main() method is invoked at install time to collect
	// metadata. The plugin will exit 0 and the Run([]string) method will not be
	// invoked.
	plugin.Start(new(CfCliDynatrace))
	// Plugin code should be written in the Run([]string) method,
	// ensuring the plugin environment is bootstrapped.
}
