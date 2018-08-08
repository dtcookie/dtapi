# \AutomaticallyAppliedTagsApi

All URIs are relative to *https://siz65484.live.dynatrace.com/api/config/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAutoTag1**](AutomaticallyAppliedTagsApi.md#CreateAutoTag1) | **Post** /autoTags | Creates a new auto tag.
[**CreateOrUpdateAutoTag1**](AutomaticallyAppliedTagsApi.md#CreateOrUpdateAutoTag1) | **Put** /autoTags/{id} | Updates an existing auto tag.
[**DeleteAutoTag1**](AutomaticallyAppliedTagsApi.md#DeleteAutoTag1) | **Delete** /autoTags/{id} | Deletes the specified auto tag.
[**GetAllAutoTagConfigs1**](AutomaticallyAppliedTagsApi.md#GetAllAutoTagConfigs1) | **Get** /autoTags | Lists all configured auto tags.
[**GetSingleAutoTagConfig1**](AutomaticallyAppliedTagsApi.md#GetSingleAutoTagConfig1) | **Get** /autoTags/{id} | Shows the config properties of the specified auto tag.
[**ValidateAutoTag1**](AutomaticallyAppliedTagsApi.md#ValidateAutoTag1) | **Post** /autoTags/validator/{id} | Validate updates of existing auto tags for the &#x60;PUT /autoTags/{id}&#x60; request.
[**ValidateAutoTag2**](AutomaticallyAppliedTagsApi.md#ValidateAutoTag2) | **Post** /autoTags/validator | Validates new auto tags for the &#x60;POST /autoTags&#x60; request.


# **CreateAutoTag1**
> AutoTagStub CreateAutoTag1(ctx, optional)
Creates a new auto tag.

The body must not provide an ID as IDs are automatically assigned.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CreateAutoTag1Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CreateAutoTag1Opts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **autoTag** | [**optional.Interface of AutoTag**](AutoTag.md)|  | 

### Return type

[**AutoTagStub**](AutoTagStub.md)

### Authorization

[WriteConfigToken](../README.md#WriteConfigToken)

### HTTP request headers

 - **Content-Type**: application/json; charset=utf-8
 - **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateOrUpdateAutoTag1**
> CreateOrUpdateAutoTag1(ctx, id, optional)
Updates an existing auto tag.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | [**string**](.md)| The ID of the auto tag to be updated.   If you set the ID in the body as well, it must match this ID. | 
 **optional** | ***CreateOrUpdateAutoTag1Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CreateOrUpdateAutoTag1Opts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **autoTag** | [**optional.Interface of AutoTag**](AutoTag.md)|  | 

### Return type

 (empty response body)

### Authorization

[WriteConfigToken](../README.md#WriteConfigToken)

### HTTP request headers

 - **Content-Type**: application/json; charset=utf-8
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteAutoTag1**
> DeleteAutoTag1(ctx, id)
Deletes the specified auto tag.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | [**string**](.md)| The ID of the auto tag to be deleted. | 

### Return type

 (empty response body)

### Authorization

[WriteConfigToken](../README.md#WriteConfigToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAllAutoTagConfigs1**
> AutoTagStubList GetAllAutoTagConfigs1(ctx, )
Lists all configured auto tags.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**AutoTagStubList**](AutoTagStubList.md)

### Authorization

[ReadConfigToken](../README.md#ReadConfigToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSingleAutoTagConfig1**
> AutoTag GetSingleAutoTagConfig1(ctx, id, optional)
Shows the config properties of the specified auto tag.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | [**string**](.md)| The ID of the required auto tag. | 
 **optional** | ***GetSingleAutoTagConfig1Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetSingleAutoTagConfig1Opts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **includeProcessGroupReferences** | **optional.Bool**| Flag to include process group references to the response.    Process group references aren&#39;t compatible across environments. When this flag is set to false, conditions with process group references will be removed. If that leads to a rule having no conditions, the entire rule will be removed. | [default to false]

### Return type

[**AutoTag**](AutoTag.md)

### Authorization

[ReadConfigToken](../README.md#ReadConfigToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ValidateAutoTag1**
> ValidateAutoTag1(ctx, id, optional)
Validate updates of existing auto tags for the `PUT /autoTags/{id}` request.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | [**string**](.md)| The ID of the auto tag to be validated. | 
 **optional** | ***ValidateAutoTag1Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ValidateAutoTag1Opts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **autoTag** | [**optional.Interface of AutoTag**](AutoTag.md)|  | 

### Return type

 (empty response body)

### Authorization

[WriteConfigToken](../README.md#WriteConfigToken)

### HTTP request headers

 - **Content-Type**: application/json; charset=utf-8
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ValidateAutoTag2**
> ValidateAutoTag2(ctx, optional)
Validates new auto tags for the `POST /autoTags` request.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ValidateAutoTag2Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ValidateAutoTag2Opts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **autoTag** | [**optional.Interface of AutoTag**](AutoTag.md)|  | 

### Return type

 (empty response body)

### Authorization

[WriteConfigToken](../README.md#WriteConfigToken)

### HTTP request headers

 - **Content-Type**: application/json; charset=utf-8
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

