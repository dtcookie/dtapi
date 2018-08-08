package dtcmd

import (
	"fmt"

	dtapi "github.com/dtcookie/dtapi/libdtapi"
	dtapiconf "github.com/dtcookie/dtapi/libdtapiconf"
	log "github.com/dtcookie/dtapi/libdtlog"
)

// Tenant is a wrapper around the Tenant object defined in
// the package "github.com/dtcookie/dtapi/libdtapi".
//
// Functionality is reduced down to the use cases of current
// command line utilities and cloud cli plugins.
type Tenant struct {
	// ActivityCallback allows the user to override
	// what is getting logged out during certain operations.
	ActivityCallback     ActivityCallback
	applicationDetection *applicationDetectionAPI
	webApplications      *webApplicationAPI
	base                 *dtapi.Tenant
}

// NewTenant creates a new Tenant object for accessing the
// Dynatrace REST API.
//
// Parameters:
// 	- credentials:	for authentication
// Returns:
//	- *Tenant	the fully initialized Tentant object
func NewTenant(credentials Credentials) *Tenant {
	base := dtapi.NewTenant(credentials.Environment, credentials.APIToken)
	tenant := Tenant{}

	tenant.ActivityCallback = &defaultActivityCallback{}
	tenant.base = base

	appDetectionAPI := applicationDetectionAPI{}
	appDetectionAPI.tenant = &tenant
	appDetectionAPI.base = base.APIs.ApplicationDetection
	tenant.applicationDetection = &appDetectionAPI

	webAppAPI := webApplicationAPI{}
	webAppAPI.tenant = &tenant
	webAppAPI.base = base.APIs.WebApplications
	tenant.webApplications = &webAppAPI

	return &tenant
}

// PrintClusterVersion prints out the cluster version of
// the Dynatrace Tenant
func (tenant *Tenant) PrintClusterVersion() error {
	version, err := tenant.base.APIs.Cluster.Version()
	fmt.Println("Dynatrace Cluster Version: " + version)
	return err
}

// UnBindWebApplication deletes application detection rules
// for the Web Application identified by the given name.
//
// Only application detection rules related to the given list
// of domains will be removed.
//
// Parameters:
//	- name		the name of the web application. Its identifier will
//				get resolved using that name
//	- domains	the domains the remaining application detection rules
//				may not match against
// Returns:
//	- bool	'true' if there was an update necessary, 'false' otherwise
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
	if appDetectionConfig, err = tenant.applicationDetection.get(); err != nil {
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
	if err = tenant.applicationDetection.base.Update(appDetectionConfig); err != nil {
		return false, err
	}
	return true, nil
}

// BindWebApplication creates application detection rules for the Web Application
// identified by the given 'webAppName', so that they cover the given 'domains'
//
// If the current application detection configuration already covers one of the given
// domains, no additional application detection rule will be created.
//
// Parameters:
//	- webAppName	the name of the Web Application. Its identifier will get resolved
//					internally based on that name
//	- domains		the domains to cover with the newly created application detection rules
// Returns:
//	- bool	'true' if there was an update necessary, 'false' otherwise
func (tenant *Tenant) BindWebApplication(webAppName string, domains []string) (bool, error) {
	var err error

	var webAppConfig dtapiconf.WebApplicationConfig
	foundConfig := false
	var configStubs []dtapiconf.WebApplicationConfigStub
	if configStubs, err = tenant.webApplications.getAll(); err != nil {
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
		if webAppConfig, err = tenant.webApplications.get(webAppID); err != nil {
			return false, err
		}
		foundConfig = true
	}

	var appDetectionConfig dtapiconf.ApplicationDetectionConfig
	if appDetectionConfig, err = tenant.applicationDetection.get(); err != nil {
		return false, err
	}

	if matchingDetectionRule, matchingDomain := tenant.applicationDetection.findBindingDetectionRule(appDetectionConfig.DetectionRules, domains, webAppID); matchingDetectionRule != nil {
		log.Fail("the route ", log.Cyan(matchingDomain), " already matches a detection rule for web application ", log.Cyan(webAppNamesByID[matchingDetectionRule.ApplicationIdentifier]), ". (", matchingDetectionRule.FilterConfig.ApplicationMatchType, " ", matchingDetectionRule.FilterConfig.Pattern, ").")
		log.Println("  remove that detection rule or run with ", log.Yellow("-f"), " to remove them automatically.")
		return false, nil
	}

	if !foundConfig {
		tenant.ActivityCallback.OnCreateWebApp(webAppName)
		if webAppConfig, err = tenant.webApplications.createFromDefault(webAppName); err != nil {
			return false, err
		}
	} else {
		if tenant.webApplications.enableRUM(webAppConfig) != nil {
			return false, err
		}
	}

	ok, err = tenant.applicationDetection.bindDomains(webAppConfig.Identifier, appDetectionConfig, domains)
	return true, err
}

// resolveWebAppID provides the identifier of a Web Application
// if the only thing known about it is the name.
//
// While also the name of a Web Application is being kept unique
// within a Tenant, it's not the identifier that is being used
// in other configuration settings.
//
// Parameters:
//	- name	the name of the Web Application
// Returns:
//	- *string	the identifier of the Web Application matching the
//				given name or 'nil' if none could be found.
func (tenant *Tenant) resolveWebAppID(name string) (*string, error) {
	var err error

	var configStubs []dtapiconf.WebApplicationConfigStub
	if configStubs, err = tenant.webApplications.getAll(); err != nil {
		return nil, err
	}
	for _, configStub := range configStubs {
		if configStub.Name == name {
			return &configStub.Identifier, nil
		}
	}
	return nil, nil
}

// defaultActivityCallback is the default implementation of
// the interface ActivityCallback
type defaultActivityCallback struct {
	ActivityCallback
}

// OnCreateWebApp simply logs out that a web application is about
// to be created via REST
func (callback *defaultActivityCallback) OnCreateWebApp(name string) {
	log.Println("... creating web application ", log.Cyan(name))
}
