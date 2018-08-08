# ManagementZoneRule

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Type of entities for which the management zone rule should be evaluated and associated to if it matches. | 
**Enabled** | **bool** | Management zone evaluation rule enabled/disabled | 
**PropagationTypes** | **[]string** | Types of propagation to be done on a matching management zone. Only some may be valid for the specified &#39;type&#39;. | [optional] 
**Conditions** | [**[]EntityRuleEngineCondition**](EntityRuleEngineCondition.md) | Actual conditions of the rule when the rule should match for an entity of the specified &#39;type&#39;.Only if all conditions match for an entity, the evaluation is positive and the management zone is associated with the entity. | 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


