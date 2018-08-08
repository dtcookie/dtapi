# \ApplicationDetectionConfigApi

All URIs are relative to *https://siz65484.live.dynatrace.com/api/config/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetConfiguration**](ApplicationDetectionConfigApi.md#GetConfiguration) | **Get** /applicationDetection | Lists all available application detection rules.
[**UpdateConfiguration**](ApplicationDetectionConfigApi.md#UpdateConfiguration) | **Put** /applicationDetection | Updates an application detection configuration.
[**ValidateConfiguration**](ApplicationDetectionConfigApi.md#ValidateConfiguration) | **Post** /applicationDetection/validator | Validates the application detection configuration for the &#x60;PUT /applicationDetection&#x60; request.


# **GetConfiguration**
> ApplicationDetectionConfig GetConfiguration(ctx, )
Lists all available application detection rules.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ApplicationDetectionConfig**](ApplicationDetectionConfig.md)

### Authorization

[ReadConfigToken](../README.md#ReadConfigToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateConfiguration**
> UpdateConfiguration(ctx, optional)
Updates an application detection configuration.

All the application detection rules are stored in a single configuration. When you want to edit a rule or create a new one, you have to put **all** existing rules in the request body, otherwise you'll lose them. Execute the `GET /applicationDetection` request to view the list of existing rules. Then modify an existing rule or add a new one, and keep the rest of the rules as they are.   The order of rules is important. Dynatrace evaluates rules from top to bottom. To add a new rule to the configuration, execute the `GET /applicationDetection` request, and place the new rule in the required position among the existing rules.    Validate the configuration with the `POST /validator` request before submitting it.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***UpdateConfigurationOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UpdateConfigurationOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **applicationDetectionConfig** | [**optional.Interface of ApplicationDetectionConfig**](ApplicationDetectionConfig.md)|  | 

### Return type

 (empty response body)

### Authorization

[WriteConfigToken](../README.md#WriteConfigToken)

### HTTP request headers

 - **Content-Type**: application/json; charset=utf-8
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ValidateConfiguration**
> ValidateConfiguration(ctx, optional)
Validates the application detection configuration for the `PUT /applicationDetection` request.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ValidateConfigurationOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ValidateConfigurationOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **applicationDetectionConfig** | [**optional.Interface of ApplicationDetectionConfig**](ApplicationDetectionConfig.md)|  | 

### Return type

 (empty response body)

### Authorization

[WriteConfigToken](../README.md#WriteConfigToken)

### HTTP request headers

 - **Content-Type**: application/json; charset=utf-8
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

