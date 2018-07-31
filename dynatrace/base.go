package dynatrace

import (
	"net/http"

	"github.com/dtcookie/dtapi/config"
	"github.com/dtcookie/dtapi/environment"
)

type clients struct {
	env  *environment.APIClient
	conf *config.APIClient
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
