# ResponseTimeDegradationThresholdConfig

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResponseTimeThresholdMilliseconds** | **int32** | Response time during any 5-minute period to trigger an alert, milliseconds. | 
**SlowestResponseTimeThresholdMilliseconds** | **int32** | Response time of the 10% slowest during any 5-minute period to trigger an alert, milliseconds. | 
**LoadThreshold** | **string** | Mininal service load to detect response time degradation.   Response time degradation of services with smaller load won&#39;t trigger alerts. | 
**Sensitivity** | **string** | Sensitivity of the threshold.   With &#x60;low&#x60; sensitivity high statistical confidence is used, so brief violations (for example, due to a surge in load) won&#39;t trigger alerts.   With &#x60;high&#x60; sensitivity no statistical confidence is used. Each violation triggers alert. | 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


