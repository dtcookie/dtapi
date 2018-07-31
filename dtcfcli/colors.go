package main

import (
	"fmt"

	"github.com/fatih/color"
)

type colorstype struct {
}

var colors = colorstype{}

func (c colorstype) cyan(text string) {
	color.New(color.FgCyan, color.Bold).Printf(text)
}

func (c colorstype) red(text string) {
	color.New(color.FgRed, color.Bold).Printf(text)
}

func (c colorstype) green(text string) {
	color.New(color.FgGreen, color.Bold).Printf(text)
}

// PrintfLn PrintfLn
func PrintfLn(format string, a ...interface{}) (n int, err error) {
	return fmt.Println(fmt.Sprintf(format, a))
}
