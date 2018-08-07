package dtcmd

import (
	"errors"
	"fmt"
	"time"

	dtapi "github.com/dtcookie/dtapi/libdtapi"
	dtapiconf "github.com/dtcookie/dtapi/libdtapiconf"
	log "github.com/dtcookie/dtapi/libdtlog"
)

// WebApplicationAPI TODO: documentation
type WebApplicationAPI struct {
	tenant *Tenant
	base   *dtapi.WebApplicationConfigAPI
}

// Exists TODO: documentation
func (api *WebApplicationAPI) Exists(ID string) (bool, error) {
	configStubs, _, err := api.base.ListConfigurations()
	if err != nil {
		return false, api.tenant.setLastError(err)
	}
	for _, configStub := range configStubs {
		if configStub.Identifier == ID {
			return true, nil
		}
	}
	return false, nil
}

// NamedDelete TODO: documentation
func (api *WebApplicationAPI) NamedDelete(name string) error {
	configStubs, _, err := api.base.ListConfigurations()
	if err != nil {
		return api.tenant.setLastError(err)
	}

	for _, configStub := range configStubs {
		fmt.Println(configStub.Name)
		if name == configStub.Name {
			return api.tenant.setLastError(api.Delete(configStub.Identifier, configStub.Name))
		}
	}
	log.Println("A web application with the name ", log.Cyan(name), " does not exist.")
	return api.tenant.setLastError(nil)
}

// Delete TODO: documentation
func (api *WebApplicationAPI) Delete(ID string, name string) error {
	if err := api.tenant.ApplicationDetection.UnBind(ID, name); err != nil {
		return api.tenant.setLastError(err)
	}
	log.Println("... deleting web application ", log.Cyan(name))
	if _, err := api.base.ForID(ID).Delete(); err != nil {
		return api.tenant.setLastError(err)
	}
	log.Println("... validating")
	numAttempts := 0
	for numAttempts = 0; numAttempts < 10; numAttempts++ {
		exists, err := api.Exists(ID)
		if err != nil {
			return api.tenant.setLastError(err)
		}
		if !exists {
			return api.tenant.setLastError(nil)
		}
		time.Sleep(1 * time.Second)
	}
	if numAttempts == 10 {
		log.Println(log.Yellow("WARNING"), " even after 10 seconds the deletion of the web application was not confirmed by the tenant.")
	}

	return api.tenant.setLastError(nil)
}

// Get TODO: documentation
func (api *WebApplicationAPI) Get(ID string) (dtapiconf.WebApplicationConfig, error) {
	return api.base.ForID(ID).Get()
}

// GetAll TODO: documentation
func (api *WebApplicationAPI) GetAll() ([]dtapiconf.WebApplicationConfigStub, error) {
	configStubs, _, err := api.base.ListConfigurations()
	return configStubs, api.tenant.setLastError(err)
}

// List TODO: documentation
func (api *WebApplicationAPI) List() error {
	configStubs, _, err := api.base.ListConfigurations()
	if err != nil {
		return api.tenant.setLastError(err)
	}

	maxLen := 0
	for _, configStub := range configStubs {
		maxLen = max(maxLen, len(configStub.Name))
	}
	maxLen = maxLen + 2

	if len(configStubs) == 0 {
		return api.tenant.setLastError(nil)
	}

	log.Println(log.Gray("name"), spc(maxLen-len("name")), log.Gray("real user monitoring"), spc(maxLen-len("real user monitoring")))

	for _, configStub := range configStubs {
		log.Println(log.Cyan(configStub.Name), spc(maxLen-len(configStub.Name)), "on", spc(maxLen-len("on")))
	}
	return api.tenant.setLastError(nil)
}

// EnableRUM TODO: documentation
func (api *WebApplicationAPI) EnableRUM(webAppConfig dtapiconf.WebApplicationConfig) error {
	if webAppConfig.RealUserMonitoringEnabled {
		return nil
	}
	log.Println("... enabling real user monitoring")
	webAppConfig.RealUserMonitoringEnabled = true
	webAppAPI := api.base.ForID(webAppConfig.Identifier)
	return webAppAPI.CreateOrUpdate(webAppConfig)
}

// CreateFromDefault TODO: documentation
func (api *WebApplicationAPI) CreateFromDefault(name string) (dtapiconf.WebApplicationConfig, error) {
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
