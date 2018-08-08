# AutoTag

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metadata** | [**ConfigurationMetadata**](ConfigurationMetadata.md) |  | [optional] 
**Id** | **string** | The ID of the rule-based tag. | [optional] 
**Name** | **string** | The name of the rule-based tag. This is the actual tag value which is associated with entities.If a &#39;valueFormat&#39; is defined inside a rule, the name of this tag will represent the tag key for this rule. | 
**Rules** | [**[]AutoTagRule**](AutoTagRule.md) | The list of rules when to associate this tag with an entity. Each rule is evaluated independently of all other rules. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


