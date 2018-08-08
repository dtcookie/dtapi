# CapturedMethod

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Method** | [**MethodReference**](MethodReference.md) |  | 
**Capture** | **string** | What to capture from the method. | 
**ArgumentIndex** | **int32** | The index of the argument to capture. Set &#x60;0&#x60; to capture the return value, &#x60;1&#x60; or higher to capture a mehtod argument.    Required if the **capture** is set to &#x60;ARGUMENT&#x60;.   Not applicable in other cases. | [optional] 
**DeepObjectAccess** | **string** | The getter chain to apply to the captured object. It is required in one of the following cases:   The **capture** is set to &#x60;THIS&#x60;.    The **capture** is set to &#x60;ARGUMENT&#x60;, and the argument is not a primitive, a primitive wrapper class, a string, or an array.    Not applicable in other cases. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


