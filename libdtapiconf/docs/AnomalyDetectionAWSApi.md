# \AnomalyDetectionAWSApi

All URIs are relative to *https://siz65484.live.dynatrace.com/api/config/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAwsAnomalyDetectionConfig1**](AnomalyDetectionAWSApi.md#GetAwsAnomalyDetectionConfig1) | **Get** /anomalyDetection/aws | Get configuration of anomaly detection for AWS.
[**UpdateAwsAnomalyDetectionConfig1**](AnomalyDetectionAWSApi.md#UpdateAwsAnomalyDetectionConfig1) | **Put** /anomalyDetection/aws | Update configuration of anomaly detection for AWS.
[**ValidateAwsAnomalyDetectionConfig1**](AnomalyDetectionAWSApi.md#ValidateAwsAnomalyDetectionConfig1) | **Post** /anomalyDetection/aws/validator | Validate configuration of anomaly detection for AWS (for PUT request).


# **GetAwsAnomalyDetectionConfig1**
> AwsAnomalyDetectionConfig GetAwsAnomalyDetectionConfig1(ctx, )
Get configuration of anomaly detection for AWS.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**AwsAnomalyDetectionConfig**](AwsAnomalyDetectionConfig.md)

### Authorization

[ReadConfigToken](../README.md#ReadConfigToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateAwsAnomalyDetectionConfig1**
> UpdateAwsAnomalyDetectionConfig1(ctx, optional)
Update configuration of anomaly detection for AWS.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***UpdateAwsAnomalyDetectionConfig1Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UpdateAwsAnomalyDetectionConfig1Opts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **awsAnomalyDetectionConfig** | [**optional.Interface of AwsAnomalyDetectionConfig**](AwsAnomalyDetectionConfig.md)|  | 

### Return type

 (empty response body)

### Authorization

[WriteConfigToken](../README.md#WriteConfigToken)

### HTTP request headers

 - **Content-Type**: application/json; charset=utf-8
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ValidateAwsAnomalyDetectionConfig1**
> ValidateAwsAnomalyDetectionConfig1(ctx, optional)
Validate configuration of anomaly detection for AWS (for PUT request).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ValidateAwsAnomalyDetectionConfig1Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ValidateAwsAnomalyDetectionConfig1Opts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **awsAnomalyDetectionConfig** | [**optional.Interface of AwsAnomalyDetectionConfig**](AwsAnomalyDetectionConfig.md)|  | 

### Return type

 (empty response body)

### Authorization

[WriteConfigToken](../README.md#WriteConfigToken)

### HTTP request headers

 - **Content-Type**: application/json; charset=utf-8
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

