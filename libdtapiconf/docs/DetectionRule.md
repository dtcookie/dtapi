# DetectionRule

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The ID of the detection rule. | [optional] 
**Enabled** | **bool** | Rule enabled/disabled. | 
**FileName** | **string** | The PHP file containing the class or methods to instrument.   Required for PHP custom service.    Not applicable to Java and .NET. | [optional] 
**ClassName** | **string** | The fully qualified class or interface to instrument.   Required for Java and .NET custom services.    Not applicable to PHP. | [optional] 
**Matcher** | **string** | Matcher applying to the class name. &#x60;STARTS_WITH&#x60; can only be used if there is at least one annotation defined. Default value is &#x60;EQUALS&#x60;. | [optional] 
**MethodRules** | [**[]MethodRule**](MethodRule.md) | List of methods to instrument. | 
**Annotations** | **[]string** | Additional annotations filter of the rule.   Only classes where all listed annotations are available in the class itself or any of its superclasses are instrumented.   Not applicable to PHP. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


