# ConfigurationHistory

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metadata** | [**ConfigurationHistoryMetadata**](ConfigurationHistoryMetadata.md) |  | [optional] 
**From** | **int64** | Start timestamp of the report timeframe, in UTC milliseconds. | [optional] 
**To** | **int64** | End timestamp of the report timeframe, in UTC milliseconds. | [optional] 
**ResultsTruncated** | **bool** | If true, the queried timespan contained more results than could be delivered at once so only the most recent results were delivered, older ones truncated. | [optional] 
**Results** | [**[]ConfigurationInstanceDto**](ConfigurationInstanceDto.md) | List of all configuration changes during the specified timeframe. | [optional] 
**TimestampOfFirstTruncatedResult** | **int64** | If the result limit was exceeded, the timestamp of the first result that was truncated. Use this as your query&#39;s \&quot;to\&quot; to get the next batch of results. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


