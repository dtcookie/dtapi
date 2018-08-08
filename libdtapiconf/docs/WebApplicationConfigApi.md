# \WebApplicationConfigApi

All URIs are relative to *https://siz65484.live.dynatrace.com/api/config/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateConfiguration1**](WebApplicationConfigApi.md#CreateConfiguration1) | **Post** /applications/web | Create a new web application configuration. The payload must not provide an id as that will be automatically assigned.
[**CreateOrUpdateConfiguration1**](WebApplicationConfigApi.md#CreateOrUpdateConfiguration1) | **Put** /applications/web/{id} | Store the configured web application configuration. If the payload has the id property set, it must match the id in the URL.
[**CreateOrUpdateDefaultConfiguration1**](WebApplicationConfigApi.md#CreateOrUpdateDefaultConfiguration1) | **Put** /applications/web/default | Update the default web application configuration.
[**DeleteConfiguration1**](WebApplicationConfigApi.md#DeleteConfiguration1) | **Delete** /applications/web/{id} | Delete the web application configuration with the given id.
[**GetConfiguration2**](WebApplicationConfigApi.md#GetConfiguration2) | **Get** /applications/web/{id} | Get the web application configuration with the given id.
[**GetDataPrivacySettings1**](WebApplicationConfigApi.md#GetDataPrivacySettings1) | **Get** /applications/web/{id}/dataPrivacy | Get the web application&#39;s data privacy settings.
[**GetDefaultApplication1**](WebApplicationConfigApi.md#GetDefaultApplication1) | **Get** /applications/web/default | Get the default web application configuration.
[**GetDefaultApplicationDataPrivacySettings1**](WebApplicationConfigApi.md#GetDefaultApplicationDataPrivacySettings1) | **Get** /applications/web/default/dataPrivacy | Get the default web application&#39;s data privacy settings.
[**ListConfigurations1**](WebApplicationConfigApi.md#ListConfigurations1) | **Get** /applications/web | List all application configurations.
[**ListDataPrivacySettings1**](WebApplicationConfigApi.md#ListDataPrivacySettings1) | **Get** /applications/web/dataPrivacy | List all applications&#39; data privacy settings.
[**UpdateDataPrivacySettings1**](WebApplicationConfigApi.md#UpdateDataPrivacySettings1) | **Put** /applications/web/{id}/dataPrivacy | Update the web application&#39;s data privacy settings.
[**UpdateDefaultApplicationDataPrivacySettings1**](WebApplicationConfigApi.md#UpdateDefaultApplicationDataPrivacySettings1) | **Put** /applications/web/default/dataPrivacy | Update the default web application&#39;s data privacy settings.
[**ValidateConfiguration2**](WebApplicationConfigApi.md#ValidateConfiguration2) | **Post** /applications/web/validator | Validates new web application configuration for the &#x60;POST /&#x60; request.
[**ValidateConfiguration3**](WebApplicationConfigApi.md#ValidateConfiguration3) | **Post** /applications/web/validator/{id} | Validates new web application configuration for the &#x60;PUT /{id}&#x60; request.
[**ValidateDataPrivacySettings1**](WebApplicationConfigApi.md#ValidateDataPrivacySettings1) | **Post** /applications/web/{id}/dataPrivacy/validator | Validates new data privacy settings for the &#x60;PUT /{id}/dataPrivacy&#x60; request.
[**ValidateDefaultApplicationDataPrivacySettings1**](WebApplicationConfigApi.md#ValidateDefaultApplicationDataPrivacySettings1) | **Post** /applications/web/default/dataPrivacy/validator | Validates new data privacy settings for the &#x60;PUT /default/dataPrivacy&#x60; request.
[**ValidateDefaultConfiguration1**](WebApplicationConfigApi.md#ValidateDefaultConfiguration1) | **Post** /applications/web/default/validator | Validates new default web application configuration for the &#x60;PUT /default&#x60; request.


# **CreateConfiguration1**
> WebApplicationConfigStub CreateConfiguration1(ctx, optional)
Create a new web application configuration. The payload must not provide an id as that will be automatically assigned.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CreateConfiguration1Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CreateConfiguration1Opts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webApplicationConfig** | [**optional.Interface of WebApplicationConfig**](WebApplicationConfig.md)|  | 

### Return type

[**WebApplicationConfigStub**](WebApplicationConfigStub.md)

### Authorization

[WriteConfigToken](../README.md#WriteConfigToken)

### HTTP request headers

 - **Content-Type**: application/json; charset=utf-8
 - **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateOrUpdateConfiguration1**
> CreateOrUpdateConfiguration1(ctx, id, optional)
Store the configured web application configuration. If the payload has the id property set, it must match the id in the URL.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***CreateOrUpdateConfiguration1Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CreateOrUpdateConfiguration1Opts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **webApplicationConfig** | [**optional.Interface of WebApplicationConfig**](WebApplicationConfig.md)|  | 

### Return type

 (empty response body)

### Authorization

[WriteConfigToken](../README.md#WriteConfigToken)

### HTTP request headers

 - **Content-Type**: application/json; charset=utf-8
 - **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateOrUpdateDefaultConfiguration1**
> CreateOrUpdateDefaultConfiguration1(ctx, optional)
Update the default web application configuration.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CreateOrUpdateDefaultConfiguration1Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CreateOrUpdateDefaultConfiguration1Opts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webApplicationConfig** | [**optional.Interface of WebApplicationConfig**](WebApplicationConfig.md)|  | 

### Return type

 (empty response body)

### Authorization

[WriteConfigToken](../README.md#WriteConfigToken)

### HTTP request headers

 - **Content-Type**: application/json; charset=utf-8
 - **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteConfiguration1**
> DeleteConfiguration1(ctx, id)
Delete the web application configuration with the given id.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[WriteConfigToken](../README.md#WriteConfigToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetConfiguration2**
> WebApplicationConfig GetConfiguration2(ctx, id)
Get the web application configuration with the given id.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 

### Return type

[**WebApplicationConfig**](WebApplicationConfig.md)

### Authorization

[ReadConfigToken](../README.md#ReadConfigToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDataPrivacySettings1**
> ApplicationDataPrivacy GetDataPrivacySettings1(ctx, id)
Get the web application's data privacy settings.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 

### Return type

[**ApplicationDataPrivacy**](ApplicationDataPrivacy.md)

### Authorization

[ReadConfigToken](../README.md#ReadConfigToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDefaultApplication1**
> WebApplicationConfig GetDefaultApplication1(ctx, )
Get the default web application configuration.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**WebApplicationConfig**](WebApplicationConfig.md)

### Authorization

[ReadConfigToken](../README.md#ReadConfigToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDefaultApplicationDataPrivacySettings1**
> ApplicationDataPrivacy GetDefaultApplicationDataPrivacySettings1(ctx, )
Get the default web application's data privacy settings.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ApplicationDataPrivacy**](ApplicationDataPrivacy.md)

### Authorization

[ReadConfigToken](../README.md#ReadConfigToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListConfigurations1**
> WebApplicationConfigStubListDto ListConfigurations1(ctx, )
List all application configurations.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**WebApplicationConfigStubListDto**](WebApplicationConfigStubListDto.md)

### Authorization

[ReadConfigToken](../README.md#ReadConfigToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListDataPrivacySettings1**
> ApplicationDataPrivacyListDto ListDataPrivacySettings1(ctx, )
List all applications' data privacy settings.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ApplicationDataPrivacyListDto**](ApplicationDataPrivacyListDto.md)

### Authorization

[ReadConfigToken](../README.md#ReadConfigToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateDataPrivacySettings1**
> UpdateDataPrivacySettings1(ctx, id, optional)
Update the web application's data privacy settings.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***UpdateDataPrivacySettings1Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UpdateDataPrivacySettings1Opts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **applicationDataPrivacy** | [**optional.Interface of ApplicationDataPrivacy**](ApplicationDataPrivacy.md)|  | 

### Return type

 (empty response body)

### Authorization

[WriteConfigToken](../README.md#WriteConfigToken)

### HTTP request headers

 - **Content-Type**: application/json; charset=utf-8
 - **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateDefaultApplicationDataPrivacySettings1**
> UpdateDefaultApplicationDataPrivacySettings1(ctx, optional)
Update the default web application's data privacy settings.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***UpdateDefaultApplicationDataPrivacySettings1Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UpdateDefaultApplicationDataPrivacySettings1Opts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **applicationDataPrivacy** | [**optional.Interface of ApplicationDataPrivacy**](ApplicationDataPrivacy.md)|  | 

### Return type

 (empty response body)

### Authorization

[WriteConfigToken](../README.md#WriteConfigToken)

### HTTP request headers

 - **Content-Type**: application/json; charset=utf-8
 - **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ValidateConfiguration2**
> ValidateConfiguration2(ctx, optional)
Validates new web application configuration for the `POST /` request.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ValidateConfiguration2Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ValidateConfiguration2Opts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webApplicationConfig** | [**optional.Interface of WebApplicationConfig**](WebApplicationConfig.md)|  | 

### Return type

 (empty response body)

### Authorization

[WriteConfigToken](../README.md#WriteConfigToken)

### HTTP request headers

 - **Content-Type**: application/json; charset=utf-8
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ValidateConfiguration3**
> ValidateConfiguration3(ctx, id, optional)
Validates new web application configuration for the `PUT /{id}` request.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***ValidateConfiguration3Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ValidateConfiguration3Opts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **webApplicationConfig** | [**optional.Interface of WebApplicationConfig**](WebApplicationConfig.md)|  | 

### Return type

 (empty response body)

### Authorization

[WriteConfigToken](../README.md#WriteConfigToken)

### HTTP request headers

 - **Content-Type**: application/json; charset=utf-8
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ValidateDataPrivacySettings1**
> ValidateDataPrivacySettings1(ctx, id, optional)
Validates new data privacy settings for the `PUT /{id}/dataPrivacy` request.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***ValidateDataPrivacySettings1Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ValidateDataPrivacySettings1Opts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **applicationDataPrivacy** | [**optional.Interface of ApplicationDataPrivacy**](ApplicationDataPrivacy.md)|  | 

### Return type

 (empty response body)

### Authorization

[WriteConfigToken](../README.md#WriteConfigToken)

### HTTP request headers

 - **Content-Type**: application/json; charset=utf-8
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ValidateDefaultApplicationDataPrivacySettings1**
> ValidateDefaultApplicationDataPrivacySettings1(ctx, optional)
Validates new data privacy settings for the `PUT /default/dataPrivacy` request.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ValidateDefaultApplicationDataPrivacySettings1Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ValidateDefaultApplicationDataPrivacySettings1Opts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **applicationDataPrivacy** | [**optional.Interface of ApplicationDataPrivacy**](ApplicationDataPrivacy.md)|  | 

### Return type

 (empty response body)

### Authorization

[WriteConfigToken](../README.md#WriteConfigToken)

### HTTP request headers

 - **Content-Type**: application/json; charset=utf-8
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ValidateDefaultConfiguration1**
> ValidateDefaultConfiguration1(ctx, optional)
Validates new default web application configuration for the `PUT /default` request.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ValidateDefaultConfiguration1Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ValidateDefaultConfiguration1Opts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webApplicationConfig** | [**optional.Interface of WebApplicationConfig**](WebApplicationConfig.md)|  | 

### Return type

 (empty response body)

### Authorization

[WriteConfigToken](../README.md#WriteConfigToken)

### HTTP request headers

 - **Content-Type**: application/json; charset=utf-8
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

