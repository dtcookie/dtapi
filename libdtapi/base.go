package dtapi

import (
	"net/http"

	dtapiconf "github.com/dtcookie/dtapi/libdtapiconf"
	dtapienv "github.com/dtcookie/dtapi/libdtapienv"
)

type clients struct {
	env  *dtapienv.APIClient
	conf *dtapiconf.APIClient
}

type baseTenant struct {
	clients clients
}

func check204(httpResponse *http.Response, err error) (bool, error) {
	if httpResponse.StatusCode == 204 {
		return true, nil
	} else if httpResponse.StatusCode == 400 {
		return false, nil
	} else if err != nil {
		return false, err
	}
	return false, nil
}
