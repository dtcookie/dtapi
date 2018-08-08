# RequestAttribute

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metadata** | [**ConfigurationMetadata**](ConfigurationMetadata.md) |  | [optional] 
**Id** | **string** | The ID of the request attribute. | [optional] 
**Name** | **string** | The name of the request attribute. | 
**Enabled** | **bool** | Request attribute enabled/disabled. | 
**DataType** | **string** | The data type of the request attribute. | 
**DataSources** | [**[]DataSource**](DataSource.md) | The list of data sources. | 
**Normalization** | **string** | String values transformation.    If the **dataType** is not &#x60;string&#x60;, set the &#x60;Original&#x60; here. | 
**Aggregation** | **string** | Aggregation type for the request values. | 
**Confidential** | **bool** | Confidential data flag. Set &#x60;true&#x60; to treat the captured data as confidential. | 
**SkipPersonalDataMasking** | **bool** | Personal data masking flag. Set &#x60;true&#x60; to skip masking. Warning: This will potentially access personalized data. | 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


