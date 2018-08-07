package dtcmd

import (
	"errors"
	"regexp"
	"strings"

	dtapi "github.com/dtcookie/dtapi/libdtapi"
	dtapiconf "github.com/dtcookie/dtapi/libdtapiconf"
	log "github.com/dtcookie/dtapi/libdtlog"
)

// ApplicationDetectionAPI TODO: documentation
type ApplicationDetectionAPI struct {
	tenant *Tenant
	base   *dtapi.ApplicationDetectionConfigAPI
}

func (api *ApplicationDetectionAPI) removeRulesOfWebApp(detectionRules []dtapiconf.ApplicationDetectionRule, webAppID string) []dtapiconf.ApplicationDetectionRule {
	result := make([]dtapiconf.ApplicationDetectionRule, 0)
	for _, detectionRule := range detectionRules {
		if detectionRule.ApplicationIdentifier != webAppID {
			result = append(result, detectionRule)
		}
	}
	return result
}

// UnBind TODO: documentation
func (api *ApplicationDetectionAPI) UnBind(ID string, name string) error {
	log.Println("... deleting application detection rules for web application ", log.Cyan(name))
	appDetectionConfig, err := api.base.Get()
	if err != nil {
		return api.tenant.setLastError(err)
	}
	appDetectionConfig.DetectionRules = api.removeRulesOfWebApp(appDetectionConfig.DetectionRules, ID)
	return api.tenant.setLastError(api.base.Update(appDetectionConfig))
}

// Get TODO: documentation
func (api *ApplicationDetectionAPI) Get() (dtapiconf.ApplicationDetectionConfig, error) {
	return api.base.Get()
}

// FindBindingDetectionRule TODO: documentation
func (api *ApplicationDetectionAPI) FindBindingDetectionRule(detectionRules []dtapiconf.ApplicationDetectionRule, domains []string, except string) (*dtapiconf.ApplicationDetectionRule, string) {
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

func removeFrom(elements []string, element string) []string {
	result := make([]string, 0)
	for _, elem := range elements {
		if elem != element {
			result = append(result, elem)
		}
	}
	return result
}

// BindDomains TODO: documentation
func (api *ApplicationDetectionAPI) BindDomains(ID string, appDetectionConfig dtapiconf.ApplicationDetectionConfig, domains []string) (bool, error) {
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
		return false, errors.New("unable to update application detections rules even after 10 attempts")
	}
	return true, nil
}

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

func patternMatchesDomains(pattern string, domains []string, applicationMatchType string) string {
	for _, domain := range domains {
		if patternMatchesDomain(pattern, domain, applicationMatchType) {
			return domain
		}
	}
	return ""
}

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
