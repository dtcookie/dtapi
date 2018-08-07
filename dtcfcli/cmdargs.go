package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"strings"

	log "github.com/dtcookie/dtapi/libdtlog"
)

type argstype struct {
}

var argp = argstype{}

func (argv argstype) consume(args []string) (string, []string) {
	if len(args) == 0 {
		return "", args
	}
	return args[0], args[1:]
}

func dumpArgs(method string, args []string) {
	fmt.Println("Arguments within '" + method + "':")
	for _, arg := range args {
		fmt.Println("    " + arg)
	}
}

type commandFlags struct {
	Environment string
	APIToken    string
	Name        string
	Force       bool
}

func (cmdFlags commandFlags) String() string {
	bytes, err := json.MarshalIndent(cmdFlags, "", "  ")
	if err != nil {
		return err.Error()
	}
	return string(bytes)
}

func (cmdFlags *commandFlags) defineStringFlag(flagSet *flag.FlagSet, target *string, name string, defaultValue string, description string) {
	flagSet.StringVar(target, name, defaultValue, description)
}

func (cmdFlags *commandFlags) defineBoolFlag(flagSet *flag.FlagSet, target *bool, name string, defaultValue bool, description string) {
	flagSet.BoolVar(target, name, defaultValue, description)
}

func (cmdFlags *commandFlags) parseFlags(flagSet *flag.FlagSet, args []string) error {
	subt := 0

	if len(args) > 0 {
		subt = 1
	}
	err := flagSet.Parse(args[subt:])

	if err != nil {
		errMsg := err.Error()
		if strings.HasPrefix(errMsg, "flag provided but not defined: ") {
			flagName := errMsg[len("flag provided but not defined: "):]
			return fmt.Errorf("unkown flag '%s'", flagName)
		}
		return err
	}
	if cmdFlags.Environment == "" {
		cmdFlags.Environment = os.Getenv("DT_ENVIRONMENT")
	}
	if cmdFlags.APIToken == "" {
		cmdFlags.APIToken = os.Getenv("DT_APITOKEN")
	}
	return nil
}

func (cmdFlags *commandFlags) parse(args []string) error {
	flagSet := flag.NewFlagSet("dynatrace", flag.ContinueOnError)
	cmdFlags.defineStringFlag(flagSet, &cmdFlags.Environment, "environment", "", "Your Dynatrace Environment")
	cmdFlags.defineStringFlag(flagSet, &cmdFlags.APIToken, "api-token", "", "the API Token for accessing the Dynatrace REST API")
	cmdFlags.defineStringFlag(flagSet, &cmdFlags.Name, "name", "", "the name of the resulting web application")
	cmdFlags.defineBoolFlag(flagSet, &cmdFlags.Force, "f", false, "")
	return cmdFlags.parseFlags(flagSet, args)
}

func (cmdFlags *commandFlags) validate() bool {
	if cmdFlags.Environment == "" {
		log.Fail("No Dynatrace Environment ID specified. Use ", log.Cyan("-environment <environmentid>"), " or define the system environment variable ", log.Cyan("DT_ENVIRONMENT"), ".")
		return false
	}
	if cmdFlags.APIToken == "" {
		log.Fail("No API Token specified. Use ", log.Cyan("-api-token <api-token>"), " or define the system environment variable ", log.Cyan("DT_APITOKEN"), ".")
		return false
	}
	return true
}
