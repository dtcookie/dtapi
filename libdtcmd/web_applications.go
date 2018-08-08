package dtcmd

import (
	"errors"
	"fmt"
	"time"

	dtapi "github.com/dtcookie/dtapi/libdtapi"
	dtapiconf "github.com/dtcookie/dtapi/libdtapiconf"
	log "github.com/dtcookie/dtapi/libdtlog"
)

// webApplicationAPI is a wrapper around the more general
// WebApplicationConfigAPI offered via
// "github.com/dtcookie/dtapi/libdtapiconf"
//
// Functionality is restricted to use cases limited to the ones
// covered by current command line utilities.
// Utilization is meant to be package internal by the public
// functions of the Tenant object within this package.
type webApplicationAPI struct {
	tenant *Tenant
	base   *dtapi.WebApplicationConfigAPI
}

// exists checks whether a web application exists with the given
// unique identifier
//
// Parameters:
//	- ID	the unique identifier of a Web Application
// Returns:
//	- bool	'true' if a Web Application with that ID exists,
//			'false' otherwise
func (api *webApplicationAPI) exists(ID string) (bool, error) {
	var err error
	var configStubs []dtapiconf.WebApplicationConfigStub

	if configStubs, _, err = api.base.ListConfigurations(); err != nil {
		return false, err
	}
	for _, configStub := range configStubs {
		if configStub.Identifier == ID {
			return true, nil
		}
	}
	return false, nil
}

// namedDelete deletes the Web Application identified with the
// given 'name', including any application detection rules for that
// Web Application.
//
// Parameters:
//	- name	the name of the Web Application. Its unique identifier
//			will get resolved automatically
func (api *webApplicationAPI) namedDelete(name string) error {
	var err error
	var configStubs []dtapiconf.WebApplicationConfigStub

	if configStubs, _, err = api.base.ListConfigurations(); err != nil {
		return err
	}

	for _, configStub := range configStubs {
		fmt.Println(configStub.Name)
		if name == configStub.Name {
			return api.delete(configStub.Identifier, configStub.Name)
		}
	}
	log.Println("A web application with the name ", log.Cyan(name), " does not exist.")
	return nil
}

// delete deletes the Web Application identified with the
// given 'ID', including any application detection rules for that
// Web Application.
//
// Parameters:
//	- ID	the unique identifier of the Web Application
//	- name	the name of the Web Application (for logging purposes).
func (api *webApplicationAPI) delete(ID string, name string) error {
	var err error
	if err = api.tenant.applicationDetection.unbind(ID, name); err != nil {
		return err
	}
	log.Println("... deleting web application ", log.Cyan(name))
	if _, err = api.base.ForID(ID).Delete(); err != nil {
		return err
	}
	log.Println("... validating")
	numAttempts := 0
	for numAttempts = 0; numAttempts < 10; numAttempts++ {
		var exists bool
		if exists, err = api.exists(ID); err != nil {
			return err
		}
		if !exists {
			return nil
		}
		time.Sleep(1 * time.Second)
	}
	if numAttempts == 10 {
		log.Println(log.Yellow("WARNING"), " even after 10 seconds the deletion of the web application was not confirmed by the tenant.")
	}

	return nil
}

// get retrieves the detailed configuration for the Web Application identified
// by the given 'ID'.
//
// In case such a Web Application does not exists a 404 Not Found error can be
// expected. Therefore using this function requires a previous check for whether
// that application actually exists.
//
// Parameters:
// 	- ID	the unique identifier of the Web Application
// Returns:
//	- WebApplicationConfig	the configuration of the Web Application identified
//							by the given 'ID'
func (api *webApplicationAPI) get(ID string) (dtapiconf.WebApplicationConfig, error) {
	return api.base.ForID(ID).Get()
}

// getAll retrieves name and unique identifier for all currently existing
// Web Applications.
//
// Returns:
//	- []WebApplicationConfigStub	a list of name/ID pairs
func (api *webApplicationAPI) getAll() ([]dtapiconf.WebApplicationConfigStub, error) {
	configStubs, _, err := api.base.ListConfigurations()
	return configStubs, err
}

// list prints out a tabular list of currently configured Web Applications
func (api *webApplicationAPI) list() error {
	var err error
	var configStubs []dtapiconf.WebApplicationConfigStub

	if configStubs, _, err = api.base.ListConfigurations(); err != nil {
		return err
	}

	maxLen := 0
	for _, configStub := range configStubs {
		maxLen = max(maxLen, len(configStub.Name))
	}
	maxLen = maxLen + 2

	if len(configStubs) == 0 {
		return nil
	}

	log.Println(log.Gray("name"), spc(maxLen-len("name")), log.Gray("real user monitoring"), spc(maxLen-len("real user monitoring")))

	for _, configStub := range configStubs {
		log.Println(log.Cyan(configStub.Name), spc(maxLen-len(configStub.Name)), "on", spc(maxLen-len("on")))
	}
	return nil
}

// enableRUM checks whether the given Web Application has currently enabled RUM
// Monitoring. If not, it is getting enabled.
//
// Parameters:
//	- webAppConfig	the configuration of a previously fetched Web Application
func (api *webApplicationAPI) enableRUM(webAppConfig dtapiconf.WebApplicationConfig) error {
	if webAppConfig.RealUserMonitoringEnabled {
		return nil
	}
	log.Println("... enabling real user monitoring")
	webAppConfig.RealUserMonitoringEnabled = true
	webAppAPI := api.base.ForID(webAppConfig.Identifier)
	return webAppAPI.CreateOrUpdate(webAppConfig)
}

// createFromDefault creates a new Web Application with the given 'name'.
//
// As a template for the new Web Application the Default Web Application is
// being used.
//
// This function expects that checks whether a Web Application with the given
// name exists have already been done.
// The creation won't fail, but the eventual name of the Web Application may
// eventually differ from what has been requested.
//
// Parameters:
//	- name	the name of the new Web Application
// Returns:
//	- WebApplicationConfig	the configuration of the Web Application that has been
//							created.
func (api *webApplicationAPI) createFromDefault(name string) (dtapiconf.WebApplicationConfig, error) {
	var webAppConfig dtapiconf.WebApplicationConfig
	var err error
	if webAppConfig, err = api.base.Default.Get(); err != nil {
		return webAppConfig, err
	}
	webAppConfig.Identifier = ""
	webAppConfig.Name = name
	webAppConfig.RealUserMonitoringEnabled = true
	var configStub dtapiconf.WebApplicationConfigStub
	if configStub, _, err = api.base.Create(webAppConfig); err != nil {
		return webAppConfig, err
	}
	fmt.Println("... validating")
	fetchAttempts := 1
	for fetchAttempts = 1; fetchAttempts <= 10; fetchAttempts++ {
		webAppConfig, err = api.base.ForID(configStub.Identifier).Get()
		if err != nil {
			if err.Error() == "404 Not Found" {
				time.Sleep(2 * time.Second)
			} else {
				return webAppConfig, err
			}
		} else {
			break
		}
	}
	if fetchAttempts == 10 {
		return webAppConfig, errors.New("Failed to fetch application details after creation")
	}
	return webAppConfig, nil
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func spc(len int) string {
	s := ""
	i := 0
	for i = 0; i < len; i++ {
		s = s + " "
	}
	return s
}
