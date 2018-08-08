# \RequestAttributesApi

All URIs are relative to *https://siz65484.live.dynatrace.com/api/config/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateConfiguration2**](RequestAttributesApi.md#CreateConfiguration2) | **Post** /requestAttributes | Creates a new request attribute.
[**CreateOrUpdateConfiguration2**](RequestAttributesApi.md#CreateOrUpdateConfiguration2) | **Put** /requestAttributes/{id} | Updates an existing request attribute.
[**DeleteConfiguration2**](RequestAttributesApi.md#DeleteConfiguration2) | **Delete** /requestAttributes/{id} | Deletes the specified request attribute.
[**GetConfiguration3**](RequestAttributesApi.md#GetConfiguration3) | **Get** /requestAttributes/{id} | Shows the properties of the specified request attribute.
[**ListConfigurations2**](RequestAttributesApi.md#ListConfigurations2) | **Get** /requestAttributes | Lists all available request attributes.
[**ValidateConfiguration5**](RequestAttributesApi.md#ValidateConfiguration5) | **Post** /requestAttributes/validator | Validates new request attributes for the &#x60;POST /requestAttributes&#x60; request.
[**ValidateConfiguration6**](RequestAttributesApi.md#ValidateConfiguration6) | **Post** /requestAttributes/validator/{id} | Validate updates of existing request attribute for the &#x60;PUT /requestAttributes/{id}&#x60; request.


# **CreateConfiguration2**
> RequestAttributeStub CreateConfiguration2(ctx, optional)
Creates a new request attribute.

The body must not provide an ID as IDs are automatically assigned.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CreateConfiguration2Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CreateConfiguration2Opts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestAttribute** | [**optional.Interface of RequestAttribute**](RequestAttribute.md)|  | 

### Return type

[**RequestAttributeStub**](RequestAttributeStub.md)

### Authorization

[CaptureRequestDataToken](../README.md#CaptureRequestDataToken)

### HTTP request headers

 - **Content-Type**: application/json; charset=utf-8
 - **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateOrUpdateConfiguration2**
> CreateOrUpdateConfiguration2(ctx, id, optional)
Updates an existing request attribute.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | [**string**](.md)| The ID of the request attribute to be updated.   If you set the ID in the body as well, it must match this ID. | 
 **optional** | ***CreateOrUpdateConfiguration2Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CreateOrUpdateConfiguration2Opts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **requestAttribute** | [**optional.Interface of RequestAttribute**](RequestAttribute.md)|  | 

### Return type

 (empty response body)

### Authorization

[CaptureRequestDataToken](../README.md#CaptureRequestDataToken)

### HTTP request headers

 - **Content-Type**: application/json; charset=utf-8
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteConfiguration2**
> DeleteConfiguration2(ctx, id)
Deletes the specified request attribute.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | [**string**](.md)| The ID of the request attribute to be deleted. | 

### Return type

 (empty response body)

### Authorization

[CaptureRequestDataToken](../README.md#CaptureRequestDataToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetConfiguration3**
> RequestAttribute GetConfiguration3(ctx, id, optional)
Shows the properties of the specified request attribute.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | [**string**](.md)| The ID of the required request attribute. | 
 **optional** | ***GetConfiguration3Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetConfiguration3Opts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **includeProcessGroupReferences** | **optional.Bool**| Flag to include process group references to the response.    Process Group group references aren&#39;t compatible across environments. | [default to false]

### Return type

[**RequestAttribute**](RequestAttribute.md)

### Authorization

[ReadConfigToken](../README.md#ReadConfigToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListConfigurations2**
> RequestAttributeStubListDto ListConfigurations2(ctx, )
Lists all available request attributes.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**RequestAttributeStubListDto**](RequestAttributeStubListDto.md)

### Authorization

[ReadConfigToken](../README.md#ReadConfigToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ValidateConfiguration5**
> ValidateConfiguration5(ctx, optional)
Validates new request attributes for the `POST /requestAttributes` request.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ValidateConfiguration5Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ValidateConfiguration5Opts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestAttribute** | [**optional.Interface of RequestAttribute**](RequestAttribute.md)|  | 

### Return type

 (empty response body)

### Authorization

[CaptureRequestDataToken](../README.md#CaptureRequestDataToken)

### HTTP request headers

 - **Content-Type**: application/json; charset=utf-8
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ValidateConfiguration6**
> ValidateConfiguration6(ctx, id, optional)
Validate updates of existing request attribute for the `PUT /requestAttributes/{id}` request.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | [**string**](.md)| The ID of the request attribute to be validated. | 
 **optional** | ***ValidateConfiguration6Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ValidateConfiguration6Opts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **requestAttribute** | [**optional.Interface of RequestAttribute**](RequestAttribute.md)|  | 

### Return type

 (empty response body)

### Authorization

[CaptureRequestDataToken](../README.md#CaptureRequestDataToken)

### HTTP request headers

 - **Content-Type**: application/json; charset=utf-8
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

