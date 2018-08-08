package dtapi

import (
	dtapiconf "github.com/dtcookie/dtapi/libdtapiconf"
	dtapienv "github.com/dtcookie/dtapi/libdtapienv"
)

// envService is the base type for any service
// accessing the Environment REST API
type envService struct {
	client *dtapienv.APIClient
}

// confService is the base type for any service
// accessing the Config REST API
type confService struct {
	client *dtapiconf.APIClient
}
