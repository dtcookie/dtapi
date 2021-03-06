/*
 * Dynatrace Environment API
 *
 * Documentation of the Dynatrace REST API. Refer to the [help page](https://www.dynatrace.com/support/help/shortlink/section-api) to read about use-cases and examples.
 *
 * API version: 1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package dtapienv


type Host struct {
	// Dynatrace entity ID of the required entity.   You can find them in the URL of the corresponding Dynatrace page, for example, `HOST-007`.
	EntityId string `json:"entityId,omitempty"`
	// The name of the Dynatrace entity, displayed in the UI.
	DisplayName string `json:"displayName,omitempty"`
	// Customized name of the entity
	CustomizedName string `json:"customizedName,omitempty"`
	// Discovered name of the entity
	DiscoveredName string `json:"discoveredName,omitempty"`
	// Timestamp in UTC milliseconds when the entity was detected for the first time.
	FirstSeenTimestamp int64 `json:"firstSeenTimestamp,omitempty"`
	// Timestamp in UTC milliseconds when the entity was detected for the last time.
	LastSeenTimestamp int64 `json:"lastSeenTimestamp,omitempty"`
	// The list of entity tags.
	Tags []TagInfo `json:"tags,omitempty"`
	FromRelationships map[string]interface{} `json:"fromRelationships,omitempty"`
	ToRelationships map[string]interface{} `json:"toRelationships,omitempty"`
	// 
	PaasType string `json:"paasType,omitempty"`
	// 
	OsVersion string `json:"osVersion,omitempty"`
	// 
	UserLevel string `json:"userLevel,omitempty"`
	// 
	OsType string `json:"osType,omitempty"`
	// 
	LogicalCpuCores int32 `json:"logicalCpuCores,omitempty"`
	// 
	OsArchitecture string `json:"osArchitecture,omitempty"`
	// 
	IsMonitoringCandidate bool `json:"isMonitoringCandidate,omitempty"`
	// 
	PaasMemoryLimit int64 `json:"paasMemoryLimit,omitempty"`
	// 
	CloudType string `json:"cloudType,omitempty"`
	// 
	HypervisorType string `json:"hypervisorType,omitempty"`
	// 
	MonitoringMode string `json:"monitoringMode,omitempty"`
	// 
	AzureComputeModeName string `json:"azureComputeModeName,omitempty"`
	// 
	Bitness string `json:"bitness,omitempty"`
	AgentVersion AgentVersion `json:"agentVersion,omitempty"`
	// 
	OpenstackAvZone string `json:"openstackAvZone,omitempty"`
	// 
	IpAddresses []string `json:"ipAddresses,omitempty"`
	// 
	CpuCores int32 `json:"cpuCores,omitempty"`
	// 
	AzureHostNames []string `json:"azureHostNames,omitempty"`
	// 
	AzureSiteNames []string `json:"azureSiteNames,omitempty"`
	// 
	ConsumedHostUnits string `json:"consumedHostUnits,omitempty"`
	// 
	AzureSku string `json:"azureSku,omitempty"`
}
