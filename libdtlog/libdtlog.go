package dtlog

import (
	"fmt"

	dtapiconf "github.com/dtcookie/dtapi/libdtapiconf"
	dtapienv "github.com/dtcookie/dtapi/libdtapienv"
	"github.com/fatih/color"
)

// ColorSection TODO: documentation
type colorized struct {
	fmt.Stringer
	color *color.Color
	text  string
}

func (c colorized) String() string {
	c.color.Print(c.text)
	return ""
}

// Cyan TODO: documentation
func Cyan(text string) colorized {
	return colorized{
		color: color.New(color.FgCyan, color.Bold),
		text:  text,
	}
}

// Red TODO: documentation
func Red(text string) colorized {
	return colorized{
		color: color.New(color.FgRed, color.Bold),
		text:  text,
	}
}

// Green TODO: documentation
func Green(text string) colorized {
	return colorized{
		color: color.New(color.FgGreen, color.Bold),
		text:  text,
	}
}

// Gray TODO: documentation
func Gray(text string) colorized {
	return colorized{
		color: color.New(color.FgHiBlack, color.Bold),
		text:  text,
	}
}

// Yellow TODO: documentation
func Yellow(text string) colorized {
	return colorized{
		color: color.New(color.FgYellow, color.Bold),
		text:  text,
	}
}

// Println TODO: documentation
func Println(v ...interface{}) {
	for _, value := range v {
		fmt.Print(value)
	}
	fmt.Println()
}

// FAILED TODO: documentation
func FAILED() {
	Println(Red("FAILED"))
}

// Fail TODO: documentation
func Fail(v ...interface{}) bool {
	FAILED()
	for _, value := range v {
		fmt.Print(value)
	}
	fmt.Println()
	return false
}

// FailErrorEnvelope TODO: documentation
func FailErrorEnvelope(errEnv *dtapiconf.ErrorEnvelope) bool {
	err := errEnv.Error
	Fail("Status Code: ", Red(fmt.Sprintf("%d", err.Code)))
	Println(err.Message)
	if len(err.ConstraintViolations) > 0 {
		for _, violation := range err.ConstraintViolations {
			Println(" - ", Red(violation.Path), ": ", violation.Message)
		}
	}
	return false
}

// FailConfGenericOpenAPIError TODO: documentation
func FailConfGenericOpenAPIError(err *dtapiconf.GenericOpenAPIError) bool {
	model := err.Model()
	switch errEnv := model.(type) {
	case dtapiconf.ErrorEnvelope:
		return FailErrorEnvelope(&errEnv)
	default:
		return Fail(err.Error())
	}
}

// FailEnvGenericOpenAPIError TODO: documentation
func FailEnvGenericOpenAPIError(err *dtapienv.GenericOpenAPIError) bool {
	return Fail(err.Error())
}

// FailError TODO: documentation
func FailError(err error) bool {
	switch goaerr := err.(type) {
	case dtapiconf.GenericOpenAPIError:
		return FailConfGenericOpenAPIError(&goaerr)
	case dtapienv.GenericOpenAPIError:
		return FailEnvGenericOpenAPIError(&goaerr)
	default:
		return Fail(err.Error())
	}
}
