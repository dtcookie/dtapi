# \CustomServicesApi

All URIs are relative to *https://siz65484.live.dynatrace.com/api/config/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Delete**](CustomServicesApi.md#Delete) | **Delete** /customServices/{technology}/{id} | Deletes the specified custom service
[**GetItem**](CustomServicesApi.md#GetItem) | **Get** /customServices/{technology}/{id} | Gets the definition of the specified custom service
[**GetList**](CustomServicesApi.md#GetList) | **Get** /customServices/{technology} | Lists all custom services of the specified technology along with their definitions
[**Post**](CustomServicesApi.md#Post) | **Post** /customServices/{technology} | Creates a custom service
[**PutItem**](CustomServicesApi.md#PutItem) | **Put** /customServices/{technology}/{id} | Updates the specified custom service.
[**PutList**](CustomServicesApi.md#PutList) | **Put** /customServices/{technology} | Updates all custom services of the specified technology
[**ValidateItem**](CustomServicesApi.md#ValidateItem) | **Post** /customServices/{technology}/validator | Validate the new custom service for the &#x60;POST /customServices/{technology}&#x60; request
[**ValidateList**](CustomServicesApi.md#ValidateList) | **Put** /customServices/{technology}/validator | Validates custom services for the &#x60;PUT /customServices/&#x60; request.


# **Delete**
> Delete(ctx, technology, id)
Deletes the specified custom service

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **technology** | **string**| Technology of the custom service to delete. | 
  **id** | [**string**](.md)| The ID of the custom service to delete. | 

### Return type

 (empty response body)

### Authorization

[WriteConfigToken](../README.md#WriteConfigToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetItem**
> CustomService GetItem(ctx, technology, id, optional)
Gets the definition of the specified custom service

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **technology** | **string**| Technology of the custom service you&#39;re inquiring. | 
  **id** | [**string**](.md)| The ID of the custom service you&#39;re inquiring. | 
 **optional** | ***GetItemOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetItemOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **includeProcessGroupReferences** | **optional.Bool**| Flag to include process group references to the response.    Process group references aren&#39;t compatible across environments.   &#x60;false&#x60; is used if the value is not set. | [default to false]

### Return type

[**CustomService**](CustomService.md)

### Authorization

[ReadConfigToken](../README.md#ReadConfigToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetList**
> CustomServiceList GetList(ctx, technology, optional)
Lists all custom services of the specified technology along with their definitions

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **technology** | **string**| Technology of the required custom services. | 
 **optional** | ***GetListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetListOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **includeProcessGroupReferences** | **optional.Bool**| Flag to include process group references to the response.    Process group references aren&#39;t compatible across environments.   &#x60;false&#x60; is used if the value is not set. | [default to false]

### Return type

[**CustomServiceList**](CustomServiceList.md)

### Authorization

[ReadConfigToken](../README.md#ReadConfigToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Post**
> Post(ctx, technology, optional)
Creates a custom service

In the body of the request, neither custom service nor its rules can have the ID. All IDs will be generated automatically by Dynatrace.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **technology** | **string**| Technology of the new custom service. | 
 **optional** | ***PostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PostOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **customService** | [**optional.Interface of CustomService**](CustomService.md)| JSON body of the request containing definition of the new custom service. 

You must not specify the IDs for the custom service or any of its rules. | 

### Return type

 (empty response body)

### Authorization

[WriteConfigToken](../README.md#WriteConfigToken)

### HTTP request headers

 - **Content-Type**: application/json; charset=utf-8
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutItem**
> PutItem(ctx, technology, id, optional)
Updates the specified custom service.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **technology** | **string**| Technology of the custom service to update. | 
  **id** | [**string**](.md)| The ID of the custom service to update.   The ID of the custom service in the body of the request must match this ID. | 
 **optional** | ***PutItemOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PutItemOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **customService** | [**optional.Interface of CustomService**](CustomService.md)| JSON body of the request containing updated definition of the custom service. | 

### Return type

 (empty response body)

### Authorization

[WriteConfigToken](../README.md#WriteConfigToken)

### HTTP request headers

 - **Content-Type**: application/json; charset=utf-8
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutList**
> PutList(ctx, technology, optional)
Updates all custom services of the specified technology

This request affects **all** the custom services of the specified technology. If a custom service is not presented in the body of the request, it will be deleted.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **technology** | **string**| Technology of custom services to update. | 
 **optional** | ***PutListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PutListOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **customServiceList** | [**optional.Interface of CustomServiceList**](CustomServiceList.md)| JSON body of the request containing updated definitions of custom services. | 

### Return type

 (empty response body)

### Authorization

[WriteConfigToken](../README.md#WriteConfigToken)

### HTTP request headers

 - **Content-Type**: application/json; charset=utf-8
 - **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ValidateItem**
> ValidateItem(ctx, technology, optional)
Validate the new custom service for the `POST /customServices/{technology}` request

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **technology** | **string**| Technology of the custom service to validate. | 
 **optional** | ***ValidateItemOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ValidateItemOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **customService** | [**optional.Interface of CustomService**](CustomService.md)| JSON body of the request containing definition of the custom service to validate. | 

### Return type

 (empty response body)

### Authorization

[WriteConfigToken](../README.md#WriteConfigToken)

### HTTP request headers

 - **Content-Type**: application/json; charset=utf-8
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ValidateList**
> ValidateList(ctx, technology, optional)
Validates custom services for the `PUT /customServices/` request.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **technology** | **string**| Technology of the custom services to validate. | 
 **optional** | ***ValidateListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ValidateListOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **customServiceList** | [**optional.Interface of CustomServiceList**](CustomServiceList.md)| JSON body of the request containing definitions of the custom services to validate. 

It must contain definitions of **all** custom services of the required technology. | 

### Return type

 (empty response body)

### Authorization

[WriteConfigToken](../README.md#WriteConfigToken)

### HTTP request headers

 - **Content-Type**: application/json; charset=utf-8
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

