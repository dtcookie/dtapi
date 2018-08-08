package dtcmd

import (
	"errors"
	"regexp"
	"strings"

	dtapi "github.com/dtcookie/dtapi/libdtapi"
	dtapiconf "github.com/dtcookie/dtapi/libdtapiconf"
	log "github.com/dtcookie/dtapi/libdtlog"
)

// applicationDetectionAPI is a wrapper around the general
// Application Detection API offered by the package
// "github.com/dtcookie/dtapi/libdtapiconf".
//
// It's not available outside the this package, but expected
// to be utilized by the public Tenant functions within this
// package.
//
// Functionality is reduced to the requirements the current
// command line utilities and several cli plugins require.
type applicationDetectionAPI struct {
	tenant *Tenant
	base   *dtapi.ApplicationDetectionConfigAPI
}

// unbind deletes all the application detection rules for the
// specified Dynatrace Web Application
//
// Parameters:
// 	- ID	the identifier of the Web Application.
//	- name	the name of the Web Application (for logging purposes).
func (api *applicationDetectionAPI) unbind(ID string, name string) error {
	var err error
	var appDetectionConfig dtapiconf.ApplicationDetectionConfig
	log.Println("... deleting application detection rules for web application ", log.Cyan(name))
	if appDetectionConfig, err = api.base.Get(); err != nil {
		return err
	}
	appDetectionConfig.DetectionRules = api.removeRulesOfWebApp(appDetectionConfig.DetectionRules, ID)
	return api.base.Update(appDetectionConfig)
}

// Get retrieves the current application detection settings
func (api *applicationDetectionAPI) get() (dtapiconf.ApplicationDetectionConfig, error) {
	return api.base.Get()
}

// bindDomains creates the necessary application detection rules for a Web Applications
// so that they cover the specified domain names.
//
// Any current application detection rules will remain configured.
//
// Parameters:
// 	- ID					the unique identifier of the Web Application
//  - appDetectionConfig	the current application detection config of the web app
// Returns:
//	- bool	'true' if the application detection config was updated, 'false' if it wasn't
//			necessary to do so or if an error occured
func (api *applicationDetectionAPI) bindDomains(ID string, appDetectionConfig dtapiconf.ApplicationDetectionConfig, domains []string) (bool, error) {
	for _, detectionRule := range appDetectionConfig.DetectionRules {
		matchingDomain := ruleMatchesDomain(detectionRule, domains)
		if matchingDomain != "" {
			domains = removeFrom(domains, matchingDomain)
		}
	}

	if len(domains) == 0 {
		return false, nil
	}
	for _, route := range domains {
		detectionRule := dtapiconf.ApplicationDetectionRule{
			ApplicationIdentifier: ID,
			FilterConfig: dtapiconf.ApplicationFilter{
				Pattern:                route,
				ApplicationMatchType:   "EQUALS",
				ApplicationMatchTarget: "DOMAIN",
			},
		}
		appDetectionConfig.DetectionRules = append(appDetectionConfig.DetectionRules, detectionRule)
		log.Println("... creating application detection rule for route ", log.Cyan(route))
	}
	updateAttempts := 1
	for updateAttempts = 1; updateAttempts <= 10; updateAttempts++ {
		if err := api.base.Update(appDetectionConfig); err == nil {
			return true, nil
		}
	}
	if updateAttempts == 10 {
		return true, errors.New("unable to update application detections rules even after 10 attempts")
	}
	return true, nil
}

// findBindingDetectionRule queries within the given application detection rules for
// entries that are matching one of the given domains, but doesn't match the given
// web application identifier (parameter except).
//
// Purpose of this function is to figure out, whether for the given domains there are
// already existing application detection rules, but for a different web application
// than the one specified via parameter 'except'.
//
// Parameters:
//	- detectionRules	the application detection rules to query for matches
//	- domains			the domains the detection rules may possibly match
//	- except			the unique identifier of a web application the detection
//						rule must not match or "" if looking for matches for any
//						web applications
// Returns:
//	- *ApplicationDetectionRule	the first application detection rule found that matches
//								either one of the given domains
//	- string					the domain the application detection rule matched
func (api *applicationDetectionAPI) findBindingDetectionRule(detectionRules []dtapiconf.ApplicationDetectionRule, domains []string, except string) (*dtapiconf.ApplicationDetectionRule, string) {
	if len(detectionRules) == 0 {
		return nil, ""
	}
	if len(domains) == 0 {
		return nil, ""
	}
	for _, detectionRule := range detectionRules {
		matchingDomain := ruleMatchesDomain(detectionRule, domains)
		if matchingDomain != "" {
			if detectionRule.ApplicationIdentifier != except {
				return &detectionRule, matchingDomain
			}
		}
	}
	return nil, ""
}

