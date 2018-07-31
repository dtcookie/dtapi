package main

type argstype struct {
}

var argp = argstype{}

func (argv argstype) consume(args []string) (string, []string) {
	if len(args) == 0 {
		return "", args
	}
	return args[0], args[1:]
}
