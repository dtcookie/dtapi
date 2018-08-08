# \AnomalyDetectionDiskEventsApi

All URIs are relative to *https://siz65484.live.dynatrace.com/api/config/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDiskEventConfig1**](AnomalyDetectionDiskEventsApi.md#CreateDiskEventConfig1) | **Post** /anomalyDetection/diskEvents | Create a new disk event rule.
[**DeleteDiskEventConfig1**](AnomalyDetectionDiskEventsApi.md#DeleteDiskEventConfig1) | **Delete** /anomalyDetection/diskEvents/{id} | Delete the disk alert rule with the given id.
[**GetDiskEventConfig1**](AnomalyDetectionDiskEventsApi.md#GetDiskEventConfig1) | **Get** /anomalyDetection/diskEvents/{id} | Get the disk event rule with the given id.
[**ListDiskEventConfigs1**](AnomalyDetectionDiskEventsApi.md#ListDiskEventConfigs1) | **Get** /anomalyDetection/diskEvents | Get the list of disk event rules.
[**UpdateOrCreateDiskEventConfig1**](AnomalyDetectionDiskEventsApi.md#UpdateOrCreateDiskEventConfig1) | **Put** /anomalyDetection/diskEvents/{id} | Update the disk event rule with the given id or create it if it doesn&#39;t exist yet.
[**ValidateDiskEventConfig1**](AnomalyDetectionDiskEventsApi.md#ValidateDiskEventConfig1) | **Post** /anomalyDetection/diskEvents/validator | Validate configuration of disk event rule.


# **CreateDiskEventConfig1**
> DiskEventAnomalyDetectionConfigStub CreateDiskEventConfig1(ctx, optional)
Create a new disk event rule.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CreateDiskEventConfig1Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CreateDiskEventConfig1Opts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **diskEventAnomalyDetectionConfig** | [**optional.Interface of DiskEventAnomalyDetectionConfig**](DiskEventAnomalyDetectionConfig.md)|  | 

### Return type

[**DiskEventAnomalyDetectionConfigStub**](DiskEventAnomalyDetectionConfigStub.md)

### Authorization

[WriteConfigToken](../README.md#WriteConfigToken)

### HTTP request headers

 - **Content-Type**: application/json; charset=utf-8
 - **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteDiskEventConfig1**
> DeleteDiskEventConfig1(ctx, id)
Delete the disk alert rule with the given id.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | [**string**](.md)|  | 

### Return type

 (empty response body)

### Authorization

[WriteConfigToken](../README.md#WriteConfigToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDiskEventConfig1**
> DiskEventAnomalyDetectionConfig GetDiskEventConfig1(ctx, id)
Get the disk event rule with the given id.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | [**string**](.md)|  | 

### Return type

[**DiskEventAnomalyDetectionConfig**](DiskEventAnomalyDetectionConfig.md)

### Authorization

[ReadConfigToken](../README.md#ReadConfigToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListDiskEventConfigs1**
> DiskEventAnomalyDetectionConfigList ListDiskEventConfigs1(ctx, )
Get the list of disk event rules.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**DiskEventAnomalyDetectionConfigList**](DiskEventAnomalyDetectionConfigList.md)

### Authorization

[ReadConfigToken](../README.md#ReadConfigToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateOrCreateDiskEventConfig1**
> UpdateOrCreateDiskEventConfig1(ctx, id, optional)
Update the disk event rule with the given id or create it if it doesn't exist yet.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | [**string**](.md)|  | 
 **optional** | ***UpdateOrCreateDiskEventConfig1Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UpdateOrCreateDiskEventConfig1Opts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **diskEventAnomalyDetectionConfig** | [**optional.Interface of DiskEventAnomalyDetectionConfig**](DiskEventAnomalyDetectionConfig.md)|  | 

### Return type

 (empty response body)

### Authorization

[WriteConfigToken](../README.md#WriteConfigToken)

### HTTP request headers

 - **Content-Type**: application/json; charset=utf-8
 - **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ValidateDiskEventConfig1**
> ValidateDiskEventConfig1(ctx, optional)
Validate configuration of disk event rule.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ValidateDiskEventConfig1Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ValidateDiskEventConfig1Opts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **diskEventAnomalyDetectionConfig** | [**optional.Interface of DiskEventAnomalyDetectionConfig**](DiskEventAnomalyDetectionConfig.md)|  | 

### Return type

 (empty response body)

### Authorization

[WriteConfigToken](../README.md#WriteConfigToken)

### HTTP request headers

 - **Content-Type**: application/json; charset=utf-8
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

