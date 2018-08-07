package dtapi

import (
	dtapiconf "github.com/dtcookie/dtapi/libdtapiconf"
	dtapienv "github.com/dtcookie/dtapi/libdtapienv"
)

type envService struct {
	client *dtapienv.APIClient
}

type confService struct {
	client *dtapiconf.APIClient
}
