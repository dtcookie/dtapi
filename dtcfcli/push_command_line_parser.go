package main

import "flag"

// PushArguments TODO: documentation
type PushArguments struct {
	Manifest string
}

// ParsePushFlags TODO: documentation
func ParsePushFlags(args []string) (*PushArguments, error) {
	var err error

	flagSet := flag.NewFlagSet("dynatrace", flag.ContinueOnError)
	manifestPath := ""
	flagSet.StringVar(&manifestPath, "f", "manifest.yml", "")
	if err = flagSet.Parse(args); err != nil {
		return nil, err
	}

	return &PushArguments{
		Manifest: manifestPath,
	}, nil
}
