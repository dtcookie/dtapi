package dtcmd

import (
	"fmt"

	dtapi "github.com/dtcookie/dtapi/libdtapi"
	dtapiconf "github.com/dtcookie/dtapi/libdtapiconf"
	log "github.com/dtcookie/dtapi/libdtlog"
)

// Tenant TODO: documentation
type Tenant struct {
	ApplicationDetection *ApplicationDetectionAPI
	WebApplications      *WebApplicationAPI
	base                 *dtapi.Tenant
	lastError            error
	ActivityCallback     ActivityCallback
}

// NewTenant TODO: documentation
func NewTenant(environment string, apiToken string) *Tenant {
	base := dtapi.NewTenant(environment, apiToken)
	tenant := Tenant{}

	tenant.ActivityCallback = &DefaultActivityCallback{}
	tenant.base = base
	tenant.lastError = nil

	appDetectionAPI := ApplicationDetectionAPI{}
	appDetectionAPI.tenant = &tenant
	appDetectionAPI.base = base.APIs.ApplicationDetection
	tenant.ApplicationDetection = &appDetectionAPI

	webAppAPI := WebApplicationAPI{}
	webAppAPI.tenant = &tenant
	webAppAPI.base = base.APIs.WebApplications
	tenant.WebApplications = &webAppAPI

	return &tenant
}

func (tenant *Tenant) setLastError(err error) error {
	tenant.lastError = err
	return err
}

// LogLastErr TODO: documentation
func (tenant *Tenant) LogLastErr() bool {
	if tenant.lastError != nil {
		return log.FailError(tenant.lastError)
	}
	return true
}

// CheckClusterVersion TODO: documentation
func (tenant *Tenant) CheckClusterVersion() error {
	logVersion := false
	version, err := tenant.base.APIs.Cluster.Version()
	if err != nil {
		return tenant.setLastError(err)
	}
	if logVersion {
		fmt.Println("Dynatrace Cluster Version: " + version)
	}
	return tenant.setLastError(nil)
}

func (tenant *Tenant) resolveWebAppID(name string) (*string, error) {
	var err error

	var configStubs []dtapiconf.WebApplicationConfigStub
	if configStubs, err = tenant.WebApplications.GetAll(); err != nil {
		return nil, err
	}
	for _, configStub := range configStubs {
		if configStub.Name == name {
			return &configStub.Identifier, nil
		}
	}
	return nil, nil
}

// UnBindWebApplication TODO: documentation
func (tenant *Tenant) UnBindWebApplication(name string, domains []string) (bool, error) {
	var err error

	var webAppID *string
	if webAppID, err = tenant.resolveWebAppID(name); err != nil {
		return false, err
	}

	if webAppID == nil {
		log.Println("A web application with the name ", log.Cyan(name), " does not exist in dynatrace.")
		log.Println("You may want to specify the option ", log.Yellow("-name <webapp>"), " in order to specfify a different name.")
		return false, nil
	}

	var appDetectionConfig dtapiconf.ApplicationDetectionConfig
	if appDetectionConfig, err = tenant.ApplicationDetection.Get(); err != nil {
		return false, err
	}
	remainingDetectionrules := make([]dtapiconf.ApplicationDetectionRule, 0)
	for _, detectionRule := range appDetectionConfig.DetectionRules {
		if detectionRule.ApplicationIdentifier == *webAppID {
			matchingDomain := ruleMatchesDomain(detectionRule, domains)
			if matchingDomain != "" {
				log.Println("... removing detection rule matching the route ", log.Cyan(matchingDomain))
				continue
			}
		}
		remainingDetectionrules = append(remainingDetectionrules, detectionRule)
	}

	appDetectionConfig.DetectionRules = remainingDetectionrules
	if err = tenant.ApplicationDetection.base.Update(appDetectionConfig); err != nil {
		return false, err
	}
	return true, nil
}

// BindWebApplication BindWebApplication
func (tenant *Tenant) BindWebApplication(webAppName string, domains []string) (bool, error) {
	var err error

	var webAppConfig dtapiconf.WebApplicationConfig
	foundConfig := false
	var configStubs []dtapiconf.WebApplicationConfigStub
	if configStubs, err = tenant.WebApplications.GetAll(); err != nil {
		return false, err
	}
	webAppNamesByID := map[string]string{}
	webAppIdsByName := map[string]string{}
	for _, configStub := range configStubs {
		webAppNamesByID[configStub.Identifier] = configStub.Name
		webAppIdsByName[configStub.Name] = configStub.Identifier
	}
	var webAppID string
	var ok bool
	if webAppID, ok = webAppIdsByName[webAppName]; ok {
		if webAppConfig, err = tenant.WebApplications.Get(webAppID); err != nil {
			return false, err
		}
		foundConfig = true
	}

	var appDetectionConfig dtapiconf.ApplicationDetectionConfig
	if appDetectionConfig, err = tenant.ApplicationDetection.Get(); err != nil {
		return false, err
	}

	if matchingDetectionRule, matchingDomain := tenant.ApplicationDetection.FindBindingDetectionRule(appDetectionConfig.DetectionRules, domains, webAppID); matchingDetectionRule != nil {
		log.Fail("the route ", log.Cyan(matchingDomain), " already matches a detection rule for web application ", log.Cyan(webAppNamesByID[matchingDetectionRule.ApplicationIdentifier]), ". (", matchingDetectionRule.FilterConfig.ApplicationMatchType, " ", matchingDetectionRule.FilterConfig.Pattern, ").")
		log.Println("  remove that detection rule or run with ", log.Yellow("-f"), " to remove them automatically.")
		return false, nil
	}

	if !foundConfig {
		tenant.ActivityCallback.OnCreateWebApp(webAppName)
		if webAppConfig, err = tenant.WebApplications.CreateFromDefault(webAppName); err != nil {
			return false, err
		}
	} else {
		if tenant.WebApplications.EnableRUM(webAppConfig) != nil {
			return false, err
		}
	}

	ok, err = tenant.ApplicationDetection.BindDomains(webAppConfig.Identifier, appDetectionConfig, domains)
	return true, err
}

// ActivityCallback TODO: documentation
type ActivityCallback interface {
	OnCreateWebApp(name string)
}

type DefaultActivityCallback struct {
	ActivityCallback
}

func (callback *DefaultActivityCallback) OnCreateWebApp(name string) {
	log.Println("... creating web application ", log.Cyan(name))
}
