# \DataPrivacyApi

All URIs are relative to *https://siz65484.live.dynatrace.com/api/config/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetDataPrivacySettings2**](DataPrivacyApi.md#GetDataPrivacySettings2) | **Get** /dataPrivacy | Lists the global data privacy settings.
[**UpdateDataPrivacySettings2**](DataPrivacyApi.md#UpdateDataPrivacySettings2) | **Put** /dataPrivacy | Updates the global data privacy settings.
[**ValidateConfiguration4**](DataPrivacyApi.md#ValidateConfiguration4) | **Post** /dataPrivacy/validator | Validates new data privacy settings for the &#x60;PUT /dataPrivacy&#x60; request.


# **GetDataPrivacySettings2**
> DataPrivacy GetDataPrivacySettings2(ctx, )
Lists the global data privacy settings.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**DataPrivacy**](DataPrivacy.md)

### Authorization

[ReadConfigToken](../README.md#ReadConfigToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateDataPrivacySettings2**
> UpdateDataPrivacySettings2(ctx, optional)
Updates the global data privacy settings.

This request updates global settings, affecting all your applications. To update application-specific data privacy settings, use the `PUT /applications/web/{id}/dataPrivacy` request.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***UpdateDataPrivacySettings2Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UpdateDataPrivacySettings2Opts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dataPrivacy** | [**optional.Interface of DataPrivacy**](DataPrivacy.md)|  | 

### Return type

 (empty response body)

### Authorization

[WriteConfigToken](../README.md#WriteConfigToken)

### HTTP request headers

 - **Content-Type**: application/json; charset=utf-8
 - **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ValidateConfiguration4**
> ValidateConfiguration4(ctx, optional)
Validates new data privacy settings for the `PUT /dataPrivacy` request.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ValidateConfiguration4Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ValidateConfiguration4Opts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dataPrivacy** | [**optional.Interface of DataPrivacy**](DataPrivacy.md)|  | 

### Return type

 (empty response body)

### Authorization

[WriteConfigToken](../README.md#WriteConfigToken)

### HTTP request headers

 - **Content-Type**: application/json; charset=utf-8
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

