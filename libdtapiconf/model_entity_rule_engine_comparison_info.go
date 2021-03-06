/*
 * Dynatrace Configuration API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package dtapiconf


// Type-specific comparison information for attributes. This is an abstract version of the comparison information, in its place use one of the actual implementations for the allowed types: * STRING -> EntityRuleEngineStringComparisonInfo * INTEGER -> EntityRuleEngineIntegerComparisonInfo * SERVICE_TYPE -> EntityRuleEngineServiceTypeComparisonInfo * PAAS_TYPE -> EntityRuleEnginePaasTypeComparisonInfo * CLOUD_TYPE -> EntityRuleEngineCloudTypeComparisonInfo * AZURE_SKU -> EntityRuleEngineAzureSkuComparisonInfo * AZURE_COMPUTE_MODE -> EntityRuleEngineAzureComputeModeComparisonInfo * ENTITY_ID -> EntityRuleEngineEntityIdComparisonInfo * SIMPLE_TECH -> EntityRuleEngineSimpleTechComparisonInfo * SERVICE_TOPOLOGY -> EntityRuleEngineServiceTopologyComparisonInfo * DATABASE_TOPOLOGY -> EntityRuleEngineDatabaseTopologyComparisonInfo * OS_TYPE -> EntityRuleEngineOsTypeComparisonInfo * HYPERVISOR_TYPE -> EntityRuleEngineHypervisorTypeComparisonInfo * IP_ADDRESS -> EntityRuleEngineIpAddressComparisonInfo * OS_ARCHITECTURE -> EntityRuleEngineOsArchitectureComparisonInfo * BITNESS -> EntityRuleEngineBitnessComparisonInfo * APPLICATION_TYPE -> EntityRuleEngineApplicationTypeComparisonInfo * MOBILE_PLATFORM -> EntityRuleEngineMobilePlatformComparisonInfo * CUSTOM_APPLICATION_TYPE -> EntityRuleEngineCustomApplicationTypeComparisonInfo * DCRUM_DECODER_TYPE -> EntityRuleEngineDcrumDecoderTypeComparisonInfo * SYNTHETIC_ENGINE_TYPE -> EntityRuleEngineSyntheticEngineTypeComparisonInfo * TAG -> EntityRuleEngineTagComparisonInfo
type EntityRuleEngineComparisonInfo struct {
	Operator Enum `json:"operator"`
	// Holds the actual value which should be used for comparison. Other parameters of this condition define how this value should be compared to actually stored values.
	Value map[string]interface{} `json:"value,omitempty"`
	// Negate the comparison.
	Negate bool `json:"negate"`
}
