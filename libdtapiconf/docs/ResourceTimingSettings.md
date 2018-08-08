# ResourceTimingSettings

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**W3cResourceTimings** | **bool** | W3C resource timings for third party/CDN monitoring settings enabled/disabled. | 
**NonW3cResourceTimings** | **bool** | Timing for JavaScript files and images on non W3C supported browsers enabled/disabled. | 
**NonW3cResourceTimingsInstrumentationDelay** | **int32** | Instrumentation delay for monitoring resource and image resource impact in browsers that don&#39;t offer W3C resource timings. Valid values range from 0 to 9999.  Only effective if **nonW3cResourceTimings** is enabled. | 
**ResourceTimingCaptureType** | **string** | Defines how detailed resource timings are captured.  Only effective if **w3cResourceTimings** or **nonW3cResourceTimings** is enabled. | 
**ResourceTimingsDomainLimit** | **int32** | Limits the number of domains for which W3C resource timings are captured.  Only effective if **resourceTimingCaptureType** is &#x60;CAPTURE_LIMITED_SUMMARIES&#x60;. | 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


