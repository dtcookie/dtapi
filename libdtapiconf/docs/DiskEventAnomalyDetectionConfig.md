# DiskEventAnomalyDetectionConfig

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metadata** | [**ConfigurationMetadata**](ConfigurationMetadata.md) |  | [optional] 
**Id** | **string** | The ID of the disk event. | [optional] 
**Name** | **string** | The name of the disk event. | 
**Enabled** | **bool** | Whether the disk event is enabled. | 
**Metric** | **string** | The metric to monitor. | 
**Threshold** | **float64** | The event threshold. If the metric is &#x60;LowDiskSpace&#x60; or &#x60;LowInodes&#x60;, this is a percentage. For &#x60;ReadTimeExceeding&#x60; or &#x60;WriteTimeExceeding&#x60;, milliseconds. | 
**ViolatingSamples** | **int32** | The minimum number of violating samples for an event to be triggered. Must not exceed &#x60;samples&#x60;. | 
**Samples** | **int32** | The number of samples that are inspected. | 
**DiskNameFilter** | [**DiskNameFilterDto**](DiskNameFilterDto.md) |  | [optional] 
**TagFilters** | [**[]TagFilterDto**](TagFilterDto.md) | Only apply this rule to hosts having these tags. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