// removeRulesOfWebApp removes any application detection rules of the given rules
// that are matching the Web Application identified by the given 'webAppID'.
//
// This function does NOT update anything via REST calls, but simply removes entries
// from the given array.
//
// Parameters:
//	- detectionRules:	the application detection rules to remove entries from
//	- webAppID:			the unique identifier of the Web Application a detection
//						rule needs to match in order to get removed from the list
// Returns:
//	- []ApplicationDetectionRule	a list of application detection rules that doesn't
//									contain any rules matching the given 'webAppID'
func (api *applicationDetectionAPI) removeRulesOfWebApp(detectionRules []dtapiconf.ApplicationDetectionRule, webAppID string) []dtapiconf.ApplicationDetectionRule {
	result := make([]dtapiconf.ApplicationDetectionRule, 0)
	for _, detectionRule := range detectionRules {
		if detectionRule.ApplicationIdentifier != webAppID {
			result = append(result, detectionRule)
		}
	}
	return result
}

// ruleMatchesDomain queries within the given domain names for the first one
// that is covered by the given application detection rule.
//
// Parameters:
//	- detectionRule	the application detection rule that is potentially
//					covering one of the given domains.
//	- domains		the domains the application detection rule is potentially
//					covering.
// Returns:
//	- string	the first domain found to be covered by the given application
//				detection rule or "" if none is covered
func ruleMatchesDomain(detectionRule dtapiconf.ApplicationDetectionRule, domains []string) string {
	filterConfig := detectionRule.FilterConfig
	applicationMatchTarget := filterConfig.ApplicationMatchTarget // URL, DOMAIN
	if applicationMatchTarget != "DOMAIN" {
		return ""
	}
	pattern := filterConfig.Pattern
	applicationMatchType := filterConfig.ApplicationMatchType // CONTAINS, BEGINS_WITH, ENDS_WITH, MATCHES, EQUALS

	return patternMatchesDomains(pattern, domains, applicationMatchType)
}

// patternMatchesDomains queries within the given domains for the first one
// matching the given pattern based on the given application match type.
//
// Parameters:
//	- pattern				the pattern the resulting domain must match
//	- domains				the domains to check whether they match with the pattern
//	- applicationMatchType	either CONTAINS, BEGINS_WITH, ENDS_WITH, MATCHES or EQUALS
// Returns:
//	- string	the first domain in the given list of domains that matches the given pattern
//				based on the given applicationMatchType
func patternMatchesDomains(pattern string, domains []string, applicationMatchType string) string {
	for _, domain := range domains {
		if patternMatchesDomain(pattern, domain, applicationMatchType) {
			return domain
		}
	}
	return ""
}

// patternMatchesDomain checks whether the given domain matches the given pattern
// based on the given application match type.
//
// Parameters:
//	- pattern				the pattern the resulting domain must match
//	- domain				the domain to check whether it matches with the pattern
//	- applicationMatchType	either CONTAINS, BEGINS_WITH, ENDS_WITH, MATCHES or EQUALS
// Returns:
//	- bool	'true' if the pattern matches the domain, 'false' otherwise
func patternMatchesDomain(pattern string, domain string, applicationMatchType string) bool {
	switch applicationMatchType {
	case "CONTAINS":
		return strings.HasPrefix(domain, pattern) || strings.HasSuffix(domain, pattern)
	case "BEGINS_WITH":
		return strings.HasPrefix(domain, pattern)
	case "ENDS_WITH":
		return strings.HasSuffix(domain, pattern)
	case "MATCHES":
		matched, _ := regexp.MatchString(pattern, domain)
		return matched
	case "EQUALS":
		return pattern == domain
	default:
		return false

	}
}

// removeFrom is a helper method. It simply removes elements
// from a list of strings.
func removeFrom(elements []string, element string) []string {
	result := make([]string, 0)
	for _, elem := range elements {
		if elem != element {
			result = append(result, elem)
		}
	}
	return result
}
