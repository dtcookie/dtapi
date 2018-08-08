# \ManagementZonesApi

All URIs are relative to *https://siz65484.live.dynatrace.com/api/config/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateManagementZone1**](ManagementZonesApi.md#CreateManagementZone1) | **Post** /managementZones | Creates a new management zone.
[**CreateOrUpdateManagementZone1**](ManagementZonesApi.md#CreateOrUpdateManagementZone1) | **Put** /managementZones/{id} | Updates an existing management zone.
[**DeleteManagementZone1**](ManagementZonesApi.md#DeleteManagementZone1) | **Delete** /managementZones/{id} | Deletes the specified management zone.
[**GetAllManagementZoneConfigs1**](ManagementZonesApi.md#GetAllManagementZoneConfigs1) | **Get** /managementZones | Lists all configured management zones.
[**GetSingleManagementZoneConfig1**](ManagementZonesApi.md#GetSingleManagementZoneConfig1) | **Get** /managementZones/{id} | Shows the config properties of the specified management zone.
[**ValidateManagementZone1**](ManagementZonesApi.md#ValidateManagementZone1) | **Post** /managementZones/validator/{id} | Validate updates of existing management zone for the &#x60;PUT /managementZones/{id}&#x60; request.
[**ValidateManagementZone2**](ManagementZonesApi.md#ValidateManagementZone2) | **Post** /managementZones/validator | Validates new management zone for the &#x60;POST /managementZones&#x60; request.


# **CreateManagementZone1**
> ManagementZoneStub CreateManagementZone1(ctx, optional)
Creates a new management zone.

The body must not provide an ID as IDs are automatically assigned.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CreateManagementZone1Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CreateManagementZone1Opts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **managementZone** | [**optional.Interface of ManagementZone**](ManagementZone.md)|  | 

### Return type

[**ManagementZoneStub**](ManagementZoneStub.md)

### Authorization

[WriteConfigToken](../README.md#WriteConfigToken)

### HTTP request headers

 - **Content-Type**: application/json; charset=utf-8
 - **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateOrUpdateManagementZone1**
> CreateOrUpdateManagementZone1(ctx, id, optional)
Updates an existing management zone.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID of the management zone to be updated.   If you set the ID in the body as well, it must match this ID. | 
 **optional** | ***CreateOrUpdateManagementZone1Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CreateOrUpdateManagementZone1Opts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **managementZone** | [**optional.Interface of ManagementZone**](ManagementZone.md)|  | 

### Return type

 (empty response body)

### Authorization

[WriteConfigToken](../README.md#WriteConfigToken)

### HTTP request headers

 - **Content-Type**: application/json; charset=utf-8
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteManagementZone1**
> DeleteManagementZone1(ctx, id)
Deletes the specified management zone.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID of the management zone to be deleted. | 

### Return type

 (empty response body)

### Authorization

[WriteConfigToken](../README.md#WriteConfigToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAllManagementZoneConfigs1**
> ManagementZoneStubList GetAllManagementZoneConfigs1(ctx, )
Lists all configured management zones.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ManagementZoneStubList**](ManagementZoneStubList.md)

### Authorization

[ReadConfigToken](../README.md#ReadConfigToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSingleManagementZoneConfig1**
> ManagementZone GetSingleManagementZoneConfig1(ctx, id, optional)
Shows the config properties of the specified management zone.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID of the required management zone. | 
 **optional** | ***GetSingleManagementZoneConfig1Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetSingleManagementZoneConfig1Opts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **includeProcessGroupReferences** | **optional.Bool**| Flag to include process group references to the response.    Process group references aren&#39;t compatible across environments. When this flag is set to false, conditions with process group references will be removed. If that leads to a rule having no conditions, the entire rule will be removed. | [default to false]

### Return type

[**ManagementZone**](ManagementZone.md)

### Authorization

[ReadConfigToken](../README.md#ReadConfigToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ValidateManagementZone1**
> ValidateManagementZone1(ctx, id, optional)
Validate updates of existing management zone for the `PUT /managementZones/{id}` request.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID of the management zone to be validated. | 
 **optional** | ***ValidateManagementZone1Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ValidateManagementZone1Opts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **managementZone** | [**optional.Interface of ManagementZone**](ManagementZone.md)|  | 

### Return type

 (empty response body)

### Authorization

[WriteConfigToken](../README.md#WriteConfigToken)

### HTTP request headers

 - **Content-Type**: application/json; charset=utf-8
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ValidateManagementZone2**
> ValidateManagementZone2(ctx, optional)
Validates new management zone for the `POST /managementZones` request.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ValidateManagementZone2Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ValidateManagementZone2Opts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **managementZone** | [**optional.Interface of ManagementZone**](ManagementZone.md)|  | 

### Return type

 (empty response body)

### Authorization

[WriteConfigToken](../README.md#WriteConfigToken)

### HTTP request headers

 - **Content-Type**: application/json; charset=utf-8
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

