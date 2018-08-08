# \AnomalyDetectionHostsApi

All URIs are relative to *https://siz65484.live.dynatrace.com/api/config/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetHostEventsConfig1**](AnomalyDetectionHostsApi.md#GetHostEventsConfig1) | **Get** /anomalyDetection/hosts | Get configuration of anomaly detection for hosts.
[**UpdateHostEventsConfig1**](AnomalyDetectionHostsApi.md#UpdateHostEventsConfig1) | **Put** /anomalyDetection/hosts | Update configuration of anomaly detection for hosts.
[**ValidateHostEventsConfig1**](AnomalyDetectionHostsApi.md#ValidateHostEventsConfig1) | **Post** /anomalyDetection/hosts/validator | Validate configuration of anomaly detection for hosts (for PUT request).


# **GetHostEventsConfig1**
> HostsAnomalyDetectionConfig GetHostEventsConfig1(ctx, )
Get configuration of anomaly detection for hosts.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**HostsAnomalyDetectionConfig**](HostsAnomalyDetectionConfig.md)

### Authorization

[ReadConfigToken](../README.md#ReadConfigToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateHostEventsConfig1**
> UpdateHostEventsConfig1(ctx, optional)
Update configuration of anomaly detection for hosts.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***UpdateHostEventsConfig1Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UpdateHostEventsConfig1Opts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **hostsAnomalyDetectionConfig** | [**optional.Interface of HostsAnomalyDetectionConfig**](HostsAnomalyDetectionConfig.md)|  | 

### Return type

 (empty response body)

### Authorization

[WriteConfigToken](../README.md#WriteConfigToken)

### HTTP request headers

 - **Content-Type**: application/json; charset=utf-8
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ValidateHostEventsConfig1**
> ValidateHostEventsConfig1(ctx, optional)
Validate configuration of anomaly detection for hosts (for PUT request).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ValidateHostEventsConfig1Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ValidateHostEventsConfig1Opts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **hostsAnomalyDetectionConfig** | [**optional.Interface of HostsAnomalyDetectionConfig**](HostsAnomalyDetectionConfig.md)|  | 

### Return type

 (empty response body)

### Authorization

[WriteConfigToken](../README.md#WriteConfigToken)

### HTTP request headers

 - **Content-Type**: application/json; charset=utf-8
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

