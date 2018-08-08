# \AnomalyDetectionVMwareApi

All URIs are relative to *https://siz65484.live.dynatrace.com/api/config/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetVMwareAnomalyDetectionConfig1**](AnomalyDetectionVMwareApi.md#GetVMwareAnomalyDetectionConfig1) | **Get** /anomalyDetection/vmware | Get configuration of anomaly detection for VMware.
[**UpdateVMwareAnomalyDetectionConfig1**](AnomalyDetectionVMwareApi.md#UpdateVMwareAnomalyDetectionConfig1) | **Put** /anomalyDetection/vmware | Update configuration of anomaly detection for VMware.
[**ValidateVMwareAnomalyDetectionConfig1**](AnomalyDetectionVMwareApi.md#ValidateVMwareAnomalyDetectionConfig1) | **Post** /anomalyDetection/vmware/validator | Validate configuration of anomaly detection for VMware (for PUT request).


# **GetVMwareAnomalyDetectionConfig1**
> VMwareAnomalyDetectionConfig GetVMwareAnomalyDetectionConfig1(ctx, )
Get configuration of anomaly detection for VMware.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**VMwareAnomalyDetectionConfig**](VMwareAnomalyDetectionConfig.md)

### Authorization

[ReadConfigToken](../README.md#ReadConfigToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateVMwareAnomalyDetectionConfig1**
> UpdateVMwareAnomalyDetectionConfig1(ctx, optional)
Update configuration of anomaly detection for VMware.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***UpdateVMwareAnomalyDetectionConfig1Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UpdateVMwareAnomalyDetectionConfig1Opts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **vMwareAnomalyDetectionConfig** | [**optional.Interface of VMwareAnomalyDetectionConfig**](VMwareAnomalyDetectionConfig.md)|  | 

### Return type

 (empty response body)

### Authorization

[WriteConfigToken](../README.md#WriteConfigToken)

### HTTP request headers

 - **Content-Type**: application/json; charset=utf-8
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ValidateVMwareAnomalyDetectionConfig1**
> ValidateVMwareAnomalyDetectionConfig1(ctx, optional)
Validate configuration of anomaly detection for VMware (for PUT request).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ValidateVMwareAnomalyDetectionConfig1Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ValidateVMwareAnomalyDetectionConfig1Opts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **vMwareAnomalyDetectionConfig** | [**optional.Interface of VMwareAnomalyDetectionConfig**](VMwareAnomalyDetectionConfig.md)|  | 

### Return type

 (empty response body)

### Authorization

[WriteConfigToken](../README.md#WriteConfigToken)

### HTTP request headers

 - **Content-Type**: application/json; charset=utf-8
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

